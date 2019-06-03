# DNS for Services and Pods

本文提供了 Kubernetes 对 DNS 支持的概述


<!-- vim-markdown-toc GFM -->

* [Introduction](#introduction)
    * [What things get DNS names](#what-things-get-dns-names)
* [Services](#services)
    * [A records](#a-records)
    * [SRV records](#srv-records)
* [Pods](#pods)
    * [Pod's hostname and subdomain fields](#pods-hostname-and-subdomain-fields)
    * [Pod's DNS Policy](#pods-dns-policy)
    * [Pod's DNS Config](#pods-dns-config)

<!-- vim-markdown-toc -->

## Introduction

Kubernetes DNS 在集群上 Schedules 了一个 DNS `pod` 和 `Service`，然后配置
kubelets 使相互独立的容器通过 DNS Services IP 解析 DNS name

### What things get DNS names

每个 service (包括 DNS server 本身) 都会在集群中分配一个 DNS name。默认
client pod DNS 搜索的区域为所属的 `namespace` 和 集群定义的 `domain`

by example:

假设在 Kubernetes namespace `bar` 中有一个 service name `foo` ，一个 pod 运行在
namespace `bar` 可以通过 DNS query `foo` 查询到该服务。 一个 pod 运行在 namespace `quux` 可以通过 DNS query `foo.bar` 查询到该服务



## Services

### A records

"Normal" (not headless，有cluster IP) 服务会被分配一个 DNS A 记录，例如：
`my-svc.my-namespace.svc.cluster.local` 会被解析为 Service 的集群 IP

"Headless" (没有 cluster IP) 服务也会分配一个 DNS A 记录，例如：
`my-svc.my-namespace.svc.cluster.local` 。不像 Normal
服务，它会解析为一个服务若干 pods 的 IP 集合， 客户端使用这个 IP
集合，或通过标准的 round-robin 方式挑选 IP

### SRV records

为 named ports 创建 SRV 记录，它属于 normal 或 [Headless Services](https://github.com/RosenLo/notes/blob/master/k8s/concepts.md#headless-services) 的一部分。

SRV record 的组成形式如： `_my-port-name._my-port-protocol.my-svc.my-namespace.svc.cluster.local`

对于 `Normal` 服务解析为端口号和服务名记录: `my-svc.my-namespace.svc.cluster.local`

对于 `Headless` 服务解析为多条包含端口和服务名记录: `auto-generated-name.my-svc.my-namespace.svc.cluster.local`


## Pods

### Pod's hostname and subdomain fields

现在创建一个 pod， 主机名默认就是 Pod's `metadata.name` 字段的值

Pod Spec 有一个可选字段 `hostname` , 可以用来指定 hostname 。当指定 `hostname`
将会优先使用，不会生成默认主机名

Pod Sepc 也有一个可选字段 `subdomain` ，用来指定 subdomain 。例如：一个 Pod
`hostname` 为 `foo`， `subdomain` 为 `bar` ，`namespace` 为 `my-namespace`，
FQDN (fully qualified domain name) 就是 `foo.bar.my-namespace.svc.cluster.local`

Example:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: default-subdomain
spec:
  selector:
    name: busybox
  clusterIP: None
  ports:
  - name: foo # Actually, no port is needed.
    port: 1234
    targetPort: 1234
---
apiVersion: v1
kind: Pod
metadata:
  name: busybox1
  labels:
    name: busybox
spec:
  hostname: busybox-1
  subdomain: default-subdomain
  containers:
  - image: busybox:1.28
    command:
      - sleep
      - "3600"
    name: busybox
---
apiVersion: v1
kind: Pod
metadata:
  name: busybox2
  labels:
    name: busybox
spec:
  hostname: busybox-2
  subdomain: default-subdomain
  containers:
  - image: busybox:1.28
    command:
      - sleep
      - "3600"
    name: busybox
```

如果在同一个 namespace 中存在一个 pod subdomain 为 `default-subdomain` ，和一个
headless 服务，名字也是 `default-subdomain` , 这个 pod 将会看到自己的 FQDN
名称为 `busybox-1.default-subdomain.my-namespace.svc.cluster.local` 。 DNS
提供一个 A 记录指向 Pod IP ， `busybox1` 和 `busybox2` 两者皆有独立的 A 记录

这个 Endpoint object 可以指定 `hostname` 为任意的 endpoint 地址以及它的 IP


### Pod's DNS Policy

DNS policies 可以基于每个 pod 制定，在 Pod.Sepc.`dnsPolicy` 字段。 现在 Kubernetes 支持以下 pod-specific DNS
policies

- `Default` - Pod 继承当前运行 Node 节点的 DNS 配置。更多细节请参考 [related discussion](https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#inheriting-dns-from-the-node)
- `ClusterFirst` - 如果 DNS 查询与配置的域名后缀不匹配，如 `www.kubernetes.io`， 将会被转发到该 Node 继承的上游 DNS Server 。 集群管理员可能配置了额外的上游 DNS Server
- `ClusterFirstWithHostNet` -  对于直接运行在主机网络的 Pods，需要显式申明 DNS Policy
- `None` - 允许 Pod 忽略 Kubernets 环境中 DNS 的配置，假设所有 DNS 的配置由
  `Pod.Spec.dnsConfig` 字段提供 。更多细节请参考 [pods dns config](#pods-dns-config)


**Note:** `Default` 并不是默认的 DNS Policy，如果 `dnsPolicy` 没有明确指定，默认为 `ClusterFrist`

下面 Pod 的示例展示了 DNS Policy 为 `ClusterFirstWithHostNet` ，因为它的
`hostNetwork` 为 `true`

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox
  namespace: default
spec:
  containers:
  - image: busybox:1.28
    command:
      - sleep
      - "3600"
    imagePullPolicy: IfNotPresent
    name: busybox
  restartPolicy: Always
  hostNetwork: true
  dnsPolicy: ClusterFirstWithHostNet
```

### Pod's DNS Config

Pods DNS Config 允许用户更好的控制 Pod DNS Settings

`dnsConfig` 字段是可选的，可以与任何 `dnsPolicy` 配置使用。如果 Pods `dnsPolicy` 设置为 `None`， 那么 `dnsConfig` 字段必须指定

以下是用户可以在 `dnsConfig` 字段中指定的属性：

- `namespace` - DNS Server 的 IP 列表，最多三个。如果 `dnsPolicy` 为 `None` ，
  那最少需要有一个 IP ，否则这个属性是可选的。IP 列表会和 DNS policy 的 DNS
Server IP 去重合并
- `searches` - 搜索域，最多六个，可选属性，会和 DNS Policy 的搜索域去重合并
- `options` - 一个可选的对象列表。每个对象包含一个 `name` 字段（必填）和一个
  `value` 字段（可选）。会和 DNS Policy 的 options 去重合并

以下是自定义 DNS 的 Pod 一个示例：

service/networking/custom-dns.yaml

```yaml
apiVersion: v1
kind: Pod
metadata:
  namespace: default
  name: dns-example
spec:
  containers:
    - name: test
      image: nginx
  dnsPolicy: "None"
  dnsConfig:
    nameservers:
      - 1.2.3.4
    searches:
      - ns1.svc.cluster.local
      - my.dns.search.suffix
    options:
      - name: ndots
        value: "2"
      - name: edns0
```

当上面的 Pod 创建成功，`/etc/resolv.conf` 将会有以下内容：

```
nameserver 1.2.3.4
search ns1.svc.cluster.local my.dns.search.suffix
options ndots:2 edns0
```

对于 IPv6 设置，search 和 nameserver 将会类似于以下内容：

```
$ kubectl exec -it dns-example -- cat /etc/resolv.conf
nameserver fd00:79:30::a
search default.svc.cluster.local svc.cluster.local cluster.local
options ndots:5
```

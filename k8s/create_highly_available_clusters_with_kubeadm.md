<!-- vim-markdown-toc GFM -->

* [使用 Kubeadm 安装 Kubernetes 集群记录](#使用-kubeadm-安装-kubernetes-集群记录)
    * [Environment](#environment)
    * [Prepaering](#prepaering)
    * [Installing kubeadm](#installing-kubeadm)
        * [Installing Docker](#installing-docker)
        * [Installing kubeadm, kubelet and kubectl](#installing-kubeadm-kubelet-and-kubectl)
    * [Stacked control plane and etcd nodes](#stacked-control-plane-and-etcd-nodes)
        * [Initialization](#initialization)
            * [Installing Pod Network](#installing-pod-network)
            * [Configure SSH](#configure-ssh)
            * [Copy Certificate To Other Control plane](#copy-certificate-to-other-control-plane)
        * [节点加入集群](#节点加入集群)
        * [查看各组件状态](#查看各组件状态)
        * [Reset](#reset)
        * [Troubleshooting](#troubleshooting)
    * [Reference](#reference)

<!-- vim-markdown-toc -->

# 使用 Kubeadm 安装 Kubernetes 集群记录

## Environment

- CentOS 7.2.1511

## Prepaering

Master  |   ip  |
node1
node2

以下所有操作都是以`root` 用户执行


## Installing kubeadm

### Installing Docker

- 安装 Docker 所需要的包

```bash
yum install -y yum-utils device-mapper-persistent-data lvm2
```

- 添加 Docker repo

```bash
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
```

- 安装指定版本

```bash
yum install docker-ce-18.06.0.ce
```


### Installing kubeadm, kubelet and kubectl

因为墙内的原因无法选择官方 repo 源安装，这里使用 aliyun 提供的 repo

```bash
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
EOF
```

关闭 selinux, 需要允许容器去访问宿主机文件系统

```bash
setenforce 0

sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
```

有 issue 报告了在 RHEL/CentOS 7 下因为 iptables
导致流量路由不正确，需要以下设置

```bash
cat <<EOF >  /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
```

需要关闭 swap ， kubelet 才能正常工作

```bash
swapoff -a

vi /etc/stab # 注释掉 swap

echo "vm.swappiness=0" >> /etc/sysctl.d/k8s.conf

sysctl --system
```

```
yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes

systemctl enable kubelet && systemctl start kubelet
```

## Stacked control plane and etcd nodes

Kubeadm HA topology - stacked etcd

<div> <img src="../assets/kubeadm-ha-topology-stacked-etcd.svg"/> </div><br>

### Initialization

```
kubeadm init --apiserver-advertise-address=0.0.0.0 --pod-network-cidr=10.244.0.0/16  --kubernetes-version=v1.13.0
```

成功后会得到如下信息:

```
[init] Using Kubernetes version: v1.13.0
[preflight] Running pre-flight checks
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Activating the kubelet service
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [k8s10-001 localhost] and IPs [10.10.151.45 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [k8s10-001 localhost] and IPs [10.10.151.45 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [k8s10-001 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.10.151.45]
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
[control-plane] Creating static Pod manifest for "kube-scheduler"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 23.003435 seconds
[uploadconfig] storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.13" in namespace kube-system with the configuration for the kubelets in the cluster
[patchnode] Uploading the CRI Socket information "/var/run/dockershim.sock" to the Node API object "k8s10-001" as an annotation
[mark-control-plane] Marking the node k8s10-001 as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node k8s10-001 as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: c726la.x9qplglrr1uyxcbi
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstraptoken] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstraptoken] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstraptoken] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstraptoken] creating the "cluster-info" ConfigMap in the "kube-public" namespace
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes master has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

You can now join any number of machines by running the following on each node
as root:

  kubeadm join 10.10.151.45:6443 --token c726la.x9qplglrr1uyxcbi --discovery-token-ca-cert-hash sha256:abd6143f181f188eea9cc8d039510789163e11d3b49815b79d8f220c3204f731
```

通过上面的输出，可以看到有以下关键步骤：

- [kubelet-start] 生成kubelet的配置文件"/var/lib/kubelet/config.yaml"
- [certificates] 生成相关的各种证书
- [kubeconfig] 生成相关的kubeconfig文件
- [bootstraptoken] 生成token记录下来，后边使用kubeadm join往集群中添加节点时会用到


查看集群状态

```bash
kubectl get cs

NAME                 STATUS    MESSAGE              ERROR
controller-manager   Healthy   ok
scheduler            Healthy   ok
etcd-0               Healthy   {"health": "true"}
```


#### Installing Pod Network

```bash
wget https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

kubectl apply -f  kube-flannel.yml
```


#### Configure SSH

生成密钥

```bash
ssh-keygen
```

配置免密登录

```bash
#!/bin/env bash
CONTROL_PLANE_IPS="k8s09-001 k8s09-002"
for host in ${CONTROL_PLANE_IPS}; do
    cat ~/.ssh/id_rsa.pub | ssh root@$host "cat - >> ~/.ssh/authorized_keys"
done
```

#### Copy Certificate To Other Control plane

在 master 上执行下面的 shell 脚本

```bash
#!/bin/env bash
USER=root
CONTROL_PLANE_IPS="k8s09-001 k8s09-002"
for host in ${CONTROL_PLANE_IPS}; do
    scp /etc/kubernetes/pki/ca.crt "${USER}"@$host:
    scp /etc/kubernetes/pki/ca.key "${USER}"@$host:
    scp /etc/kubernetes/pki/sa.key "${USER}"@$host:
    scp /etc/kubernetes/pki/sa.pub "${USER}"@$host:
    scp /etc/kubernetes/pki/front-proxy-ca.crt "${USER}"@$host:
    scp /etc/kubernetes/pki/front-proxy-ca.key "${USER}"@$host:
    scp /etc/kubernetes/pki/etcd/ca.crt "${USER}"@$host:etcd-ca.crt
    scp /etc/kubernetes/pki/etcd/ca.key "${USER}"@$host:etcd-ca.key
    scp /etc/kubernetes/admin.conf "${USER}"@$host:
done
```

在 node 上执行下面的 shell 脚本

```bash
#!/bin/env bash
USER=root
mkdir -p /etc/kubernetes/pki/etcd
mv /${USER}/ca.crt /etc/kubernetes/pki/
mv /${USER}/ca.key /etc/kubernetes/pki/
mv /${USER}/sa.pub /etc/kubernetes/pki/
mv /${USER}/sa.key /etc/kubernetes/pki/
mv /${USER}/front-proxy-ca.crt /etc/kubernetes/pki/
mv /${USER}/front-proxy-ca.key /etc/kubernetes/pki/
mv /${USER}/etcd-ca.crt /etc/kubernetes/pki/etcd/ca.crt
mv /${USER}/etcd-ca.key /etc/kubernetes/pki/etcd/ca.key
mv /${USER}/admin.conf /etc/kubernetes/admin.conf
```

### 节点加入集群

```bash
kubeadm join k8s10-001:6443 --token c726la.x9qplglrr1uyxcbi --discovery-token-ca-cert-hash sha256:abd6143f181f188eea9cc8d039510789163e11d3b49815b79d8f220c3204f731 --experimental-control-plane
```

- experimental-control-plane: 这个 flag 表示加入集群作为 control plane

### 查看各组件状态

```bash
kubectl -n kube-system get pods,svc,deployment -o wide
```

### Reset

```bash
kubeadm reset
ifconfig cni0 down
ip link delete cni0
ifconfig flannel.1 down
ip link delete flannel.1
rm -rf /var/lib/cni/
```

### Troubleshooting

- 查看kubelet 日志

    ```bash
    journalctl -f -u kubelet
    ```


## Reference
- [Docker Docs](https://docs.docker.com/)
- [Installing kubeadm](https://kubernetes.io/docs/setup/independent/install-kubeadm/)
- [Creating Highly Available Clusters with kubeadm](https://kubernetes.io/docs/setup/independent/high-availability/)
- [HA Topology](https://kubernetes.io/docs/setup/independent/ha-topology/)

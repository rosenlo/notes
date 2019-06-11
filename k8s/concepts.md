The Concepts section helps you quickly learn about the parts of the Kubernetes system and abstractions Kubernetes uses to represent your cluster, and helps you
obtain adopt understanding of how Kubernetes works.


<!-- vim-markdown-toc GFM -->

* [Overview](#overview)
* [Kubernetes Objects](#kubernetes-objects)
    * [Understanding Kubernetes Objects](#understanding-kubernetes-objects)
        * [Object Spec and Stauts](#object-spec-and-stauts)
        * [Describing a Kubernetes Object](#describing-a-kubernetes-object)
            * [Required Fields](#required-fields)
    * [Pods](#pods)
        * [Termination of Pods](#termination-of-pods)
        * [Disruptions](#disruptions)
            * [Voluntary and Involuntary Disruptions](#voluntary-and-involuntary-disruptions)
            * [Dealing with Disruptions](#dealing-with-disruptions)
    * [Service](#service)
        * [Headless Services](#headless-services)
            * [With selectors](#with-selectors)
            * [Without selectors](#without-selectors)
    * [Volumes](#volumes)
    * [Namespace](#namespace)
    * [ReplicaSet](#replicaset)
    * [Deployment](#deployment)
    * [StatefulSet](#statefulset)
        * [Using StatefulSet](#using-statefulset)
        * [Limitations](#limitations)
        * [Components](#components)
        * [Pod Selector](#pod-selector)
        * [Pod Identity](#pod-identity)
            * [Ordinal Index](#ordinal-index)
            * [Stable Network ID](#stable-network-id)
            * [Stable Storage](#stable-storage)
            * [Pod Name Label](#pod-name-label)
        * [Deployment and Scaling Guarantees](#deployment-and-scaling-guarantees)
            * [Pod Management Policies](#pod-management-policies)
            * [OrderedReady Pod Management](#orderedready-pod-management)
            * [Parallel Pod Management](#parallel-pod-management)
        * [Update Strategies](#update-strategies)
            * [On Delete](#on-delete)
            * [Rolling Updates](#rolling-updates)
                * [Partitions](#partitions)
                * [Forced Rollback](#forced-rollback)
    * [DaemonSet](#daemonset)
    * [Job](#job)
* [Kubernetes Control Plane](#kubernetes-control-plane)
    * [Kubernetes Master](#kubernetes-master)
    * [Kubernetes Nodes](#kubernetes-nodes)

<!-- vim-markdown-toc -->

# Overview

- You can use Kubernetes API objects to describe your cluster's state:
    - What applications or other workloads you want to run
    - What container images they use
    - The number of reqplicas
    - What network and disk resource you want to make available
    - ...
- You can use the command-line interface, `kubectl`
- You can also use the Kubernetes API directly to interact with the cluster and set or modify you desired state.
- The Kubernetes Master is a collection of three proccesses that run on a single node in your cluster, which is designated as the master node. Those processes are:
    - [Kube-apiserver](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/)
    - [Kube-controller-manager](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/)
    - [Kube-scheduler](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/)
- Each individual non-master node in your cluster runs two processes:
    - [Kublet](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/), which communicates with the Kubernetes Master.
    - [Kube-proxy](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/), a network proxy which reflects Kubernetes networking services on each node.


# Kubernetes Objects

The Basic Kubernetes objects include:

<!-- GFM-TOC -->
* [Pods](#Pods)
* [Service](#Service)
    * [Headless Services](#headless-services)
* [Volumes](#Volumes)
* [Namespace](#Namespace)
<!-- GFM-TOC -->

The Higher-level Kubernetes objects called Controllers, Controllers build upon
the basic objects, and provide additional functionality and convenience
features. they include:

<!-- GFM-TOC -->
* [ReplicaSet](#ReplicaSet)
* [Deployment](#Deployment)
* [StatefulSet](#StatefulSet)
* [DaemonSet](#DaemonSet)
* [Job](#job)
<!-- GFM-TOC -->


## Understanding Kubernetes Objects

Kubernetes Objects 在 Kubernetes 系统中表现为持久化实体，用来展示你应用的状态。

- 容器应用运行在哪个节点
- 应用的可用资源
- 管理应用的策略如：重启、升级、容错

要使用 Kubernetes Objects 去创建、修改、删除。可以使用 [Kubernetes API](#Kubernetes-API) 的命令行工具 `kubectl` 。
还可以在自己的程序中直接使用 [Client Libraries](https://kubernetes.io/docs/reference/using-api/client-libraries/)

### Object Spec and Stauts

每个 Kubernetes Object 都包含了两个嵌套对象字段（Spec、Status），它们管理着 Object 配置。

- Spec - 描述 Object 的所需状态，希望具备哪些特征，用户必须提供
- Status - 描述 Object 真实的状态，由 Kubernetes 系统提供，且时时刻刻在管理
  Object 的状态与 Spec 提供的状态保持一致

For example:

一个 Kubernetes Deployment 是一个 Object
。它可以表示一个应用在集群中运行，当你创建 Deployment
，希望运行三个实例，需要在 Spec 中设置 `replicas: 3` 。Kubernetes
系统会读取 Deployment Spec 然后启动三个实例运行你的应用，如果有失败或 status
改变，Kubernetes 修复状态直到与 Spec 中状态一致。

更多关于 Object Spec, Status 和 metadata 在 [Kubernetes API Conventions](#https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md)


### Describing a Kubernetes Object

当你在 Kubernetes 中创建一个 object， 你需要提供一些必要的信息关于这个 object 如：`name`

当你使用 Kubernetes API 去创建一个 object （或直接通过`kubectl`） ，这个 API
请求 body 是 JSON 格式，但大多数情况下，我们会写一个`.yaml`的文件然后使用`kubectl` 创建 object

For example：

```yaml
apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2 # tells deployment to run 2 pods matching the template
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```

```shell
kubectl apply -f https://k8s.io/examples/application/deployment.yaml --record
```

输出类似于:

```
deployment.apps/nginx-deployment created
```

#### Required Fields

在 `.yaml` 文件中需要设置以下几个字段

- `apiVersion` - 哪个 Kubernetes API 版本
- `kind` - 你想创建的类型
- `metadata` - 有助于唯一识别的数据，包含了 `name` , UID 和 可选的 `namespace`

更多关于 Kubernetes spec 字段:

- [Kubernetes API Reference](#https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/)
- [Sepc for Pod](#https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#podspec-v1-core)
- [Spec for Deployment](#https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#deploymentspec-v1-apps)



## Pods

A Pod are the smallest deployable unit of computing that can be created and
managed in Kubernetes, pods represent running processes on nodes in cluster.

- A Pod always run on a **Node**
- Shared stroage, as Volumes
- Networking, as a unique cluster IP address
- Information about how to run each container, such as the container image version or specific ports to use

<div> <img src="../assets/pods.svg" width="500"/> </div><br>

Pod 的共享上下文是一组 Linux namespaces, cgroups 以及可能的隔离方面。与 Docker
容器的隔离相似。在 Pod 的上下文中，个别的应用程度可能会应用进一步的子隔离。

Pod 里的容器共享 IP 和端口，可以通过`localhost`
找到彼此。通过标准进程间通讯（如： `SystemV semaphores`,
`POSIX`）共享内存相互通讯。不同 Pods 中的容器有不同的 IP
，如果没有[特殊配置](https://kubernetes.io/docs/concepts/policy/pod-security-policy/)，不能相互通讯

应用在 Pod 中可以访问共享[Volumes](#Volumes)，它被定义为 Pod
的一部分，可挂载在每个应用文件系统。

就 Docker 构造而言，Pod 建模为一组具有共享 namespaces 和 共享文件系统卷的 Docker 容器。

与单个应用容器一样， Pod 被认为是相对短暂（而不是需要持久化）的实体。正如 [Pod
生命周期](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)讨论一样，Pods
创建分配一个唯一 ID(UID)，然后调度到 Nodes 直到被终止（根据重启 Policy ）或被删除。如果 Node
失联，调度在该 Node 上的 Pods 将在超时后删除。给定的 Pod （由 UID
定义）不能重新调度到新的 Node；相反，它可以被相同的 Pod
替换，甚至可以用相同的名字，不过会生成一个新的 UID 。更多细节参考 [replication
controller](#https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/)

当有些东西被认为与 Pod 有同样的生命周期，例如[Volumes](#Volumes)，那么意味着与
Pod 共存亡。假设 Pod 由于某种原因被删除了，甚至是被替换创建，那与 Pod
有相同生命周期的东西（如：Volumes）也会被删除后再创建。


### Termination of Pods

Pod 表示为集群中节点上的一个进程，重要的是允许进程优雅终止（vs 强制
Kill，应用没有机会回收清理）。用户可以请求删除并且知道进程何时终止，还能保证删除最终完成。
当用户请求删除一个 Pod  时，系统在允许强制删除 Pod 前记录宽限期。

### Disruptions

本部分适用于想要构建高可用应用程序的人，需要了解 Pod 可能会发生的中断类型。

也适用于希望执行自动集群操作（如：升级、集群自动伸缩）的集群管理员

#### Voluntary and Involuntary Disruptions

Pod 不会自己消亡，除非被人为或 controller
销毁，也有可能是因为硬件故障或软件错误导致。

硬件故障的情况叫做 Involuntary disruptions。 example：

- 物理机的硬件故障导致 Node 破坏
- 集群管理员错误删除了 VM 实例
- 云厂商虚拟化失败导致的 VM 失联
- 内核 panic
- 集群网络分区导致节点失联
- 由于节点[资源不足](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/)导致 pod 回收

除了资源不足，其他条件大部分都很常见，并不是 kubernetes 特有的。

而其他情况叫做 voluntary disruptions，包括不限于应用程序初始化或通过集群管理员初始化。 典型应用程序操作包括：

- 删除管理 pod 的 deployment 或者其他 controller
- 更新 deployment template 引起的 restart
- 直接删除一个 pod （例如：意外）

集群管理员操作包括：

- 踢掉一台 node ，修复或升级
- [集群自动伸缩](https://kubernetes.io/docs/tasks/administer-cluster/cluster-management/#cluster-autoscaler)一台 node
- 从 node 移除一个 pod ，允许其他内容与 node 相符

注意：不是所有 voluntary disruptions 都受限于 Pod Disruptions
Budgets。例如：删除 deployment 或 pods 会绕过 Pod Disruptions Budgets


#### Dealing with Disruptions

下面有些方法可以缓解 involuntary disruptions

- 确保 pod [申请所需资源](https://kubernetes.io/docs/tasks/configure-pod-container/assign-memory-resource/)
- 为应用创建多个副本，如果需要高可用（了解更多关于运行[无状态应用](https://github.com/RosenLo/notes/blob/master/k8s/run_stateless_application_deployment.md)和[有状态应用](https://github.com/RosenLo/notes/blob/master/k8s/concepts.md#statefulset)）
- 甚至运行副本的应用更高的可用性

## Service


### Headless Services

有时不需要 load-balancing 和一个 service IP，可以创建一个 `headless`
服务通过指定 `None` 在 `.spec.clusterIP` 字段

这个选项让开发者与 Kubernetes 系统解构，允许他们自由的使用服务发现系统(例如 zookeeper)。
应用仍然可以使用自注册的模式并且可以轻松地在此 API 上构建用于其他服务发现系统的适配器

对于 `Services`， clusterIP 不用分配， `kube-proxy` 不会处理这个 services
并且也没有 load-balancing 和 proxying ，DNS 自动配置依赖 service 的 `selectors`
定义

#### With selectors

对于定义了 selectors 的 headless 服务 ， endpoints controller 会在这个 API 创建 `Endpoints` 记录，然后修改 DNS 配置，返回一个直接指向 Pods IP 的 A 记录

#### Without selectors

对于未定义 selectors 的 headless 服务， endpoints controller 不会创建
`Endpoints` 记录。无论如何， DNS 系统会通过以下方式寻找和配置：

- [ExternalName](#externalname) 的 CNAME 记录
- 适用于所有其他类型, 任何共享同一个名称的 `Endpoints` 记录

## Volumes

容器中的磁盘文件是短暂存在的，这会带来一些问题，如：

- 一些需要持久化数据的应用，当容器 Crash 时，`kubelet` 会重启容器（创建一个新容器，再销毁老容器），这时文件就会丢失。
- 多个容器在一个 Pod 中共享文件

而 Kubernetes 的 `Volume` 则解决了这些问题

建议先熟悉 [Pods](#Pods)

## Namespace

## ReplicaSet
## Deployment

A Depolyment controller provides declarative updates for Pods and ReplicaSets.

- Manage the creation and scaling of pods.

## StatefulSet

管理 `Deployment` 和 `Pods` 扩展，对 Pods 的排序和唯一性提供保障。

和 `Deployment` 一样，`StatefulSet` 管理 Pod 基于相同的容器 `spec`
。不一样的是， `StatefulSet` 对每个 Pod 维护了一个持久的标识，这些 Pod 被同样的
`spec` 创建，但是不可变的。每个都有一个持久化标识，它在重新调度后都会保留。

StatefulSet 与其他 Controller 在相同的模式下运转。在 StatefulSet
对象定义的状态， StatefulSet controller 会进行任何有必要的更新达到此状态。

### Using StatefulSet

StatefulSets 对有状态的应用程序非常有用，有以下要求：

- 稳定, 独特的网络标识
- 持久化存储
- 有序，优雅的部署方式和扩展
- 自动滚动升级

根据以上，稳定和 Pod 的重新调度的持久性同义。如果你的应用不要求任何稳定标识符或顺序的部署、删除、扩展，可以无状态部署。
例如使用 [Deployment](#Deployment) 或 [ReplicaSet](#ReplicaSet) 会更适合。

### Limitations

- StatefulSet 在 1.9 版本之前是 beta 资源，在 1.5 版本之甚至还没有
- 为了确保数据安全，删除或扩展 StatefulSet 不会删除相关 volumes
- StatefulSet 现在要求创建一个 [Headless Service](#headless-service) 为 Pod
  提供网络识别
- StatefulSet 在删除时，不会提供任何终止 Pod 的保障，要达到有序、优雅终止 Pod
  ，需要在删除前伸缩 StatefulSet replica 为 0
- 当使用默认的 [Pod 管理策略](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-management-policies) (`OrderedReady`) [Rolling Updates](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#rolling-updates)，可能会进入一个破碎状态，需要 [手动干预修复](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#forced-rollback)

### Components

以下的例子演示了 StatefulSet 组件

- 一个名字叫 nginx 的Headless Service，用来控制网络域
- 一个名字叫 web 的 StatefulSet object， 声明了 3 个 nginx 容器副本在唯一的 Pod
  中启动
- volumenClaimTemplates 通过 PersistentVolume Provisioner 配置
  [PersistentVolumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) 提供稳定存储

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
  - port: 80
    name: web
  clusterIP: None
  selector:
    app: nginx
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
spec:
  selector:
    matchLabels:
      app: nginx # has to match .spec.template.metadata.labels
  serviceName: "nginx"
  replicas: 3 # by default is 1
  template:
    metadata:
      labels:
        app: nginx # has to match .spec.selector.matchLabels
    spec:
      terminationGracePeriodSeconds: 10
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "my-storage-class"
      resources:
        requests:
          storage: 1Gi
```

### Pod Selector

必须配置 StatefulSet `.spec.selector` 字段匹配 `.spec.template.metadata.labels`
的标签。 在 Kubernetes 1.8 之前， `.spec.selector` 字段省略时为默认值。
在1.8之后的版本，如果没有指定匹配 Pod Selector， 在创建 StatefulSet
时会引起 validation 错误

### Pod Identity

StatefulSet Pods 具有唯一标识，是一组具有稳定网络标识和稳定存储的顺序组成。这个标识附属在 Pod
不管它重新调度在哪个节点上都会保留。

#### Ordinal Index

对具有 N 个副本的 StatefulSet ，每个 Pod 都会分配一个顺序数字，从 0 到 N - 1
，在这个 Set 上唯一。

#### Stable Network ID

每个在 StatefulSet 中的 Pod 都从 StatefulSet 的名字和 Pod
的顺序衍生出来，这个构建的形式是 `$(statefulset name)-$(ordinal)`
，上面的示例会创建出三个 Pod 名字 `web-0,web-1,web-2` 。而这些 Pod 可以通过
[Headless Service](#headless-service) 控制访问，这个服务管理域名的形式为： `$(service name).$(namespace).svc.cluster.local` ，"cluster.local" 就是 cluster 的域，
每个 Pod 创建会得到一个 DNS name， 形式为： `$(podname).$(governing service domain)`，这个 governing service 在 StatefulSet 中的 `serviceName` 定义。

在 [limitations](#limitations) 中说到，有必要创建一个 [Headless Service](#headless-service) 作为 Pods 的网络标识

#### Stable Storage

kubernetes 为每个 VolumeClaimTemplate 创建一个 PersistentVolume 。 在上面 nginx
的示例中，每个 Pod 会接收到一个名字叫 `my-storage-class` 的 StorageClass 和
1Gib 额度的存储。如果没指定 StorageClass，默认的 StorageClass 被使用。当 Pod
被重新调度到一个节点， 它的 `volumeMounts` 会根据 PersistentVolume Claims
挂载相关的 PersistentVolume 。 注意在删除时这个 PersistentVolume
不会被删除，如果要删除需要手动处理。

#### Pod Name Label

当 StatefulSet 创建 Pod 时， 会添加一个 label: `statefulset.kubernetes.ioo/pod-name: $(pod name)` 。这个 label 允许 Service 映射到 StatefulSet 指定的 Pod

### Deployment and Scaling Guarantees

- 对于一个具有 N 个副本的 StatefulSet 来说，Pod 以顺序的形式创建: {0..N-1}
- 当 Pods 被删除，以倒序的形式删除: {N-1..0}
- 伸缩扩展 Pod 操作前置条件为：所有 Pod 必须 Running 和 Ready
- 终止 Pod 操作前置条件为：所有 Pod 必须完全 shutdown

StatefulSet 不应该指定 `pod.Spec.TerminationGracePeriodSeconds` 为 0
，这非常不安全，强烈反对。更多信息参考 [force deleting StatefulSet Pods](https://kubernetes.io/docs/tasks/run-application/force-delete-stateful-set-pod/)

当上面 nginx 的示例创建，三个 Pods 会顺序创建 web-0, web-1, web-2。如果 web-0
没有 [Running and Ready](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/) web-1 不会被创建并且 web-2 也不会被创建，直到 web-1 Running and Ready 。
如果在 web-1 Running and Ready 之后 web-0 失败了， web-2 不会被创建，直到 web-0
重新启动并且 Running and Ready

如果用户调整副本数为 `replicas=1` ，那么 web-2 首先被终止， web-1 会先等 web-2
完全 shutdown 和删除后再终止、删除。如果 web-0 在 web-2 终止、删除后 web-1 未删除的情况下状态进入 fail，那么 web-1 会保留直到 web-0 恢复 Running and Ready

#### Pod Management Policies

在 kubernetes 1.7 之后， StatefulSet 放宽了有序保障，通过提供 `.spec.podManagementPolicy` 字段来预防唯一性身份标识

#### OrderedReady Pod Management

`OrderedReady` pod management 是默认的 StatefulSet 策略，行为参考 [Deployment and Scaling Guarantees](#deployment-and-scaling-guarantees)

#### Parallel Pod Management

`Parallel` pod management 告诉 StatefulSet controller 并行去启动或删除
Pods，不等待 Pods Running and Ready 或完全 terminated。这个选项仅影响 scaling
操作， Updates 不受影响。


### Update Strategies

在 Kubernetes 1.7 之后， StatefulSet `.spec.updateSrategy`
字段允许为 StatefulSet 中的 Pod 配置和禁用容器，labels、 resource
request/limits 和 annotations 的自动 rolling updates

#### On Delete

这个 `OnDelete` 更新策略实现了 1.6 之前的遗留行为。当一个 StatefulSet
`.spec.updateStrategy.type` 设置为 `OnDelete`， StatefulSet controller
不会自动更新 Pods，用户需要手动删除 pod 去触发 controller 创建一个新的 pods

#### Rolling Updates

这个 `RollingUpdate` 更新策略实现了 Pods 的滚动升级，这也是默认策略当没有指定
`.spec.updateStrategy` 字段。

当设置为 `RollingUpdate`，这个 StatefulSet controller 会倒序删除并重建每一个
Pod，也就是删一个重建一个。

##### Partitions

`RollingUpdate` 还可以设置 partition :
`.spec.updateStrategy.rollingUpdate.partition`，如果指定了
partition，当 StatefulSet `.sepc.template` 更新，那么所有序号大于或等于 partition 的 Pod 将会进行更新，小于的则不会更新。及时被删除了也会重建之前的版本。
如果 StatefulSet 的 `.spec.updateStrategy.rollingUpdate.partition` 大于 `.spec.replicas` 则更新其 `.spec.template` 不会被传播到 Pods，在大多数情况下用不到 partition，
但对于需要稳定更新、回滚、金丝雀或分阶段发布非常有用

##### Forced Rollback

使用默认 [Pod 管理策略](#pod-management-policy) `OrderedReady` 的 [Rolling Updates](#rolling-updates) 时，如果进入 broken 状态，有可能需要手工修复。

如果更新 Pod template 到一个永远无法 Running and Ready 的配置
（如：一个有问题的二进制文件或应用层面的配置错误），StatefulSet 会停止
rollout 并且进入等待。

在这种状态下，将 Pod template 配置恢复成正确是不够的，因为一个[已知 issue](https://github.com/kubernetes/kubernetes/issues/67250)
StatefulSet 会持续去等待 broken 的 pod 恢复成 Ready
（永远不会发生），再尝试恢复成可运行的配置

在恢复 template 后，需要删除 StatefulSet 创建的所有 Pods，然后 StatefulSet
会使用恢复 template 后的配置重建创建 Pods


## DaemonSet
## Job

# Kubernetes Control Plane

## Kubernetes Master

- The Master is responsible for maintaining the desired state for your cluster
- The Master automatically handles scheduling the pods across the Nodes in the cluster
- The Master's automatic scheduling takes into account the available resources on each Node

<div> <img src="../assets/cluster.svg" width="400"/> </div><br>



## Kubernetes Nodes
A Node is a worker machine in Kubernetes and may be either a virtual or a physical machine, A Node can have multiple pods. Each Node is managed by the Master.

Every kubernetes Node runs at least:

- Kubelet, a process responsible for communication between the kubernetes Masster and the Node; it manages the Pods and the containers running on a machine.
- A container runtime (like Docker, rkt) responsible for pulling the container image from a registry, unpacking the container, and running the application.


<div> <img src="../assets/nodes.svg" width="400"/> </div><br>



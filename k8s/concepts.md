The Concepts section helps you quickly learn about the parts of the Kubernetes system and abstractions Kubernetes uses to represent your cluster, and helps you
obtain adopt understanding of how Kubernetes works.


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
* [Volume](#Volume)
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

## Volume
## Namespace

## ReplicaSet
## Deployment

A Depolyment controller provides declarative updates for Pods and ReplicaSets.

- Manage the creation and scaling of pods.

## StatefulSet
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



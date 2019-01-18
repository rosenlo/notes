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
    - [Kube-proxy](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/), a network proxy which relects Kubernetes networking services on each node.


# Kubernetes Objects

The Basic Kubernetes objects include:

<!-- GFM-TOC -->
* [Pods](#Pods)
* [Service](#Service)
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



## Pods

A Pod are the smallest deployable unit of computing that can be created and
managed in Kubernetes, pods represent running processes on nodes in cluster.

- A Pod always run on a **Node**
- Shared stroage, as Volumes
- Networking, as a unique cluster IP address
- Information about how to run each container, such as the container image version or specific ports to use

<div> <img src="../assets/pods.svg" width="500"/> </div><br>



## Service
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



# 使用 Deployment 运行一个无状态应用

本文展示如何用 Kubernetes Deployment object 运行一个 nginx 应用

<!-- vim-markdown-toc GFM -->

* [Objectives](#objectives)
* [Before you begin](#before-you-begin)
* [Creating and exploring an nginx deployment](#creating-and-exploring-an-nginx-deployment)
* [Updating the deployment](#updating-the-deployment)
* [Scaling the application by increasing the replica count](#scaling-the-application-by-increasing-the-replica-count)
* [Deleting a deployment](#deleting-a-deployment)
* [Reference](#reference)

<!-- vim-markdown-toc -->


## Objectives

- 创建一个 `nginx` deployment
- 使用`kubectl` 查看 deployment 的详细信息
- 更新 deployment

## Before you begin

在开始之前，需要有一个 kubernetes 集群和配置好和集群通讯的命令行工具 `kubectl`
。如果还没有集群，可以使用[Minikube](https://kubernetes.io/docs/setup/minikube/) 或 [Kubeadm](https://github.com/RosenLo/notes/blob/master/k8s/create_highly_available_clusters_with_kubeadm.md) 创建一个集群。
也可以直接使用官方提供的在线体验环境:

- [Katacoda](https://www.katacoda.com/courses/kubernetes/playground)
- [Play with Kubernetes ](https://labs.play-with-k8s.com/)


## Creating and exploring an nginx deployment

创建一个 YAML 格式的 Nginx Deployment 描述文件，运行 nginx:1.7.9 的 Docker image

- 创建 nginx-deployment.yaml

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

- 查看 nginx-deployment 详细信息

    ```bash
    kubectl describe deployment nginx-deployment
    ```

    输出：

    ```
    Name:                   nginx-deployment
    Namespace:              default
    CreationTimestamp:      Tue, 28 May 2019 11:13:16 +0800
    Labels:                 <none>
    Annotations:            deployment.kubernetes.io/revision: 1
                            kubectl.kubernetes.io/last-applied-configuration:
                            {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"name":"nginx-deployment","namespace":"default"},"spec":{"replica...
    Selector:               app=nginx
    Replicas:               2 desired | 2 updated | 2 total | 2 available | 0 unavailable
    StrategyType:           RollingUpdate
    MinReadySeconds:        0
    RollingUpdateStrategy:  25% max unavailable, 25% max surge
    Pod Template:
    Labels:  app=nginx
    Containers:
    nginx:
        Image:        nginx:1.7.9
        Port:         80/TCP
        Host Port:    0/TCP
        Environment:  <none>
        Mounts:       <none>
    Volumes:        <none>
    Conditions:
    Type           Status  Reason
    ----           ------  ------
    Available      True    MinimumReplicasAvailable
    Progressing    True    NewReplicaSetAvailable
    OldReplicaSets:  <none>
    NewReplicaSet:   nginx-deployment-76bf4969df (2/2 replicas created)
    Events:
    Type    Reason             Age   From                   Message
    ----    ------             ----  ----                   -------
    Normal  ScalingReplicaSet  60s   deployment-controller  Scaled up replica set nginx-deployment-76bf4969df to 2
    ```

- 查看 nginx-deployment 创建的 pods

    ```bash
    kubectl get pods -l app=nginx
    ```

    输出

    ```
    NAME                                READY   STATUS    RESTARTS   AGE
    nginx-deployment-76bf4969df-65x5g   1/1     Running   0          21m
    nginx-deployment-76bf4969df-7prjk   1/1     Running   0          21m
    ```

## Updating the deployment

可以通过应用一个新的 YAML 文件，更新 nginx image 的版本到 1.8

- 更新 nginx-deployment.yaml

    ```yaml
    apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
    kind: Deployment
    metadata:
    name: nginx-deployment
    spec:
    selector:
        matchLabels:
        app: nginx
    replicas: 2
    template:
        metadata:
        labels:
            app: nginx
        spec:
        containers:
        - name: nginx
            image: nginx:1.8 # Update the version of nginx from 1.7.9 to 1.8
            ports:
            - containerPort: 80
    ```

- Apply 更新后的 YAML 文件

    ```bash
    kubectl apply -f nginx-deployment.yaml
    ```

    输出：

    ```
    deployment.apps/nginx-deployment configured
    ```

- 观察 pods 的更新

    ```bash
    kubectl get pods -l app=nginx
    ```

    输出：
    ```
    NAME                                READY   STATUS              RESTARTS   AGE
    nginx-deployment-5896fbb489-87s64   0/1     ContainerCreating   0          15s
    nginx-deployment-76bf4969df-65x5g   1/1     Running             0          26m
    nginx-deployment-76bf4969df-7prjk   1/1     Running             0          26m

    ---

    NAME                                READY   STATUS        RESTARTS   AGE
    nginx-deployment-5896fbb489-87s64   1/1     Running       0          88s
    nginx-deployment-5896fbb489-94p47   1/1     Running       0          57s
    nginx-deployment-76bf4969df-65x5g   1/1     Terminating   0          27m

    ```

## Scaling the application by increasing the replica count

可以通过更新 YAML 文件的 `replicas` 字段来弹性伸缩 nginx 的实例数

- 更新 nginx-deployment.yaml

    ```yaml
    apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
    kind: Deployment
    metadata:
    name: nginx-deployment
    spec:
    selector:
        matchLabels:
        app: nginx
    replicas: 4 # Update the replicas from 2 to 4
    template:
        metadata:
        labels:
            app: nginx
        spec:
        containers:
        - name: nginx
            image: nginx:1.8
            ports:
            - containerPort: 80
    ```

- Apply 更新后的 YAML 文件

    ```bash
    kubectl apply -f nginx-deployment.yaml
    ```

    输出：

    ```
    deployment.apps/nginx-deployment configured
    ```

- 验证 pods 数量

    ```bash
    kubectl get pods -l app=nginx
    ```

    输出：

    ```
    NAME                                READY   STATUS    RESTARTS   AGE
    nginx-deployment-5896fbb489-5z6m2   1/1     Running   0          4s
    nginx-deployment-5896fbb489-87s64   1/1     Running   0          7m10s
    nginx-deployment-5896fbb489-94p47   1/1     Running   0          6m39s
    nginx-deployment-5896fbb489-sflcz   1/1     Running   0          4s
    ```

## Deleting a deployment

```bash
kubectl delete deployment nginx-deployment
```


## Reference

[Run a Stateless Application Using a Deployment](https://kubernetes.io/docs/tasks/run-application/run-stateless-application-deployment/)

# 在 Kubernetes 上部署 Redis Cluster

本文展示如何用 Kubernetes 运行一个 redis 集群

<!-- vim-markdown-toc GFM -->

* [Objectives](#objectives)
* [Before you begin](#before-you-begin)
* [Redis Cluster Basics](#redis-cluster-basics)
    * [What is Redis?](#what-is-redis)
    * [Why Use Redis？](#why-use-redis)
    * [What is Redis Cluster?](#what-is-redis-cluster)
* [Deploying Redis Cluster in Kubernetes](#deploying-redis-cluster-in-kubernetes)
    * [Perrequisites](#perrequisites)
    * [Deploy Redis Cluster](#deploy-redis-cluster)
* [Testing the Redis Cluster](#testing-the-redis-cluster)
* [Reference](#reference)

<!-- vim-markdown-toc -->


## Objectives

- 创建一个 StatefulSet
- 使用`kubectl` 查看 redis 集群的详细信息
- 了解 redis, redis-cluster 基本概念

## Before you begin

在开始之前，需要有一个 kubernetes 集群和配置好和集群通讯的命令行工具 `kubectl`
。如果还没有集群，可以使用[Minikube](https://kubernetes.io/docs/setup/minikube/) 或 [Kubeadm](https://github.com/RosenLo/notes/blob/master/k8s/create_highly_available_clusters_with_kubeadm.md) 创建一个集群。
也可以直接使用官方提供的在线体验环境:

- [Katacoda](https://www.katacoda.com/courses/kubernetes/playground)
- [Play with Kubernetes ](https://labs.play-with-k8s.com/)

然后需要了解以下概念：

- [Pods](https://github.com/RosenLo/notes/blob/master/k8s/concepts.md#pods)
- [Cluster DNS](https://github.com/RosenLo/notes/blob/master/k8s/dns_pod_service.md)
- [Headless Services](https://github.com/RosenLo/notes/blob/master/k8s/concepts.md#headless-services)
- [PersistentVolume Provisioning](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)
- [StatefulSets](https://github.com/RosenLo/notes/blob/master/k8s/concepts.md#statefulset)
- [PodDisruptionBudgets](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#specifying-a-poddisruptionbudget)
- [PodAntiAffinity](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature)
- [Kubectl CLI](https://kubernetes.io/docs/reference/kubectl/kubectl/)


## Redis Cluster Basics

### What is Redis?

[Redis](https://redis.io/)
是一个开源的内存数据库，通常被当成数据库，缓存或消息中间件使用。它可以存储、操作多种高级数据类型如：`lists`, `maps`, `sets`, `sorted sets`。

由于 Redis 可以接受多种形式的 keys，所以可以在 server 上执行操作，从而减少 client 端的负载。

Redis 所有数据全在内存，只有在需要持久化时才会写到磁盘。


### Why Use Redis？

- 非常快，用 C 编写，移植性好可以运行在 POSIX 系统，例如 Linux、 Mac OS X
- 最受欢迎的 key/value 数据库，也是在容器世界中最受欢迎的 NoSQL 数据库
- 缓存方案减少了直接调用后端数据库的次数，同时也减轻了后端数据库的压力
- 丰富的 Client API Library
- 开源、稳定

### What is Redis Cluster?

[Redis Cluster](https://redis.io/topics/cluster-tutorial) 是一组 Redis
实例，设计用于通过分片来提高数据库伸缩能力，集群中 master 实例至少需要有一个 slave 副本（为了满足故障切换）负责管理一个区域 hash slot （哈希槽）。如果 master 实例变得不可用，对应的 slave 会成为 master。

master 节点分配哈希槽的范围为 `0-16383`，假设有三个 master 和三个 slave ，那 A
节点哈希槽范围：`0-5000`， B 节点哈希槽范围： `5001-10000`， C 节点哈希槽范围：
`10001-16383`

内部通讯使用 [gossip protocol](https://en.wikipedia.org/wiki/Gossip_protocol) 来传播集群信息、发现新节点

<div> <img src="../assets/redis-cluster-architecture.png"/> </div><br>

## Deploying Redis Cluster in Kubernetes

### Perrequisites

- PersistentVolumes (受限于资源，本次示例为 glusterfs 静态 PV，glusterfs 的 provider 是 [heketi](https://github.com/heketi/heketi))

redis-pv.yaml

```yaml
---
apiVersion: v1
kind: PersistentVolume
metadata:
name: redis-cluster1
spec:
capacity:
    storage: 5Gi
accessModes:
    - ReadWriteOnce
glusterfs:
    endpoints: "glusterfs-cluster"
    path: "redis_cluster1"
    readOnly: false
---
apiVersion: v1
kind: PersistentVolume
metadata:
name: redis-cluster2
spec:
capacity:
    storage: 5Gi
accessModes:
    - ReadWriteOnce
glusterfs:
    endpoints: "glusterfs-cluster"
    path: "redis_cluster2"
    readOnly: false
---
apiVersion: v1
kind: PersistentVolume
metadata:
name: redis-cluster3
spec:
capacity:
    storage: 5Gi
accessModes:
    - ReadWriteOnce
glusterfs:
    endpoints: "glusterfs-cluster"
    path: "redis_cluster3"
    readOnly: false
---
apiVersion: v1
kind: PersistentVolume
metadata:
name: redis-cluster4
spec:
capacity:
    storage: 5Gi
accessModes:
    - ReadWriteOnce
glusterfs:
    endpoints: "glusterfs-cluster"
    path: "redis_cluster4"
    readOnly: false
---
apiVersion: v1
kind: PersistentVolume
metadata:
name: redis-cluster5
spec:
capacity:
    storage: 5Gi
accessModes:
    - ReadWriteOnce
glusterfs:
    endpoints: "glusterfs-cluster"
    path: "redis_cluster5"
    readOnly: false
---
apiVersion: v1
kind: PersistentVolume
metadata:
name: redis-cluster6
spec:
capacity:
    storage: 5Gi
accessModes:
    - ReadWriteOnce
glusterfs:
    endpoints: "glusterfs-cluster"
    path: "redis_cluster6"
    readOnly: false
```

查看预先创建好的 pv
```
# kubectl get pv
NAME             CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS     CLAIM    STORAGECLASS   REASON   AGE
redis-cluster1   5Gi        RWO            Retain           UnBound                                     5m
redis-cluster2   5Gi        RWO            Retain           UnBound                                     5m
redis-cluster3   5Gi        RWO            Retain           UnBound                                     5m
redis-cluster4   5Gi        RWO            Retain           UnBound                                     5m
redis-cluster5   5Gi        RWO            Retain           UnBound                                     5m
redis-cluster6   5Gi        RWO            Retain           UnBound                                     5m
```


### Deploy Redis Cluster

- 使用 StatefulSet 部署 Redis

redis-sts.yaml

```yaml
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: redis-cluster
data:
  update-node.sh: |
    #!/bin/sh
    REDIS_NODES="/data/nodes.conf"
    sed -i -e "/myself/ s/[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}/${POD_IP}/" ${REDIS_NODES}
    exec "$@"
  redis.conf: |+
    cluster-enabled yes
    cluster-require-full-coverage no
    cluster-node-timeout 15000
    cluster-config-file /data/nodes.conf
    cluster-migration-barrier 1
    appendonly yes
    protected-mode no
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: redis-cluster
spec:
  serviceName: redis-cluster
  replicas: 6
  selector:
    matchLabels:
      app: redis-cluster
  template:
    metadata:
      labels:
        app: redis-cluster
    spec:
      containers:
      - name: redis
        image: redis:5.0.1-alpine
        ports:
        - containerPort: 6379
          name: client
        - containerPort: 16379
          name: gossip
        command: ["/conf/update-node.sh", "redis-server", "/conf/redis.conf"]
        env:
        - name: POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        volumeMounts:
        - name: conf
          mountPath: /conf
          readOnly: false
        - name: data
          mountPath: /data
          readOnly: false
      volumes:
      - name: conf
        configMap:
          name: redis-cluster
          defaultMode: 0755
  volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```

```bash
kubectl  create -f redis-sts.yaml
configmap/redis-cluster created
statefulset.apps/redis-cluster created

kubectl  create -f redis-svc.yaml
service/redis-cluster created
```

- 检查 Redis 节点是否运行

```bash
kubectl get pods
NAME                               READY   STATUS    RESTARTS   AGE
redis-cluster-0                    1/1     Running   0          3m
redis-cluster-1                    1/1     Running   0          2m
redis-cluster-2                    1/1     Running   0          2m
redis-cluster-3                    1/1     Running   0          1m
redis-cluster-4                    1/1     Running   0          1m
redis-cluster-5                    1/1     Running   0          1m
```

- 检查 pv 使用情况

```
NAME             CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                            STORAGECLASS   REASON   AGE
redis-cluster1   5Gi        RWO            Retain           Bound    default/data-redis-cluster-5                             11m
redis-cluster2   5Gi        RWO            Retain           Bound    default/data-redis-cluster-1                             11m
redis-cluster3   5Gi        RWO            Retain           Bound    default/data-redis-cluster-0                             11m
redis-cluster4   5Gi        RWO            Retain           Bound    default/data-redis-cluster-4                             11m
redis-cluster5   5Gi        RWO            Retain           Bound    default/data-redis-cluster-3                             11m
redis-cluster6   5Gi        RWO            Retain           Bound    default/data-redis-cluster-2                             11m
```

- 部署 Redis Cluster

一共六个节点，前面三个节点为 master, 后面三个节点为 slave，输入 `yes`
接受配置继续部署

```bash
kubectl exec -it redis-cluster-0 -- redis-cli --cluster create --cluster-replicas 1 $(kubectl get pods -l app=redis-cluster -o jsonpath='{range.items[*]}{.status.podIP}:6379 ')
```

输出：

```
>>> Performing hash slots allocation on 6 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
Adding replica 10.244.4.59:6379 to 10.244.4.58:6379
Adding replica 10.244.2.219:6379 to 10.244.1.38:6379
Adding replica 10.244.1.39:6379 to 10.244.2.218:6379
M: 12a608cc619df436e1f364cecc527e1b3f8e4d52 10.244.4.58:6379
   slots:[0-5460] (5461 slots) master
M: c3a8c08267f6fa00c58389669fa8060919b7f808 10.244.1.38:6379
   slots:[5461-10922] (5462 slots) master
M: 9fff6947d41e3e828907ef58dc2af112169ea587 10.244.2.218:6379
   slots:[10923-16383] (5461 slots) master
S: e7be513ed7a5a24a08e752f2475dd559ab8bb056 10.244.4.59:6379
   replicates 12a608cc619df436e1f364cecc527e1b3f8e4d52
S: 1eee241f4ffc938bf7aa605b21331c50e0fcff74 10.244.2.219:6379
   replicates c3a8c08267f6fa00c58389669fa8060919b7f808
S: e482d91e46064a53eaf48f046ab4dd44262b999b 10.244.1.39:6379
   replicates 9fff6947d41e3e828907ef58dc2af112169ea587
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join
.....
>>> Performing Cluster Check (using node 10.244.4.58:6379)
M: 12a608cc619df436e1f364cecc527e1b3f8e4d52 10.244.4.58:6379
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
M: c3a8c08267f6fa00c58389669fa8060919b7f808 10.244.1.38:6379
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: 1eee241f4ffc938bf7aa605b21331c50e0fcff74 10.244.2.219:6379
   slots: (0 slots) slave
   replicates c3a8c08267f6fa00c58389669fa8060919b7f808
S: e482d91e46064a53eaf48f046ab4dd44262b999b 10.244.1.39:6379
   slots: (0 slots) slave
   replicates 9fff6947d41e3e828907ef58dc2af112169ea587
M: 9fff6947d41e3e828907ef58dc2af112169ea587 10.244.2.218:6379
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: e7be513ed7a5a24a08e752f2475dd559ab8bb056 10.244.4.59:6379
   slots: (0 slots) slave
   replicates 12a608cc619df436e1f364cecc527e1b3f8e4d52
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
```

- 查看 redis 集群详情和每个成员的角色

```bash
kubectl exec -it redis-cluster-0 -- redis-cli cluster info
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:6
cluster_size:3
cluster_current_epoch:6
cluster_my_epoch:1
cluster_stats_messages_ping_sent:21
cluster_stats_messages_pong_sent:25
cluster_stats_messages_sent:46
cluster_stats_messages_ping_received:20
cluster_stats_messages_pong_received:21
cluster_stats_messages_meet_received:5
cluster_stats_messages_received:46
```

```bash
for x in $(seq 0 5); do echo "redis-cluster-$x"; kubectl exec redis-cluster-$x -- redis-cli role; echo; done
redis-cluster-0
master
248
10.244.4.77
6379
248

redis-cluster-1
master
1400
10.244.1.40
6379
1400

redis-cluster-2
master
1400
10.244.2.218
6379
1400

redis-cluster-3
slave
10.244.4.74
6379
connected
248

redis-cluster-4
slave
10.244.2.219
6379
connected
1400

redis-cluster-5
slave
10.244.1.39
6379
connected
1400
```


## Testing the Redis Cluster

部署一个简单的 flask 应用，然后删除一个 master 节点观察集群状态

- 部署一个计数器 App

app-deployment-service.yaml

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: hit-counter-lb
spec:
  type: NodePort
  ports:
  - port: 80
    protocol: TCP
    targetPort: 5000
    nodePort: 30020
  selector:
      app: myapp
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hit-counter-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: calinrus/api-redis-ha:1.0
        ports:
        - containerPort: 5000
```

- 连续访问

```bash
for i in `seq 1 100`; do curl http://127.0.0.1:30020; sleep 0.2; done
I have been hit 1 times since deployment.
I have been hit 2 times since deployment.
I have been hit 3 times since deployment.
I have been hit 4 times since deployment.
I have been hit 5 times since deployment.
I have been hit 6 times since deployment.
I have been hit 7 times since deployment.
```

- 删除一个 master 节点

删除节点会后自动拉起

```bash
kubectl delete pod/redis-cluster-0
```

- 查看集群节点细信息

可以发现 redis-cluster-0 已经失联，master 角色切换到 redis-cluster-3

```
for x in $(seq 0 5); do echo "redis-cluster-$x"; kubectl exec redis-cluster-$x -- cat /data/nodes.conf; echo; done
redis-cluster-0
error: unable to upgrade connection: container not found ("redis")

redis-cluster-1
e482d91e46064a53eaf48f046ab4dd44262b999b 10.244.1.39:6379@16379 slave 9fff6947d41e3e828907ef58dc2af112169ea587 0 1562056088000 6 connected
9fff6947d41e3e828907ef58dc2af112169ea587 10.244.2.218:6379@16379 master - 0 1562056090718 3 connected 10923-16383
c3a8c08267f6fa00c58389669fa8060919b7f808 10.244.1.40:6379@16379 myself,master - 0 1562056089000 2 connected 5461-10922
1eee241f4ffc938bf7aa605b21331c50e0fcff74 10.244.2.219:6379@16379 slave c3a8c08267f6fa00c58389669fa8060919b7f808 0 1562056088713 5 connected
12a608cc619df436e1f364cecc527e1b3f8e4d52 10.244.4.77:6379@16379 master,fail - 1562056073060 1562056071000 1 connected
e7be513ed7a5a24a08e752f2475dd559ab8bb056 10.244.4.74:6379@16379 master - 0 1562056089716 7 connected 0-5460
vars currentEpoch 7 lastVoteEpoch 7

redis-cluster-2
e482d91e46064a53eaf48f046ab4dd44262b999b 10.244.1.39:6379@16379 slave 9fff6947d41e3e828907ef58dc2af112169ea587 0 1562056088000 6 connected
e7be513ed7a5a24a08e752f2475dd559ab8bb056 10.244.4.74:6379@16379 master - 0 1562056089000 7 connected 0-5460
9fff6947d41e3e828907ef58dc2af112169ea587 10.244.2.218:6379@16379 myself,master - 0 1562056088000 3 connected 10923-16383
12a608cc619df436e1f364cecc527e1b3f8e4d52 10.244.4.77:6379@16379 master,fail - 1562056073066 1562056071000 1 connected
c3a8c08267f6fa00c58389669fa8060919b7f808 10.244.1.40:6379@16379 master - 0 1562056089935 2 connected 5461-10922
1eee241f4ffc938bf7aa605b21331c50e0fcff74 10.244.2.219:6379@16379 slave c3a8c08267f6fa00c58389669fa8060919b7f808 0 1562056088925 5 connected
vars currentEpoch 7 lastVoteEpoch 7

redis-cluster-3
e7be513ed7a5a24a08e752f2475dd559ab8bb056 10.244.4.74:6379@16379 myself,master - 0 1562056087000 7 connected 0-5460
12a608cc619df436e1f364cecc527e1b3f8e4d52 10.244.4.77:6379@16379 master,fail - 1562056073072 1562056072272 1 connected
1eee241f4ffc938bf7aa605b21331c50e0fcff74 10.244.2.219:6379@16379 slave c3a8c08267f6fa00c58389669fa8060919b7f808 0 1562056088311 5 connected
c3a8c08267f6fa00c58389669fa8060919b7f808 10.244.1.40:6379@16379 master - 0 1562056089315 2 connected 5461-10922
9fff6947d41e3e828907ef58dc2af112169ea587 10.244.2.218:6379@16379 master - 0 1562056087310 3 connected 10923-16383
e482d91e46064a53eaf48f046ab4dd44262b999b 10.244.1.39:6379@16379 slave 9fff6947d41e3e828907ef58dc2af112169ea587 0 1562056086308 6 connected
vars currentEpoch 7 lastVoteEpoch 0

redis-cluster-4
9fff6947d41e3e828907ef58dc2af112169ea587 10.244.2.218:6379@16379 master - 0 1562056089874 3 connected 10923-16383
c3a8c08267f6fa00c58389669fa8060919b7f808 10.244.1.40:6379@16379 master - 0 1562056089000 2 connected 5461-10922
1eee241f4ffc938bf7aa605b21331c50e0fcff74 10.244.2.219:6379@16379 myself,slave c3a8c08267f6fa00c58389669fa8060919b7f808 0 1562056088000 5 connected
e7be513ed7a5a24a08e752f2475dd559ab8bb056 10.244.4.74:6379@16379 master - 0 1562056087000 7 connected 0-5460
12a608cc619df436e1f364cecc527e1b3f8e4d52 10.244.4.77:6379@16379 master,fail - 1562056073114 1562056071000 1 connected
e482d91e46064a53eaf48f046ab4dd44262b999b 10.244.1.39:6379@16379 slave 9fff6947d41e3e828907ef58dc2af112169ea587 0 1562056088873 6 connected
vars currentEpoch 7 lastVoteEpoch 0

redis-cluster-5
c3a8c08267f6fa00c58389669fa8060919b7f808 10.244.1.40:6379@16379 master - 0 1562056089000 2 connected 5461-10922
9fff6947d41e3e828907ef58dc2af112169ea587 10.244.2.218:6379@16379 master - 0 1562056088808 3 connected 10923-16383
e482d91e46064a53eaf48f046ab4dd44262b999b 10.244.1.39:6379@16379 myself,slave 9fff6947d41e3e828907ef58dc2af112169ea587 0 1562056088000 6 connected
12a608cc619df436e1f364cecc527e1b3f8e4d52 10.244.4.77:6379@16379 master,fail - 1562056073043 1562056071739 1 connected
1eee241f4ffc938bf7aa605b21331c50e0fcff74 10.244.2.219:6379@16379 slave c3a8c08267f6fa00c58389669fa8060919b7f808 0 1562056089811 5 connected
e7be513ed7a5a24a08e752f2475dd559ab8bb056 10.244.4.74:6379@16379 master - 0 1562056087000 7 connected 0-5460
vars currentEpoch 7 lastVoteEpoch 0
```

- 应用读写情况

可以看到中间切换过程会有短暂不能写入（此时 salve 还是可读状态），但随着 master 切换后，即恢复了写入，数据也没有丢失。

```bash
for i in `seq 1 100`; do curl http://127.0.0.1:30020; sleep 0.2; done
I have been hit 1 times since deployment.
I have been hit 2 times since deployment.
I have been hit 3 times since deployment.
I have been hit 4 times since deployment.
I have been hit 5 times since deployment.
I have been hit 6 times since deployment.
I have been hit 7 times since deployment.



I have been hit 8 times since deployment.
I have been hit 9 times since deployment.
I have been hit 10 times since deployment.
I have been hit 11 times since deployment.
I have been hit 12 times since deployment.
I have been hit 13 times since deployment.
```

- 查看删除节点

```bash
kubectl exec redis-cluster-0 -- redis-cli role
slave
10.244.4.78
6379
connected
10538
```

发现节点 ip 改变了，那如何修复集群节点信息呢？

**注意**，我们创建了一个 `ConfigMap` 在容器启动时会调用 `/conf/update-node.sh`
，这个脚本会修改 redis nodes 配置文件的 ip 为新容器的
ip，然后集群恢复开始同步信息

## Reference

[Deploying Redis Cluster on Top of Kubernetes](https://rancher.com/blog/2019/deploying-redis-cluster/)
[Kuberntes Docs](https://kubernetes.io/docs/home/)
[Redis docs](https://redis.io/documentation)
[Redis cluster tutorial](https://redis.io/topics/cluster-tutorial)

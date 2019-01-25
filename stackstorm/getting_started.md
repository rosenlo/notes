# StackStorm Basic

---
G!-- vim-markdown-toc GFM -->

* [StackStorm 总览](#stackstorm-总览)
    * [StackStorm 介绍](#stackstorm-介绍)
    * [StackStorm 架构图](#stackstorm-架构图)
    * [StackStorm 如何工作：](#stackstorm-如何工作)
* [安装和配置](#安装和配置)
    * [Docker](#docker)
        * [环境](#环境)
        * [下载](#下载)
        * [生成配置文件](#生成配置文件)
        * [启动容器](#启动容器)
        * [销毁容器](#销毁容器)
* [StackStorm HA Components](#stackstorm-ha-components)
    * [st2client](#st2client)
    * [st2web](#st2web)
    * [st2auth](#st2auth)
    * [st2api](#st2api)
    * [st2stream](#st2stream)
    * [st2rulesengine](#st2rulesengine)
    * [st2timersengine](#st2timersengine)
    * [st2workflowengine](#st2workflowengine)
    * [st2notifier](#st2notifier)
    * [st2sensorcontainer](#st2sensorcontainer)
* [参考](#参考)

<!-- vim-markdown-toc -->
---


## StackStorm 总览

### StackStorm 介绍

StackStorm 是一个跨服务和工具集成的自动化平台.它使你现有的基础设施和应用环境联系在一起，你可以更容易的自动化环境.它特别注重采用行动作应对事件。

StackStorm 有助于自动化常用的操作模式，一些例子：

- 促进故障排除。当系统故障被`Nagios`, `Sensu`, `New Relic`和其他监控系统捕获时触发，在物理节点、虚拟节点和应用组件上运行一系列诊断检查，然后将结果发送到交流工具上，例如`HipChat`和`JIRA`。
- 自动修复。识别和验证 OpenStack 计算节点的硬件故障，正确的隔离节点并邮件通知管理员潜在的停机时间，如果出现任何问题，停止这个工作流然后调用`PagerDuty`通知管理员。
- 持续部署。用`Jenkins`编译和测试，提供一个新的AWS集群。在负载均衡器上打开一些流量，根据`NewRelic`应用运行数据发布或回滚，

StackStorm 有助于你将这些和其他操作模式组合为规则和工作流或动作。

### StackStorm 架构图

<div> <img src="../assets/st2-architecture.jpg" /> </div><br>


### StackStorm 如何工作：

StackStorm 工作步骤大致如下：

1. `Sensors`聚合各种服务 push 或 pull 的事件
2. `Rules Engine`对事件进行规制匹配，如果匹配`Triggers`生成`Actions`写入 RabbitMQ
3. `Workers` 处理RabbitMQ中的`Actions`
4. `Actions` 调用各种服务运行`Workflow`
5. 记录日志和审计信息写入 MongoDB
6. 处理结果返回`Rules Engine`进行后续处理


StackStorm 包含了传感器和动作的可扩展的适配器集通过插件的形式融入环境。

- **Sensors** 是用于入站和出站集成的 python 插件，分别接收和监听事件。当一个外部系统产生的事件被`Sensors`处理将会通知`Triggers`发送到 StackStorm 系统。
- **Trigger** 代表外部系统产生的事件。有通用的触发器（例如：timers, webhooks）和集成的（例如：Sensu alert, JIRA issue updated）。也可以写一个`Sensor plugin`定义新的触发器类型。
- **Actions** 是处理事件的执行方式，一般由外部系统执行，有通用的动作（SSH， REST call）和集成的（OpenStack, Docker, Puppet）或一些自定义的动作，`Actions`可以是一个python插件或任意的脚本，通过新增几行`metadata`即可添加到StackStorm。`Actions`可以直接通过`CLI`或`API`来调用，或作为`rules`和`workflows`的一部分来使用。
- **Rules** 映射`Triggers`到`actions`或`workflows`，当事件触发后，通过Rule定义的标准(criteria)进行匹配，当匹配成功则映射`trigger`数据到`action`输入。
- **Workflows** 将一系列的`Action`组合成工作流（workflow），定义`Action`执行的顺序、转换条件和传递数据。即把在`Action`中定义的各种原子操作组合成一个复杂的任务进行一个自动化操作，也可以通过手动调用或`Rule`触发。
- **Packs** 是一个内容部署单位。它通过对集成（trigger和actions）和自动化（rules和workflows）进行分组简化了 StackStorm 可插拔内容的管理和分享。我们可以创建自己的`packs`分享到Github或提交到[StackStorm Exchange](https://exchange.stackstorm.org/)。
- **Audit trail** 记录并存储`Action`执行的详细细节。同时集成外部日志分析工具:`LogStash`，`Splunk`, `Statsd`，`Syslog`。

## 安装和配置


StackStorm 提供了`RPMs`、`Debs`和`Docker images`三种方式来部署使用，这里我们使用`docker-compose`。
`stackstorm/stackstorm`镜像预装了`st2`，`st2web`，`st2mistral`和`st2chatops`包。

### Docker

#### 环境

- macOs 10.13.6
- Docker for Mac 18.03.1-ce
- docker-compose version 1.21.1, build 5a3f1a3

#### 下载

```bash
git clone https://github.com/stackstorm/st2-docker
cd st2-docker
```

#### 生成配置文件
只需运行一次

```
make env
```

在当前路径下生成`conf/`目录，存储了依赖的配置文件，例如：mongo, redis, postgres, rabbitmq，默认无需修改。

```bash
ls -l conf
total 40
-rw-r--r--  1 Rosen  staff   71 Nov  5 17:54 mongo.env
-rw-r--r--  1 Rosen  staff  132 Nov  5 17:54 postgres.env
-rw-r--r--  1 Rosen  staff  117 Nov  5 17:54 rabbitmq.env
-rw-r--r--  1 Rosen  staff   73 Nov  5 17:54 redis.env
-rw-r--r--  1 Rosen  staff   40 Nov  5 17:54 stackstorm.env
```

#### 启动容器
第一次启动会从 Docker Hub 下载 images

```bash
docker-compose up -d
```

#### 销毁容器

```bash
docker-compose down
```

## StackStorm HA Components

### st2client

命令行客户端，一个与所有 st2 容器资源共享的容器，用来切换到已存在的 StackStorm
集群并执行命令行

```bash
# obtain st2client pod name
ST2CLIENT=$(kubectl get pod -l app=st2client -o jsonpath="{.items[0].metadata.name}")

# run a single st2 client command
kubectl exec -it ${ST2CLIENT} -- st2 --version

# switch into a container shell and use st2 CLI
kubectl exec -it ${ST2CLIENT} /bin/bash
```

### st2web

- StackStorm Web UI admin Dashboard
- k8s 配置包含了一个 `Pod Deployment` 服务(提供了 `2` 个副本，支持 HA )
- 代理请求到 `st2auth`, `st2api`, `st2strem`
- st2web 默认使用 `NodePort` 服务，如果需要配置放到公共网络，需要配置
`LoadBalancer` 或 其他 Proxy

### st2auth

- st2auth 管理所有的认证
- k8s 配置包含了一个 `Pod Deployment` 服务(提供了 `2` 个副本，支持 HA ) 和监听在 `9100` 端口的 `ClusterIP` 服务
- 多个 st2auth 进程可以配置在 load balancer 用 active-active 模式

### st2api

- 提供了 REST API 服务处理来自 Web UI, CLI, ChatOps 和其他 st2 组件的请求
- k8s 配置包含了 `Pod Deployment` 服务(提供了 `2` 个副本，支持 HA ) 和监听在 `9101` 端口的 `ClusterIP` 服务
- 这是 StackStorm 服务中最重要的服务之一，强烈建议增加副本数，提高处理能力

### st2stream

- 暴露了一个 server-sent 事件流， 一般由 WebUI, ChatOps 等客户端使用接收来自
st2stream 服务发送的更新信息
- k8s 配置包含了 `Pod Deployment` (提供了 `2` 个副本，支持 HA ) 和监听在 `9102` 端口的 `ClusterIP` 服务

### st2rulesengine

- 接收 triggers 进行 rule 判断是否执行 action
- k8s 配置包含了 `Pod Deployment` (提供了 `2` 个副本，支持 HA )

### st2timersengine

- 计时器引擎，本质是一个定时任务
- k8s 配置包含了 `Pod Deployment` (只提供了 `1` 个副本，不支持 HA )
- 不能工作在 active-active 模式(多个计时器会处理多个重复的 events )
- 依赖 k8s failover/reschedule 能力解决处理失败的情况

### st2workflowengine

- 驱动 `orquesta workflows` 和 `shedules action` 通过 `st2actionrunner`
  组件执行 workflows
- 多个 processes 可以运行在 active-active 模式
- k8s 配置包含了 `Pod Deployment` (提供了 `2` 个副本，支持 HA )
- 空闲时 workflow engine process 共享负载和 work

### st2notifier

- 多个 processes 可以运行在 active-active 模式
- 使用 RabbitMQ, MongoDB 根据 action 执行完成生成 triggers 以及根据 action 重新调度
- k8s 配置包含了 `Pod Deployment` (提供了 `2` 个副本，支持 HA )
- 依赖 etcd 调度

### st2sensorcontainer

- 管理 StackStorm sensors: start, stops and restarts 作为他的子进程
- k8s 配置包含了 `Pod Deployment` (只提供了 `1` 个副本，不支持 HA)
- 官方未来会重构这个组件，依赖 k8s failover/reshedule 机制分布到多个 pod
  来提高计算能力。 more details: [single-sensor-per-container mode #4179](https://github.com/StackStorm/st2/pull/4179)

### st2actionrunner

- 实际执行 actions 的worker
- k8s 配置包含了一个 `Pod Deployment` 服务(默认提供了 `5` 个副本，可以继续增加，提高处理能力)
- 依赖 etcd 调度

### st2scheduler

- 处理外部入口的执行 action 请求
- k8s 配置包含了 `Pod Deployment` (提供了 `2` 个副本，可以继续增加，提高调度能力)
- 依赖数据库版本控制来调度

### st2garbagecollector

- 根据配置清理陈旧的执行信息和其他操作的数据
- k8s 配置包含了 `Pod Deployment` (提供了 `1` 个副本，定期执行，1个副本足够了)
- 默认没有配置，所以需要配置才会执行。清理陈旧的数据会显著提高集群的处理能力，所以在生产环境强烈建议配置垃圾回收

### RabbitMQ HA Cluster

- StackStorm 内部进程通信和负载分配
- Helm Chart 默认高可用部署，默认3个节点

### MongoDB HA ReplicaSet

- 作为 StackStorm 数据库
- Helm Chart 默认高可用部署，默认三个节点(一主，两从)

### etcd

- 作为 StackStorm 分布式调度后端
- 因为 StackStorm 现在对 etcd 的依赖还不是很重，只需要部署一个节点，官方未来会计划在 Helm Chart 部署三个节点 [Issue](https://github.com/StackStorm/stackstorm-ha/issues/8)


## 参考

- [StackStorm Overview](https://docs.stackstorm.com/overview.html)
- [Install on Docker](https://docs.stackstorm.com/install/docker.html)
- [k8s HA Components](https://docs.stackstorm.com/latest/install/k8s_ha.html#components)

# StackStorm Basic

---
<!-- vim-markdown-toc GFM -->

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

<div> <img src="../../assets/st2-architecture.jpg" /> </div><br>


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

## 参考

- [StackStorm Overview](https://docs.stackstorm.com/overview.html)
- [Install on Docker](https://docs.stackstorm.com/install/docker.html)

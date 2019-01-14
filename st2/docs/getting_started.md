# StackStorm

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
* [基本组成](#基本组成)
    * [Sensors](#sensors)
    * [Triggers](#triggers)
    * [Rules](#rules)
        * [Rule Structure](#rule-structure)
    * [Timers](#timers)
    * [Workflows](#workflows)
        * [Orquesta](#orquesta)
            * [Orquesta 工作模式](#orquesta-工作模式)
            * [Orquesta Model Definition](#orquesta-model-definition)
            * [Expressions and Context](#expressions-and-context)
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


## 基本组成

### Sensors

- 将外部系统与 StackStorm 内部 event 集成，支持 push/pull
- 对 event 进行 rule 匹配触发 trigger
- Sensor 通常会注册一个 trigger，但不是必需。例如：webhook

### Triggers

- 识别入站的 event
- 由 type(string) 和可选的 parameters(object) 组成 tuple
- 通常与 Rules 组合

### Rules

- 映射 tiggers 到 actions/workflow
- 匹配 criteria (payload 字段) map 到 action input

#### Rule Structure

```yaml
---
    name: "rule_name"                      # required
    pack: "examples"                       # optional, 默认是`default`
    description: "Rule description."       # optional
    enabled: true                          # required

    trigger:                               # required, 由 sensors 触发
        type: "trigger_type_ref"

    criteria:                              # optional, trigger 字段，用来匹配
        trigger.payload_parameter_name1:
            type: "regex"
            pattern : "^value$"
        trigger.payload_parameter_name2:
            type: "iequals"
            pattern : "watchevent"

    action:                                # required, 匹配通过后执行 action
        ref: "action_ref"
        parameters:                        # optional
            foo: "bar"
            baz: "{{ trigger.payload_parameter_1 }}"
```


### Timers

- Interval
- Cron: 类似于 Linux 中的 Crontab，不过更加灵活
- DateTime: 在特定的时间执行


### Workflows
StackStorm 支持两种类型：ActionChain 和 Mistral

- ActionChain: 内置 Workflow runner。 比较简单
- Mistral: 专用的 Workflow 服务，源于 OpenStack。可以编写更复杂的 workflow
- Orquesta: 专为 StackStorm 设计的新一代 workflow engine，目前处于 beta 测试阶段，未来 Action Chain 和 Mistral 都将被替换


#### Orquesta
graph based wrkflow engine designed，Orquesta 具有以下特点：

- 顺序执行
- forks, joins
- 复杂的查询和转换
- 多个根节点，可以同时开始 task
- 多个 branchs 并行 task，然后汇聚到一个 branch，一个 branch 可以分发到多个 branchs
- 工作流追溯


##### Orquesta 工作模式

- conductor
    - 指导 workflow 走向
    - 追踪执行信息
    - 本身不执行 action，指导 provider 执行 action
    - 更新 workflow 状态
    - 管理更新历史和运行上下文
    - 确定整个 workflow 的状态和结果

- provider
    - 执行 action
    - 执行结果信息返回给 conductor


##### Orquesta Model Definition 

- Workflow Model

Attribute   |Required   |Accept Expressions |Description
------------|-----------|-------------------|----------
version     |Yes	    |No                 |The version of the spec being used in this workflow DSL.
description	|No         |No                 |The description of the workflow.
input	    |No         |Yes, see With Items|A list of input arguments for this workflow.
vars	    |No	        |Yes                |A list of variables defined for the scope of this workflow.
tasks	    |Yes        |Yes, see Task      |A dictionary of tasks that defines the intent of this workflow.
output	    |No      	|Yes, at each item  |A list of variables defined as output for the workflow.

- Task Model

Attribute	|Required	|Accept Expresssions |Description
------------|-----------|--------------------|------------
delay	    |No	        |Yes                 |If specified, the number of seconds to delay the task execution.
join	    |No	        |No                  |If specified, sets up a barrier for a group of parallel branches.
with	    |No	        |Yes, see With Items |When given a list, execute the action for each item.
action	    |No	        |Yes                 |The fully qualified name of the action to be executed.
input	    |No	        |Yes, entire dict or at each item|A dictionary of input arguments for the action execution.
next	    |No	        |See Task Transition |Define what happens after this task is completed.

- With Items Model

Attribute	|Required	|Accept Expresssions |Description
------------|-----------|--------------------|-----------
items	    |Yes	    |Yes                 |The list of items to execute the action with.
concurrency	|No	        |Yes                 |The number of items being processed concurrently.

- Task Transition Model

Attribute	|Required	|Accept Expresssions |Description
------------|-----------|--------------------|-----------
when	    |No	        |Yes                 |The criteria defined as an expression required for transition.
publish	    |No	        |Yes, at each item   |A list of key value pairs to be published into the context.
do	        |No	        |No                  |A next set of tasks to invoke when transition criteria is met.

- Engine Commands

Command	|Description
--------|----------------
noop	|No operation or do not execute anything else.
fail	|Fails the workflow execution.


- Example

```yaml
version: 1.0

description: Calculates (a + b) * (c + d)

input:
  - a: 0    # Defaults to value of 0 if input is not provided.
  - b: 0
  - c: 0
  - d: 0
  - messages
  - hosts
  - commands
  - member
  - message

tasks:
  task1:
    # Fully qualified name (pack.name) for the action.
    action: math.add

    # Assign input arguments to the action from the context.
    input:
      operand1: <% ctx(a) %>
      operand2: <% ctx(b) %>

    # Specify what to run next after the task is completed.
    next:
      - # Specify the condition in YAQL or Jinja that is required
        # for this task to transition to the next set of tasks.
        when: <% succeeded() %>

        # Publish variables on task transition. This allows for
        # variables to be published based on the task state and
        # its result.
        publish:
          - msg: task1 done
          - ab: <% result() %>

        # List the tasks to run next. Each task will be invoked
        # sequentially. If more than one tasks transition to the
        # same task and a join is specified at the subsequent
        # task (i.e task1 and task2 transition to task3 in this
        # case), then the subsequent task becomes a barrier and
        # will be invoked when condition of prior tasks are met.
        do:
          - log
          - task3

  task2:
    # Short hand is supported for input arguments. Arguments can be
    # delimited either by space, comma, or semicolon.
    action: math.add operand1=<% ctx("c") %> operand2=<% ctx("d") %>
    next:
      - when: <% succeeded() %>

        # Short hand is supported for publishing variables. Variables
        # can be delimited either by space, comma, or semicolon.
        publish: msg="task2 done", cd=<% result() %>

        # Short hand with comma delimited list is supported.
        do: log, task3

  task3:
    # Join is specified for this task. This task will be invoked
    # when the condition of all inbound task transitions are met.
    join: all
    action: math.multiple operand1=<% ctx('ab') %> operand2=<% ctx('cd') %>
    next:
      - when: <% succeeded() %>
        publish: msg="task3 done" abcd=<% result() %>
        do: log

  task4:
    # The item value can be named. For multiple lists of items.
    # The value returned from item() in this case would be a dictionary like {"message": "value"}.
    # No concurrency, short hand notation
    with: message in <% ctx(messages) %>
    action: core.echo message=<% item(message) %>

  task5:
    # For multiple lists of items. the lists need zipped first with the zip funcion,
    # and then define the keys required to access the individual values in each item.
    with: host, command in <% ctx(hosts), ctx(commands) %>
    action: core.remote host=<% item(host) %> cmd=<% item(command) %>

  task6:
    # When concurrency is required, use the formal schema.
    with:
      items: <% ctx(messages) %>
      concurrency: 2
    action: core.echo message=<% item() %>

  task7:
    action: slack.post member=<% ctx(member) %> message=<% ctx(message) %>
    next:
      - when: <% succeeded() %>
        publish: msg="Sccussfully posted message."
        do: log
      - when: <% failed() %>
        publish: msg="Unable to post message due to error: <% result() %>"
        do: notify_on_error


  # Define a reusable task to log progress. Although this task is
  # referenced by multiple tasks, since there is no join defined,
  # this task is not a barrier and will be invoked separately.
  log:
    action: core.log message=<% ctx(msg) %>

  # A workflow example that illustrates error handling. By default
  # when any task fails, the notify_on_error task will be executed
  # and the workflow will transition to the failed state.
  notify_on_error:
    action: core.log message=<% ctx(msg) %>
    next:
      # The fail specified here tells the workflow to go into
      # failed state on completion of the notify_on_error task.
      - do: fail

output:
  - result: <% ctx().abcd %>
```

##### Expressions and Context

- Type
    - [YAQL](https://yaql.readthedocs.io/en/latest/)
    - [Jinja](http://jinja.pocoo.org/)

- Workflow Runtime Context

在运行时，workflow 会维护一个字典结构的上下文管理器。`input`, `vars`
在初始化时分配， `publish`, `output` 在每个 task 完成时分配。

一旦分配了参数，可以通过`ctx` 函数访问：
- 参数访问: `ctx(foobar)`
- `.`符号访问：`ctx().foobar`




## 参考

- [Overview](https://docs.stackstorm.com/overview.html)
- [Install on Docker](https://docs.stackstorm.com/install/docker.html)

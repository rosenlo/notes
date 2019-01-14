# StackStorm

---
<!-- vim-markdown-toc GFM -->

* [StackStorm Component](#stackstorm-component)
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
            * [YAQL](#yaql)
            * [Jinja](#jinja)
            * [Workflow Operations](#workflow-operations)
* [参考](#参考)

<!-- vim-markdown-toc -->
---


## StackStorm Component

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

- Assignment Order

    顺序分配，从上至下

- Assignment Scope

    有并行分支执行时，上下文作用域在每个分支，并在有`join`时合并。如同一个变量，后写入的将会覆盖前值。

    Example:

    在下面的例子中，当 `task1` 完成，写入 `x=123` 到上下文管理器， `task2` 延迟 3 秒后覆盖写入 `x=789`

    ```yaml
    version: 1.0

    vars:
    - x

    tasks:
    # Branch 1
    task1:
        action: core.noop
        next:
        - publish: x=123
            do: task4

    # Branch 2
    task2:
        action: core.sleep delay=3
        next:
        - do: task3
    task3:
        action: core.noop
        next:
        - publish: x=789
            do: task4

    # Converge branch 1 and 2
    task4:
        join: all
        action: core.noop
    ```

##### YAQL

YAQL(Yet Another Query Language) 是 OpenStack
下的一个项目，用于复杂的数据查询和传输，在 Workflow 中定义 YAQL 表达式：<% YAQL expression %>

- Dictionaries

    - 创建一个字典，使用`dict`函数。例如：`<% dict(a=>123, b=>true) %>`, return `{"a":
123, "b": True}`
    - 获取所有 keys ：`<% ctx.(dict1).keys() %>`, return `["a", "b"]`
    - 获取所有 values ：`<% ctx.(dict1).values() %>`, return `[123, true]`
    - 连接两个字典：`<% dict(a=>123, b=>true) + dict(c=abc) %>`, return `{"a":
      123, "b": True, "c": "abc"}`
    - 获取key： `<% ctx(dict1).get(a) %>`, return `True`, 也可以给出替代值
      `<% ctx(dict1).get(d, false) %>` return `False`

- List

    - 创建一个列表，使用`list`函数。例如： `<% list(1,2,3) %>` return `[1, 2, 3]`
    - 连接两个列表：`<% list(1, 2, 3) %> + <% list(a, b, c) %>`, return `[1, 2,
      3, "a", "b", "c"]`
    - 访问列表元素，通过下标：`<% ctx(list1)[0] %>`, return `1`

- Queries

    ```json
    {
        "vms": [
            {
                "name": "vmweb1",
                "region": "us-east",
                "role": "web"
            },
            {
                "name": "vmdb1",
                "region": "us-east",
                "role": "db"
            },
            {
                "name": "vmweb2",
                "region": "us-west",
                "role": "web"
            },
            {
                "name": "vmdb2",
                "region": "us-west",
                "role": "db"
            }
        ]
    }
    ```

    - `<% ctx(vms).select(%.name) %>` 返回一个列表，包含了所有 name `['vmweb1',
    'vmdb1', 'vmweb2', 'vmdb2']`
    - `<% ctx(vms).select(%.name, %.role) %>` 返回一个列表，包含了所有 name 和
    role `[['vmweb1', 'web'], ['vmdb1', 'db'], ['vmweb2', 'web'], ['vmdb2', 'db']]`
    - `<% ctx(vms).select(%.region).distinct() %>`返回一个列表：包含了不同的 region
    `['us-east', 'us-west']`
    - `<% ctx(vms).where(%.region = 'us-east').select(%.name) %>`
    返回一个列表：筛选出 region 在 `us-east` 的机器 ['vmweb1', 'vmdb1']
    - `<% ctx(vms).where(%.region = 'us-east' and %.role = 'web').select(%.name) %>` 返回一个列表： 筛选出 web 服务在 us-east 的机器 ['vmweb1']
    - `<% let(my_region => 'us-east', my_role => 'web') -> ctx(vms).where(%.region
    = %.my_region and %.role = %.my_role).select(%.name) %>` 赋值变量的一种写法


- List to Dictionary

    表达式：`<% dict(vms=>dict(ctx(vms).select([ctx(name), ctx()]))) %>`

    返回：
    ```json
    {
        "vms": {
            "vmweb1": {
                "name": "vmweb1",
                "region": "us-east",
                "role": "web"
            },
            "vmdb1": {
                "name": "vmdb1",
                "region": "us-east",
                "role": "db"
            },
            "vmweb2": {
                "name": "vmweb2",
                "region": "us-west",
                "role": "web"
            },
            "vmdb2": {
                "name": "vmdb2",
                "region": "us-west",
                "role": "db"
            }
        }
    }
    ```


- Built-in Functions

    - [YAQL 标准库](https://yaql.readthedocs.io/en/latest/standard_library.html)

- StackStorm Functions

    - `st2kv('system.shared_key_x')` 返回 key 为`shared_key_x` 在 system
      作用域的值，需要注意的是 key 的名字需要加引号，不然 YAQL 会当做 YAQL
      语法运行。除此之外，返回的值还可以被加密，通过参数`decrypt`设置为`true`
      例如：`st2kv('st2_key_id', decrypt=true)`


##### Jinja

Jinja 表达式：`{{ Jinja expression }}`, 代码块：`{% %}`

`{`和`}`与 JSON 冲突，所以整个 Jinja 表达式与封装需为单引号或双引号


- Built-in Filters

    - [内置 Filters](http://jinja.pocoo.org/docs/2.10/templates/#list-of-builtin-filters)

- StackStorm Filters

    - 与 StackStorm Function 一致


##### Workflow Operations

- Pausing: `st2 execution pause <execution-id>`
    - 必须在 Workflow 状态为`running` 时才能 pause
    - 当前正在执行的 task 不会停止，下一个 task 将会暂停
    - 如果 pasue 操作是从子流程发出，将不会影响父流程的进行
    - 状态从 running -> pausing -> paused

- Resuming: `st2 exection resume <execution-id>`
    - 必须在 Workflow 状态为 `paused` 时才能 resume
    - 状态从 paused -> running

- Canceling: `st2 execution cancel <execution-id>`
    - 必须在 Workflow 状态为`running` 时才能 cancel
    - 当前正在执行的 task 不会停止，下一个 task 将会取消
    - 状态从 running -> canceling -> canceled

- Re-running: `st2 execution re-run <execution-id>`
    - 相当于重新执行一遍 Workflow， 未来将支持 Re-run 特定 task


## 参考

- [StackStorm Doc](https://docs.stackstorm.com/latest/index.html)
- [Install on Docker](https://docs.stackstorm.com/install/docker.html)

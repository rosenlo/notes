# StackStorm Advanced

---
<!-- vim-markdown-toc GFM -->

* [StackStorm Component](#stackstorm-component)
    * [Actions](#actions)
        * [Action Runners](#action-runners)
            * [Available Runners](#available-runners)
        * [Writing Custom Actions](#writing-custom-actions)
            * [Action Metadata](#action-metadata)
            * [Parameters in Actions](#parameters-in-actions)
            * [Action Registration](#action-registration)
        * [Overriding Runner Parameters](#overriding-runner-parameters)
        * [Environment Variables Available to Actions](#environment-variables-available-to-actions)
        * [Converting Existing Script to Actions](#converting-existing-script-to-actions)
        * [Writing Custom Python Actions](#writing-custom-python-actions)
            * [Configuration File](#configuration-file)
            * [Logging](#logging)
            * [Action Service](#action-service)
    * [Sensors](#sensors)
    * [Triggers](#triggers)
    * [Rules](#rules)
        * [Rule Structure](#rule-structure)
    * [Timers](#timers)
    * [Workflows](#workflows)
        * [Orquesta](#orquesta)
            * [Orquesta Role](#orquesta-role)
            * [Orquesta Model Definition](#orquesta-model-definition)
            * [Expressions and Context](#expressions-and-context)
            * [YAQL](#yaql)
            * [Jinja](#jinja)
            * [Workflow Operations](#workflow-operations)
    * [Pack](#pack)
        * [What is a Pack?](#what-is-a-pack)
        * [Managing Packs](#managing-packs)
        * [Discovering Packs](#discovering-packs)
        * [Installing a Pack](#installing-a-pack)
        * [Create and Contribute a Pack](#create-and-contribute-a-pack)
            * [Anatomy of a Pack](#anatomy-of-a-pack)
                * [Actions](#actions-1)
                * [Rule](#rule)
                * [Sensors](#sensors-1)
                * [Aliases](#aliases)
                * [Policies](#policies)
            * [Creating Your First Pack](#creating-your-first-pack)
        * [Pack Configuration](#pack-configuration)
            * [Basic Concepts and Terminology](#basic-concepts-and-terminology)
                * [Configuration Schema](#configuration-schema)
                * [Configuration File](#configuration-file-1)
                * [Static Configuration Values](#static-configuration-values)
                * [Dynamic Configuration Value (DCV)](#dynamic-configuration-value-dcv)
            * [Limitations](#limitations)
                * [Dynamic Config Values](#dynamic-config-values)
                * [User Context](#user-context)
    * [Datastore](#datastore)
* [Reference](#reference)

<!-- vim-markdown-toc -->
---


## StackStorm Component

### Actions

- action 是可以任意执行的代码片段，可以用各种语言编写
- 当被规则条件所匹配，通过 tigger 触发
- 多个 action 可以组成 workflow
- 也可以直接通过 CLI, API, UI 执行

#### Action Runners

- action runner 是一个用户执行 action 的运行环境
- StackStorm 预装了很多 action runner， 为了让用户专注于 action
  的实现，而不是底层运行环境


##### Available Runners

1. `local-shell-cmd` - 在 StackStorm 本机运行 Linux 命令
2. `local-shell-script` - 在 StackStorm 本机运行 script
3. `remote-shell-cmd` - 在一个或多个机器上运行 Linux 命令
4. `remote-shell-script` - 在一个或多个机器上运行 script
5. `python-script` - 这是一个 Python runner. 需要写一个 Python classes 继承StackStorm Base classes 实现 `run()` 方法. 在 StackStorm 本机运行. `run()` 方法返回一个数组，包含成功状态和结果对象，更多信息参考 [Action Runners](#action-runners)
6. `http-request` - HTTP client，可以发送 HTTP 请求
7. `action-chain` - 内置 workflow runner, 支持执行一些简单的 workflow。更多信息参考 [ActionChain](https://docs.stackstorm.com/latest/actionchain.html)
8. `mistral-v2` - OpenStack 内置的 workflow，支持复杂的 workflow。更多信息参考 [Mistral](https://docs.stackstorm.com/latest/mistral.html)
9. `cloudslang` - 也是一个workflow runner， 在 v2.9 将被移除
10. `inquirer` -  注意: 这个 runner 为了`core.ask` action 而实现的, 在其他 case 不应该引用

#### Writing Custom Actions

- 一个 YAML 元数据文件，包含 action, input
- 一个实现了 action 逻辑的脚本文件

action 脚本可以用任意语言实现，只要符合以下规则：

- 脚本成功时退出码为 `0`，失败时非 `0` （例如：1）
- 所有日志信息需要以标准错误输出，即 `stderr`

##### Action Metadata

- `name` - string
- `runner_type` - string, action 的 runner 类型
- `enabled` - bool
- `entry_point` - string, action 脚本的相对路径
- `parameters` - 字典结构，metadata 结构遵循 JSON Schema，如果 metadata 提供，input args 将会校验，否则跳过

    - `parameter` - 参数名
        - `type` - string, 参数类型
        - `description` - string, 参数描述
        - `required` - bool，是否必须
        - `position` - int, 相对位置
        - `default` - string, 默认值
        - `secret` - bool, 如果为 ture ，这个参数的值在 StackStorm 服务的 log
          中将被屏蔽
        - `immutable` - bool, 定义这个参数的默认值是否可以被覆盖

- `tags` - list, 一个带有标签的数组结构，用来提供补充信息

For example：

这是个 python runner(python-script) 。在 `send_sms.py` 文件中实现了 `run`
方法，接受三个参数(`from_number`, `to_number`, `body`)，它和 metadata 文件处在同一个位置。

```yaml
---
name: "send_sms"
runner_type: "python-script"
description: "This sends an SMS using twilio."
enabled: true
entry_point: "send_sms.py"
parameters:
    from_number:
        type: "string"
        description: "Your twilio 'from' number in E.164 format. Example +14151234567."
        required: true
        position: 0
        default: "{{config_context.from_number}}"
    to_number:
        type: "string"
        description: "Recipient number in E.164 format. Example +14151234567."
        required: true
        position: 1
        secret: true
    body:
        type: "string"
        description: "Body of the message."
        required: true
        position: 2
        default: "Hello {% if system.user %} {{ st2kv.system.user }} {% else %} dude {% endif %}!"
```

##### Parameters in Actions

- 模板文件中通过 `st2kv.system` 访问 parameters
- 在执行中通过 `action_context` 访问 paramaters
- 还可以通过 `config_context` 访问 [pack configuration](#pack-configuration)
- 在ActionChains 和 [Workflow](#workflow) 中，每一个 task 可以访问父
  `execution_id` 例如：

    ```yaml
    ...
    -
    name: "c2"
    ref: "core.local"
    parameters:
        cmd: "echo \"c2: parent exec is {{action_context.parent.execution_id}}.\""
    on-success: "c3"
    on-failure: "c4"
    ...
    ```

##### Action Registration

- 放在对应目录`/opt/stackstorm/packs/${pack_name}/`
- registry `st2 action create my_action_metadata.yaml`
- reload `st2ctl reload --register-actions`


#### Overriding Runner Parameters

Runner 的参数可以被覆盖，在一些场景如：需要自定义和简化操作

For example：

下面的 `linux.rsync` action cmd 参数被覆盖了，还传入了其他参数变量


```yaml
---
    name: 'rsync'
    runner_type: 'remote-shell-cmd'
    description: 'Copy file(s) from one place to another w/ rsync'
    enabled: true
    entry_point: ''
    parameters:
        source:
            type: 'string'
            description: 'List of files/directories to to be copied'
            required: true
        dest_server:
            type: 'string'
            description: "Destination server for rsync'd files"
            required: true
        destination:
            type: 'string'
            description: 'Destination of files/directories on target server'
            required: true
        cmd:
            immutable: true
            default: 'rsync {{args}} {{source}} {{dest_server}}:{{destination}}'
        connect_timeout:
            type: 'integer'
            description: 'SSH connect timeout in seconds'
            default: 30
        args:
            description: 'Command line arguments passed to rysnc'
            default: '-avz -e "ssh -o ConnectTimeout={{connect_timeout}}"'
```

不是所有的 runner 参数都可以被覆盖，以下是可以覆盖的参数：

- default
- description
- enum
- immutable
- required

#### Environment Variables Available to Actions

默认，`local`, `remote`, `python_runner` 可以使用以下 env 变量：

- `ST2_ACTION_PACK_NAME` - 现在的 PACK 名
- `ST2_ACTION_EXECUTION_ID` - 现在的 Execution ID
- `ST2_ACTION_API_URL` - 完整的API URL
- `ST2_ACTION_AUTH_TOKEN` - 在 running 中可用的认证 TOKEN，当 tasks 完成时销毁

For example:

```bash
#!/usr/bin/env bash

# Retrieve a list of actions by hitting the API using cURL and the information provided
# via environment variables

RESULT=$(curl -H "X-Auth-Token: ${ST2_ACTION_AUTH_TOKEN}" ${ST2_ACTION_API_URL}/actions)
echo ${RESULT}
```

#### Converting Existing Script to Actions

**Note: 如果没有参数可以跳过这步**

如果你已编写独立的程序或脚本，不限语言，可以通过下面的方法转换为 action
，非常简单。

1. 确保脚本符合约定
确保脚本带有状态码退出，成功以`0`状态码，失败以非 `0` 状态码， 如 `1`

2. 创建 metadata 文件参考 [Action Metadata](#action-metadata)

3. 更新脚本中的参数解析


- `named` - 参数不是位置参数
- `positional` - 位置参数

命名参数通过以下格式传递：

```
script.sh --param1=value --param2=value --param3=value
```

默认情况下，参数以两个破折号开始 `--` ，如果你想用一个破折号，可以用 `kwarg_op`
在 metadata 中定义，参考下面示例：

```yaml
---
name: "my_script"
runner_type: "remote-shell-script"
description: "Script which prints arguments to stdout."
enabled: true
entry_point: "script.sh"
parameters:
    key1:
        type: "string"
        required: true
    key2:
        type: "string"
        required: true
    key3:
        type: "string"
        required: true
    kwarg_op:
        type: "string"
        immutable: true
        default: "-"
```

```
script.sh -key1=value1 -key2=value2 -key3=value3
```

位置参数通过以下格式传递：

```
script.sh value2 value1 value3
```

如果只用位置参数，只需要在 metadata 文件中声明 parameters - `position`
属性，序列化基于以下规则：

- `string`, `integer`, `float` - 序列化为 string
- `boolean` - 序列化为 string 1 (true) or 0 (false)
- `array` - 序列化为逗号分隔的string (e.g. `foo,bar,baz`)
- `object` - 序列化为 JSON

For exmaple:

```yaml
---
name: "my_script"
runner_type: "remote-shell-script"
description: "Script which prints arguments to stdout."
enabled: true
entry_point: "script.sh"
parameters:
    key1:
        type: "string"
        required: true
        position: 0
    key2:
        type: "string"
        required: false
        position: 1
    key3:
        type: "string"
        required: true
        position: 3
```

```
script.sh value1 value2 value3
```

如果第二个位置参数是可选的，可以传一个空字符

```
script.sh value1 "" value3
```


#### Writing Custom Python Actions

python action 其实是一个继承 `st2common.runners.base_action.Action`
并重写 `run` 方法的模块。

看一个最小形式的示例：

Metadata 文件(`my_echo_action.yaml`):

```yaml
---
name: "echo_action"
runner_type: "python-script"
description: "Print message to standard output."
enabled: true
entry_point: "my_echo_action.py"
parameters:
    message:
        type: "string"
        description: "Message to print."
        required: true
        position: 0
```

Action 文件(`my_echo_action.py`):

```python
import sys

from st2common.runners.base_action import Action

class MyEchoAction(Action):
    def run(self, message):
        print(message)

        if message == 'working':
            return (True, message)
        return (False, message)
```

return 有两种方式：

- 如果正常结束没有引发异常，将被视为成功，返回的对象可以是任意类型
- 返回一个元组，指定状态和返回的对象

For exmaple:

- `return False` 表示执行状态成功，result 对象是`False`
- `return (False, "Failed!")` 表示执行状态失败，result 对象是 `"Failed!"`


##### Configuration File

用于存放静态配置，文件命名规范：`<pack_name>.yaml`

更多信息参考 [Pack Configuration](#pack-configuration)

##### Logging

这个 logger 是标准 python logger 来自 `logging`模块

For example：

```python
def run(self):
    ...
    success = call_some_method()

    if success:
        self.logger.info('Action successfully completed')
    else:
        self.logger.error('Action failed...')
```

##### Action Service

类似于 sensors， `action_service`
提供了一个全局（整个 workflow ）对象，可以用来在不同 task 之间做数据传递等等

For exmaple:

```python
def run(self):
  data = {'somedata': 'foobar'}

  # Add a value to the datastore
  self.action_service.set_value(name='cache', value=json.dumps(data))

  # Retrieve a value
  value = self.action_service.get_value('cache')
  retrieved_data = json.loads(value)

  # Retrieve an encrypted value
  value = self.action_service.get_value('ma_password', decrypt=True)
  retrieved_data = json.loads(value)

  # Delete a value
  self.action_service.delete_value('cache')
```


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


##### Orquesta Role

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


- For example

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

    For example:

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
下的一个项目，用于复杂的数据查询和传输，在 Workflow 中定义 YAQL 表达式：`<% YAQL expression %>`

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


### Pack

#### What is a Pack?

- pack 是扩展 StackStorm 的集成和自动化的部署单位
- 通常 pack 是根据符合和产品划定边界，例如 AWS, Docker, Ansible etc.
- pack 可以包含 [Actions](#actions), [Workflows](#workflows), [Rules](#rules), [Sensors](#sensors), [Aliases](#action-aliases)
- StackStorm 的内容始终是 pack 的一部分， 所以了解怎么去创建一个 pack 和如何工作非常重要
- 在 [StackStorm Exchange](https://exchange.stackstorm.org/) 上可以找到公共 pack

#### Managing Packs

管理 pack 主要通过命令: `st2 pack <...>`, 更多帮助 `st2 pack -h`

`list` 和 `get` 是主要查看本地 pack 信息的命令:

```bash
# List all installed packs
st2 pack list

# Get detailed information about an installed pack
st2 pack get core
```

#### Discovering Packs

搜索公共包通过命令 `st2 pack search packname`

#### Installing a Pack

默认安装会从 [StackStorm Exchange on GitHub](https://github.com/StackStorm-Exchange) 上下载 packs 到本地 `/opt/stackstorm/packs` 然后注册到 StackStorm

```bash
# Install from the Exchange by pack name
st2 pack install sensu

# You can also install multiple packs:
st2 pack install datadog github
```

安装指定 git 仓库 pack

```bash
# Install your own pack from git
st2 pack install https://github.com/emedvedev/chatops_tutorial
```

默认安装 latest 版本，如果需要特定版本，通过 `=` 指定 ，例如：

```bash
# Fetch a specific commit
st2 pack install cloudflare=776b9a4

# Or a version tag
st2 pack install cloudflare=0.1.0

# Or a branch
st2 pack install https://github.com/emedvedev/chatops_tutorial=testing
```

安装本地 pack

```bash
# Install a pack from '/tmp/bitcoin' dir
st2 pack install file:///tmp/bitcoin
```

运行 `st2 pack install` 在已安装的 pack 上会**替换**或更新版本到 latest, 如果没有指定版本

配置文件不会被覆盖，可以很轻松回滚，在生产环境建议用 latest，避免错过大变更

#### Create and Contribute a Pack

pack 有一个定义的结构体，创建一个新 pack 需要遵循这个结构，在 debugging 问题的时候有帮助

##### Anatomy of a Pack

一个典型的 pack 文件层级如下所示:

```bash
# contents of a pack folder
actions/                 #
rules/                   #
sensors/                 #
aliases/                 #
policies/                #
tests/                   #
etc/                     # any additional things (e.g code generators, scripts...)
config.schema.yaml       # configuration schema
packname.yaml.example    # example of config, used in CI
pack.yaml                # pack definition file
requirements.txt         # requirements for Python packs
requirements-tests.txt   # requirements for python tests
icon.png                 # 64x64 .png icon
```

最上面的是几个主要的文件目录 `actions`, `rules`, `sensors` 和 `policies`
还有一些分享文件:

- `pack.yaml` - Metadata 文件描述和定义了文件夹作为 pack
- `config.schema.yaml` - Schema 定义了 pack 使用的配置元素
- `requirements.txt` - python 依赖库，在 pack 安装时会自动安装依赖

更多配置 Schema 参考 [Pack Configuration](#pack-configuration)

###### Actions

```bash
# contents of actions/
actions/
   lib/
   action1.yaml
   action1.py
   action2.yaml
   action1.sh
   workflow1.yaml
   workflow2.yaml
   workflows/
     workflow1.yaml
     workflow2.yaml
```

- `actions` 文件夹包含了 action 脚本和 action metadata 文件
- 把 workflow 配置文件放在不同路径是好的做法
- 注意 `lib` 子文件夹通常用来存放公共 Python 代码，用来被 pack actions 使用

###### Rule

```bash
# contents of rules/
rules/
   rule1.yaml
   rule2.yaml
```

- `rules` 文件夹包含了 rules
- 有关如何写 rules 细节参考 [Rules](#rules)

###### Sensors

```bash
# contents of sensors/
sensors/
   common/
   sensor1.py
   sensor1.yaml
   sensor2.py
   sensor2.yaml
```

- `sensors` 文件夹包含了 sensors
- 有关如何写 sensors 细节参考 [Sensors](#sensors)

###### Aliases

```bash
 contents of aliases/
aliases/
   alias1.yaml
   alias2.yaml
```

- `aliases` 文件夹包含了 aliases
- 有关如何写 aliases 细节参考 [Aliases](#action-aliases)

###### Policies

```bash
# contents of policies/
policies/
   policy1.yaml
   policy2.yaml
```

- `policies` 文件夹包含了 Policies
- 有关如何写 policies 细节参考 [Policies](#policies)


##### Creating Your First Pack

下面的例子，我们会创建一个简单的 pack 叫做 **hello_st2**

1. 创建 pack 目录结构和相关文件

    ```bash
    # Use the name of the pack for the folder name.
    mkdir hello_st2
    cd hello_st2
    mkdir actions
    mkdir rules
    mkdir sensors
    mkdir aliases
    mkdir policies
    touch pack.yaml
    touch requirements.txt
    ```

    **Note**: 所有文件夹都是可选，只要保持空文件夹就行，唯一必须要求创建 `config.schema.yaml`文件，空 schema 文件是无效的

2. 创建 pack metadata 文件 `pack.yaml`:

    ```yaml
    ---
    ref: hello_st2
    name: Hello StackStorm
    description: Simple pack containing examples of sensor, rule, and action.
    keywords:
        - example
        - test
    version: 0.1.0
    python_versions:
    - "2"
    - "3"
    author: StackStorm, Inc.
    email: info@stackstorm.com
    ```
    **Note**: 在 metadata 中会强制运行以下规则：
    - `version` 值必须符合 [Senmantic Versioning](https://semver.org/): `0.2.5`, 而不是 `0.2`
    - `name` 值只允许包含字母，数字和下划线，除非你明确设置 `ref`
    - `email` 必须包含合法格式的邮箱地址

3. 创建 [action](#actions)，由 metadata 和 entrypoint 组成

    看一个简单的例子：

    ```yaml
    ---
    name: greet
    pack: hello_st2
    runner_type: "local-shell-script"
    description: Greet StackStorm!
    enabled: true
    entry_point: greet.sh
    parameters:
        greeting:
            type: string
            description: Greeting you want to say to StackStorm (i.e. Hello, Hi, Yo, etc.)
            required: true
            position: 1
    ```

    创建 `entry_point` 脚本：

    ```bash
    #!/bin/bash
    echo "$1, StackStorm!"
    ```

4. 创建 [sensor](#sensors)

    下面这个简单的 sensor 每隔60秒发出一个 event 到 StackStorm

    `sensors/sensor1.yaml`:

    ```yaml
    ---
    class_name: "HelloSensor"
    entry_point: "sensor1.py"
    description: "Test sensor that emits triggers."
    trigger_types:
    -
        name: "event1"
        description: "An example trigger."
        payload_schema:
        type: "object"
    ```

    `sensors/sensor1.py`:

    ```python
    import eventlet

    from st2reactor.sensor.base import Sensor


    class HelloSensor(Sensor):
        def __init__(self, sensor_service, config):
            super(HelloSensor, self).__init__(sensor_service=sensor_service, config=config)
            self._logger = self.sensor_service.get_logger(name=self.__class__.__name__)
            self._stop = False

        def setup(self):
            pass

        def run(self):
            while not self._stop:
                self._logger.debug('HelloSensor dispatching trigger...')
                count = self.sensor_service.get_value('hello_st2.count') or 0
                payload = {'greeting': 'Yo, StackStorm!', 'count': int(count) + 1}
                self.sensor_service.dispatch(trigger='hello_st2.event1', payload=payload)
                self.sensor_service.set_value('hello_st2.count', payload['count'])
                eventlet.sleep(60)

        def cleanup(self):
            self._stop = True

        # Methods required for programmable sensors.
        def add_trigger(self, trigger):
            pass

        def update_trigger(self, trigger):
            pass

        def remove_trigger(self, trigger):
            pass
    ```

5. 创建 [rule](#rules)

    下面的 rule 示例由 sensor 接受 event 通过 trigger 调用上面定义的 action

    `rules/rule1.yaml`:

    ```yaml
    ---
    name: on_hello_event1
    pack: hello_st2
    description: Sample rule firing on hello_st2.event1.
    enabled: true
    trigger:
        type: hello_st2.event1
    action:
        ref: hello_st2.greet
        parameters:
            greeting: Yo
    ```

6. 创建 [action alias](#action-aliases)

    `aliases/alias1.yaml`:

    ```yaml
    ---
    name: greet
    pack: hello_st2
    description: "Greet StackStorm"
    action_ref: "hello_st2.greet"
    formats:
    - "greet {{greeting}}"
    ```

7. 创建 [policy](#policy)

    下面的 policy 示例限制了 `greet` action 的并发数

    `policies/policy1.yaml`:

    ```yaml
    ---
    name: greet.concurrency
    pack: hello_st2
    description: Limits the concurrent executions of the greet action.
    enabled: true
    resource_ref: hello_st2.greet
    policy_type: action.concurrency
    parameters:
        threshold: 10
    ```

8. 安装 pack

    - 使用 git 和 `pack install` （推荐）

    ```bash
    # Get the code under git
    cd hello_st2
    git init && git add ./* && git commit -m "Initial commit"
    # Install from local git repo
    st2 pack install file:///$PWD
    ```

    如果本地有更新再次运行 `st2 pack install file:///$PWD`

    如果是更新在 GitHub Repo 运行 `st2 pack install https://github.com/MY/PACK`

    - 覆盖和重载

    ```bash
    mv ./hello_st2 /opt/stackstorm/packs
    st2ctl reload
    ```

#### Pack Configuration

pack 可以使用配置文件共享公共配置，例如：API
认证、连接信息、限制和阈值。这些配置在 `actions`, `sensors` 运行时可以使用

pack configuration 和 action parameters 不一样的是 pack 配置对 pack
所有资源可用，且很少更改。而 action param
调用动态传参，参数易变。例如：可能来自 rule 映射的 input event

pack configuration 遵循 infrastructure as code 理念，以 YAML 格式的文件存放在
`/opt/stackstorm/configs` 目录下，每个 pack 都需要定义自己的 schema
configuration 文件

##### Basic Concepts and Terminology

###### Configuration Schema

这个文件叫做 `config.schema.yaml` ，位于 `/opt/stackstorm/packs/<mypack>/`
目录下

```yaml
---
  api_key:
    description: "API key"
    type: "string"
    required: true
  api_secret:
    description: "API secret"
    type: "string"
    secret: true
    required: true
  region:
    description: "API region to use"
    type: "string"
    required: true
    default: "us-east-1"
  private_key_path:
    description: "Path to the private key file to use"
    type: "string"
    required: false
```

Note: `api_secret` 被标记为 `secret`，代表如果这个值被动态使用将会加密存储在数据库里

除了上面的扁平配置，schemas 还支持嵌套对象：

```yaml
---
  consumer_key:
    description: "Your consumer key."
    type: "string"
    required: true
    secret: true
  consumer_secret:
    description: "Your consumer secret."
    type: "string"
    required: true
    secret: true
  access_token:
    description: "Your access token."
    type: "string"
    required: true
    secret: true
  access_token_secret:
    description: "Your access token secret."
    type: "string"
    required: true
    secret: true
  sensor:
    description: "Sensor specific settings."
    type: "object"
    required: false
    additionalProperties: false
    properties:
      device_uids:
        type: "array"
        description: "A list of device UIDs to poll metrics for."
        items:
          type: "string"
        required: false
```


###### Configuration File

这是个 YAML 格式的文件，可以包含**静态**或**动态**的值。
命名规范是 `<pacn name>.yaml` 位于 `/opt/stackstorm/configs/`
目录，文件所属应该为 `st2:st2`

For example: pack `libcloud`, 位于 `/opt/stackstorm/configs/libcloud.yaml`

```yaml
---
  api_key: "some_api_key"
  api_secret: "{{st2kv.user.api_secret}}"  # user-scoped configuration value which is also a secret as declared in config schema
  region: "us-west-1"
  private_key_path: "{{st2kv.system.private_key_path}}"  # global datastore value
```

- Configuration 文件在运行时不会读动态配置
- 静态配置需要注册加载到 DB 通过运行 `st2ctl reload/st2-register-content`
  脚本，对 configs ，需要附加 `--register-configs` flag
- 当注册加载 configs 时会验证**静态值**的有效性，而**动态值**是 `jinja` 语法存在 DB ，在运行时解析，所以在注册加载阶段不会被验证有效性

###### Static Configuration Values

静态值从配置文件加载然后直接使用

###### Dynamic Configuration Value (DCV)

**Note**: 现在只支持 `strings` 字符串类型的DCV

DCV 是一个 `jinja` 模板表达式，在运行时解析为 [Datastore](#datastore) name(key)，然后把 datastore value 作为 configuration value

DCV 提供了额外的灵活性，还支持用户范围的 datastore
values，当你想不同用户调用 action 使用不同的配置时这非常有用

DCV 参考:

```yaml
---
  api_secret: "{{st2kv.user.api_secret}}"  # user-scoped configuration value which is also a secret as declared in config schema
  private_key_path: "{{st2kv.system.private_key_path}}"  # global datastore value
```

`api_secret` 是用户范围的DCV，意味着 `user` 将会被替换成执行这个 action 的 username

DCV 存储在 datastore 使用 CLI 和 API 配置

如果 value 在 schema 被标记为 secret ，设置 value 需要附加 `--encrypt` flag
，将会被加密存储在 datastore

```bash
st2 key set api_secret "my super secret api secret" --scope=user --encrypt
```

在上面的例子， `private_key_path` 是个常规 DCV ，常规 DCV
可以被管理员和任意用户使用

```bash
st2 key set private_key_path "/home/myuser/.ssh/my_private_rsa_key"
```

##### Limitations

Dynamic Config Values 的上下文有一些限制需要注意

###### Dynamic Config Values

- 目前仅支持 strings 字符串类型。这是为了保持功能简单和与已存在的 datastore
  操作兼容
- 如果你想使用非 string 类型的 value，可以 JSON 序列化存储在 datastore，然后在
  action/sensor 代码中反序列化

###### User Context

user context 只能通过 StackStorm API 触发的 actions 可用。

这意味着 `{{st2kv.user.some_value}}` 表达式仅在正确的 user 通过 StackStorm API
触发的 action 才会解析

如果是被 rule 触发， user context 则不可用 `{{st2kv.user}}` 会被解析为系统用户（默认: `stanley`）



### Datastore

## Reference

- [StackStorm Doc](https://docs.stackstorm.com/latest/index.html)

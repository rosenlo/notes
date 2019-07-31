


# Dockerfile 详解

## CMD

`CMD` 指令有三种形式：

- `CMD ["executable", "param1", "param2"]` （首选 exec 形式）
- `CMD ["param1", "param2"]` （作为默认的参数传递给 [ENTRYPOINT](#ENTRYPOINT)）
- `CMD command param1 param2` （shell 形式）

Dockerfile 中只允许有一个 CMD， 如果有多个，则最后一个生效。

**CMD 的主要目的是为了提供一个默认指令运行容器**。

> Note: 如果 `CMD` 用来提供默认参数给 `ENTRYPOINT` 使用，`CMD` 和 `ENTRYPOINT`
> 都必须是 JSON 数组格式。

> Note: exec 形式是解析成 JSON 数组，意味着需要使用双引号而不是单引号。

> Note: 不同于 shell 形式，exec 形式不会调用 command shell。意味着 shell
> 命令不会正常处理。举个例子： `CMD ["echo", "$HOME"]` 不会替换变量
> `$HOME`。如果希望执行 shell ，可以使用 shell 形式或 exec 形式直接执行 shell。举个例子：
> `CMD ["sh", "-c", "echo $HOME"]`

当使用 shell 或 exec 形式时 `CMD` 会在镜像运行时执行

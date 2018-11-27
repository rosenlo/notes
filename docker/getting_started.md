
<!-- vim-markdown-toc GFM -->

* [Docker](#docker)
    * [Configure Networking](#configure-networking)
        * [Network Drivers](#network-drivers)
            * [Host Networking](#host-networking)
* [Docker Compose](#docker-compose)
    * [Compose File Configuration Reference](#compose-file-configuration-reference)
        * [Network Configuration Reference](#network-configuration-reference)
            * [driver](#driver)
            * [external](#external)
            * [internal](#internal)
            * [name](#name)

<!-- vim-markdown-toc -->







## Docker

### Configure Networking

#### Network Drivers

以下为常用的几个网络驱动

- bridge: 默认网络驱动，通常用于多个独立容器之间的相互通信。
- host: 移除了容器和Docker主机之间的网络隔离，让容器直接使用Docker主机网络。仅在`Docker 17.06+` 和`swarm service可用`，详细:  [Host Networking](#host-networking)
- overlay: 让多个`Docker daemons`和`swarm services`互联，也可以让`swarm service`和独立容器或在两个`Docker daemons`上的独立容器互联，这个策略避免了在两个容器之间执行系统层级的路由。
- macvlan: 分配MAC地址给容器，让它在当前网络上显示为一个物理设备。Docker daemon通过MAC地址将流量路由到容器。


##### Host Networking
如果用`host`作容器网络的驱动，那容器网络不会与`Docker host`隔离。也就是说：如果你运行一个容器绑定了80端口，那么绑定的是主机的80端口。

这个`host`网络驱动模式只能工作在Linux hosts，不支持Docker for Mac、Docker for windows、Dcoker EE for windows server


## Docker Compose

### Compose File Configuration Reference

`Compose file`是用来管理和运行多个Docker容器应用。这可以让我们配置和管理app和依赖的多个服务，还可以自定义网络、服务、存储卷。通过`docker-compose up -d`一条简单的命令就可以创建和启动app及依赖的所有服务。非常方便开发测试以及部署。

默认文件名是`docker-compose.yml`，可以定义多个`Compose file` 通过`docker-compose -f xxx.yml` 指定。

以下配置皆为 **version 2+**

#### Network Configuration Reference

`Compose`默认会为你的app创建一个单独的网络: `app_default`，服务的每个一个容器都会加入到这个默认网络，容器间可以通过主机名或定义在`Compose file`的容器名相互访问。

如果想自定义网络配置可以通过文件顶层参数`networks`定义。

##### driver

- 单机下默认为`bridge`
- `Swarm`下默认为`overlay`，并配置为`attachable`，不能被修改。意味着其他独立的容器可以连接到`overlay`网络

##### external
```
version: '2.1'
networks:
  my-app-net:
    external: true
```
如果设置为`true`，指定这个网络被其他Compose文件所创建，如果不存在则会抛出异常。

##### internal

Docker创建的网络默认为`bridge`模式，如果设置为`true`，则只能内部访问，其他容器就访问不到了。

##### name
自定义网络名，也可以和其他参数一起使用如`external`。

```
version: '2.1'
networks:
  network1:
    external: true
    name: my-app-net
```


<!-- vim-markdown-toc GFM -->

* [Docker](#docker)
    * [Configure Networking](#configure-networking)
        * [Network Drivers](#network-drivers)
            * [Host Networking](#host-networking)
* [Docker Compose](#docker-compose)
    * [Compose File Configuration Reference](#compose-file-configuration-reference)
        * [links](#links)
        * [networks](#networks)
            * [aliases](#aliases)
        * [depends_on](#depends_on)
        * [build](#build)
            * [extra_hosts](#extra_hosts)
        * [Network Configuration Reference](#network-configuration-reference)
            * [driver](#driver)
            * [external](#external)
            * [internal](#internal)
            * [name](#name)
    * [Extend Services In Compose](#extend-services-in-compose)
        * [Multiple Compose files](#multiple-compose-files)
        * [Extending services](#extending-services)
    * [Volume Plugin](#volume-plugin)
        * [GlusterFS Volume Plugin](#glusterfs-volume-plugin)
* [Reference](#reference)

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

#### links
连接两个容器中的服务，语法（service:alias）或直接service name

容器就可以通过service name去访问其他容器了。前提容器之间处于相通的网络。

官方推荐使用[networks](#networks)

```yaml
web:
  links:
   - "db"
   - "db:database"
   - "redis"
```

#### networks
用于多个容器加入到一个网络

```yaml
services:
  some-service:
    networks:
     - some-network
     - other-network
```

##### aliases
如果有容器名冲突，可以用不同的别名(alias)代替。

一个服务可以在不同网络上有不同别名。

```yaml
version: '2'

services:
  web:
    build: ./web
    networks:
      - new

  worker:
    build: ./worker
    networks:
      - legacy

  db:
    image: mysql
    networks:
      new:
        aliases:
          - database
      legacy:
        aliases:
          - mysql

networks:
  new:
  legacy:
```

#### depends_on
- `docker-compose up`启动多个服务时，会先创建启动依赖的服务然后再启动
- `docker-compose up SERVICE`会自动拉起`SERVICE`依赖的服务。在下面的例子中，`docker-compose up web`会先创建启动`db`和`redis`然后再创建启动`web`

例：
```yaml
version: '2'
services:
  web:
    build: .
    depends_on:
      - db
      - redis
  redis:
    image: redis
  db:
    image: postgres
```

**Note**: `depends_on`不会等`db`和`redis`成`ready`状态才去启动`web`，在`started`状态时就会启动，如果需要等依赖服务`ready`，看这个[解决方案](https://docs.docker.com/compose/startup-order/)

#### build

##### extra_hosts
在build时添加hostname映射，存放于容器中的`/etc/hosts`文件。也可以在`docker --add-host`指定。

```
extra_hosts:
 - "somehost:162.242.195.82"
 - "otherhost:50.31.209.229"
```

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

### Extend Services In Compose

Compose支持两种方式复用公共配置：
- 多个Compose文件
- 用`extends`字段扩展

#### Multiple Compose files
主要用于区分不同的环境变量如：dev、test、qa、prod等环境或不同的workflow。

默认Compose会读两个文件：`docker-compose.yml`和`docker-compose.override.yml`。一般`docker-compose.yml`包含了基础配置，而`override`文件看名字就知道是用于覆盖基础配置或整个service。

多个文件用`-f`指定，Compose会根据指定文件的顺序合并最终的配置，从左至右。需要注意的是指定所有的`override`文件是基于`docker-compose.yml`的相对路径。

看一个官方例子：不同环境

**docker-compose.yml**

```yaml
web:
  image: example/my_web_app:latest
  links:
    - db
    - cache

db:
  image: postgres:latest

cache:
  image: redis:latest
```

**docker-compose.override.yml**

```yaml
web:
  build: .
  volumes:
    - '.:/code'
  ports:
    - 8883:80
  environment:
    DEBUG: 'true'

db:
  command: '-d'
  ports:
    - 5432:5432

cache:
  ports:
    - 6379:6379
```

当运行`docker-compose up`时会自动覆盖

再创建一个生产环境override文件。一般可能储存在不同`git repo`和项目代码分离开

**docker-compose.prod.yml**

```yaml
web:
  ports:
    - 80:80
  environment:
    PRODUCTION: 'true'

cache:
  environment:
    TTL: '500'
```

生产环境部署Compose files用以下命令：

```bash
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

如果新增了一个admin服务，也依赖基础Compose文件也可以通过以下命令运行：

**docker-compose.admin.yml**

```yaml
dbadmin:
  build: database_admin/
  links:
    - db
```

```bash
docker-compose -f docker-compose.yml -f docker-compose.admin.yml \
    run dbadmin
```

#### Extending services

**Note**: 目前`extends`关键字最高支持`version 2.1`，还不支持`version 3.x`  [detail](https://github.com/moby/moby/issues/31101)

相比[Multiple Compose files](#multiple-compose-files)，不仅可以在不同Compose files复用公共配置，还可以在不同项目中复用。

`links`、`volumes_from`和`depends_on`永远不要用`extends`去共用。避免隐式依赖造成未知异常。

看个官方例子加深理解一下：

可以看出是基于**service**覆盖，在web service下指定`common-services.yml`文件的webapp service

**docker-compose.yml**

```yaml
web:
  extends:
    file: common-services.yml
    service: webapp
```

**common-services.yml**

```yaml
webapp:
  build: .
  ports:
    - "8000:8000"
  volumes:
    - "/data"
```

还可以在同一个文件基于存在的service重新定义一个新的service

**docker-compose.yml**

```yaml
web:
  extends:
    file: common-services.yml
    service: webapp
  environment:
    - DEBUG=1
  cpu_shares: 5

important_web:
  extends: web
  cpu_shares: 10
````

### Volume Plugin

#### GlusterFS Volume Plugin
```bash
docker plugin install trajano/glusterfs-volume-plugin
docker volume create -d trajano/glusterfs-volume-plugin --opt servers=serverIP  volume/subPath
```

## Reference
- [Docker Docs](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/overview/)
- [Glusterfs Volume Plugin](https://github.com/trajano/docker-volume-plugins/tree/master/glusterfs-volume-plugin)

# Harbor

---

* [Harbor总览](#Harbor总览)
    * [Harbor介绍](#Harbor介绍)
    * [Harbor功能](#Harbor功能)
* [安装和配置](#安装和配置)
    * [环境](#环境)
    * [下载](#下载)
    * [配置](#配置)
        * [更新harbor.cfg](#更新harbor.cfg)
        * [添加域名解析](#添加域名解析)
        * [配置HTTPS](#配置HTTPS)
        * [更新docker-compose.cfg](#更新docker-compose.cfg)
    * [安装](#安装)
        * [启动容器](#启动容器)
        * [销毁容器](#销毁容器)
        * [故障排除](#故障排除)
* [参考](#参考)

---


## Harbor总览

### Harbor介绍

Harbor是一个开源的用于存储和分发Docker镜像的企业级Registry项目。Harbor通过添加一些企业级需要的功能特性如（安全性，身份和管理）扩展了Docker官方的Distribution项目。Harbor提升了用户构建和运行环境传输镜像的效率。同时支持多Registry节点之间的image复制，另外提供了更强的安全功能，如（用户管理，访问控制，活动审计）。

### Harbor功能

- 私有云Registry：同时支持容器镜像和`Helm chart`，`Harbor`作为Registry为私有云环境(如：容器运行和编排平台)提供服务。
- 基于角色的访问控制：用户和资源库通过项目组合起来，用户可以在同个项目的镜像有不同的权限。
- 基于策略的镜像复制：镜像可以基于多个过滤器（资源库，标记，标签）策略的在多个注册实例上复制（同步）, `Harbor`在遇到失败的时候自动重试复制。非常适合负载均衡，高可用，多数据中心，混合云场景。
- 漏洞扫描：`Harbor`定期扫描镜像漏洞并通知用户。
- LDAP/AD支持
- 镜像删除和垃圾回收
- 公证：确保镜像的真实性
- 用户面板
- 审计：所有操作可回溯
- RESTful API
- 轻松部署


### Harbor架构

<div> <img src="../assets/harbor-architecture.jpg"/> </div><br>

根据架构图可看出Harbor由5个组件组成：
- **Proxy**：代理转发所有来自浏览器和Docker Client的请求到其他后端服务。
- **Registry**：负责存储Docker images和处理Docker的`push/pull`命令，每次请求Registry都会从`token service`获得token
- **Core services**：Harbor核心服务，主要提供了以下服务：
    - **UI**：图形界面管理images
    - **Webhook**：更新日志，发起复制等其他功能
    - **Token**：根据项目的规则，负责给每一次的docker `pull/push`命令提供token
    - **Database**：存储项目、用户、规则、复制策略和镜像的元数据
- **Job services**：镜像复制，本地镜像复制（异步）到其他实例
- **Log collector**：聚集各个组件的日志存在一个地方，实质上是一个`rsyslog`服务



## 安装和配置


Harbor官方推荐安装在Linux上。本地测试这里我选择安装在MacOS。

### 环境

- macOs 10.13.6
- Docker for Mac 18.03.1-ce
- docker-compose version 1.21.1, build 5a3f1a3

### 下载
下载最新版本，目前只有offline版本。

```bash
wget https://storage.googleapis.com/harbor-releases/harbor-offline-installer-v1.6.1.tgz
tar xvf  harbor-offline-installer-v1.6.1.tgz
cd harbor
```

### 配置


#### 更新harbor.cfg

如果只是本地测试使用也不需要漏洞扫描（需要支持HTTPS）功能，只需要修改主机名，其他参数默认即可。

```
hostname = rosen.me
ui_url_protocol = https
secretkey_path = ./data
ssl_cert = ./data/cert/rosen.me.crt
ssl_cert_key = ./data/cert/rosen.me.key
```

- 必需参数
    - **hostname** 目标主机名，用来访问UI和注册服务，必须是`IP`地址或是`fully qualified domain name (FQDN)`，例如：`192.168.1.100`或`rosen.me`。不要用`localhost`或`127.0.0.1`，因为注册服务需要被外部访问。
    - **ui_url_protocol**: （http或https。默认是http）
    - **db_password**
    - **max_job_workers**
    - **customize_crt**
    - **ssl_cert**
    - **ssl_cert_key**
    - **secretkey_path**
    - ...
- 可选参数
    - **Email settings**
    - **auth_mode**
    - **harbor_admin_password** 初始密码。默认用户/密码是**admin/Harbor12345**
    - ...


#### 添加域名解析

```
sudo vim /etc/hosts

172.16.27.64 rosen.me
```

#### 配置HTTPS


- 生成CA证书

```
mkdir -p data/cert/
cd data/cert

openssl genrsa -out ca.key 4096
openssl req -x509 -new -nodes -sha512 -days 3650 \
    -subj "/C=TW/ST=Shanghai/L=Shanghai/O=example/OU=Personal/CN=rosen.me" \
    -key ca.key \
    -out ca.crt
```

- 生成服务端证书

```
# 私钥
openssl genrsa -out rosen.me.key 4096

# 证书签名
openssl req -sha512 -new \
    -subj "/C=TW/ST=Shanghai/L=Shanghai/O=example/OU=Personal/CN=rosen.me" \
    -key rosen.me.key \
    -out rosen.me.csr
```

- 生成注册实例的证书

```
cat > v3.ext <<-EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alt_names

[alt_names]
DNS.1=rosen.me
DNS.2=rosen
EOF
```

```
openssl x509 -req -sha512 -days 3650 \
    -extfile v3.ext \
    -CA ca.crt -CAkey ca.key -CAcreateserial \
    -in rosen.me.csr \
    -out rosen.me.crt
```

- 为Docker配置客户端证书

```
openssl x509 -inform PEM -in rosen.me.crt -out rosen.me.cert
mkdir -p ~/.docker/certs.d/rosen.me
cp rosen.me.cert rosen.me.key ca.crt ~/.docker/certs.d/rosen.me
```

#### 更新docker-compose.cfg

以下是修改后的内容，主要把本地绝对路径改成了相对路径，只是不想污染本地环境。

```yaml
version: '2'
services:
  log:
    image: goharbor/harbor-log:v1.6.1
    container_name: harbor-log
    restart: always
    volumes:
      - ./var/log/harbor/:/var/log/docker/:z
      - ./common/config/log/:/etc/logrotate.d/:z
    ports:
      - 127.0.0.1:1514:10514
    networks:
      - harbor
  registry:
    image: goharbor/registry-photon:v2.6.2-v1.6.1
    container_name: registry
    restart: always
    volumes:
      - ./data/registry:/storage:z
      - ./common/config/registry/:/etc/registry/:z
    networks:
      - harbor
    environment:
      - GODEBUG=netdns=cgo
    depends_on:
      - log
    logging:
      driver: "syslog"
      options:
        syslog-address: "tcp://127.0.0.1:1514"
        tag: "registry"
  postgresql:
    image: goharbor/harbor-db:v1.6.1
    container_name: harbor-db
    restart: always
    volumes:
      - ./data/database:/var/lib/postgresql/data:z
    networks:
      - harbor
    env_file:
      - ./common/config/db/env
    depends_on:
      - log
    logging:
      driver: "syslog"
      options:
        syslog-address: "tcp://127.0.0.1:1514"
        tag: "postgresql"
  adminserver:
    image: goharbor/harbor-adminserver:v1.6.1
    container_name: harbor-adminserver
    env_file:
      - ./common/config/adminserver/env
    restart: always
    volumes:
      - ./data/config/:/etc/adminserver/config/:z
      - ./data/secretkey:/etc/adminserver/key:z
      - ./data/:/data/:z
    networks:
      - harbor
    depends_on:
      - log
    logging:
      driver: "syslog"
      options:
        syslog-address: "tcp://127.0.0.1:1514"
        tag: "adminserver"
  ui:
    image: goharbor/harbor-ui:v1.6.1
    container_name: harbor-ui
    env_file:
      - ./common/config/ui/env
    restart: always
    volumes:
      - ./common/config/ui/app.conf:/etc/ui/app.conf:z
      - ./common/config/ui/private_key.pem:/etc/ui/private_key.pem:z
      - ./common/config/ui/certificates/:/etc/ui/certificates/:z
      - ./data/secretkey:/etc/ui/key:z
      - ./data/ca_download/:/etc/ui/ca/:z
      - ./data/psc/:/etc/ui/token/:z
    networks:
      - harbor
    depends_on:
      - log
      - adminserver
      - registry
    logging:
      driver: "syslog"
      options:
        syslog-address: "tcp://127.0.0.1:1514"
        tag: "ui"
  jobservice:
    image: goharbor/harbor-jobservice:v1.6.1
    container_name: harbor-jobservice
    env_file:
      - ./common/config/jobservice/env
    restart: always
    volumes:
      - ./data/job_logs:/var/log/jobs:z
      - ./common/config/jobservice/config.yml:/etc/jobservice/config.yml:z
    networks:
      - harbor
    depends_on:
      - redis
      - ui
      - adminserver
    logging:
      driver: "syslog"
      options:
        syslog-address: "tcp://127.0.0.1:1514"
        tag: "jobservice"
  redis:
    image: goharbor/redis-photon:v1.6.1
    container_name: redis
    restart: always
    volumes:
      - ./data/redis:/var/lib/redis
    networks:
      - harbor
    depends_on:
      - log
    logging:
      driver: "syslog"
      options:
        syslog-address: "tcp://127.0.0.1:1514"
        tag: "redis"
  proxy:
    image: goharbor/nginx-photon:v1.6.1
    container_name: nginx
    restart: always
    volumes:
      - ./common/config/nginx:/etc/nginx:z
    networks:
      - harbor
    ports:
      - 80:80
      - 443:443
      - 4443:4443
    depends_on:
      - postgresql
      - registry
      - ui
      - log
    logging:
      driver: "syslog"
      options:
        syslog-address: "tcp://127.0.0.1:1514"
        tag: "proxy"
networks:
  harbor:
    external: false
```

### 安装


- 安装/重装

Harbor可以与`Notary`、`Clair`（漏洞扫描）和`chartmuseum`（Helm图表存储服务）一起安装通过`--with-notary`、`--with-clair`和`--with-chartmuseum`指定。也可以不指定，最小化安装。

```
sudo ./install.sh --with-notary --with-clair --with-chartmuseum
```


- 检查各个服务运行状态

```
docker-compose -f docker-compose.yml -f docker-compose.notary.yml -f docker-compose.clair.yml -f docker-compose.chartmuseum.yml ps

       Name                     Command                  State                                    Ports
-------------------------------------------------------------------------------------------------------------------------------------
chartmuseum          /docker-entrypoint.sh            Up (healthy)   9999/tcp
clair                /docker-entrypoint.sh            Up (healthy)   6060/tcp, 6061/tcp
harbor-adminserver   /harbor/start.sh                 Up (healthy)
harbor-db            /entrypoint.sh postgres          Up (healthy)   5432/tcp
harbor-jobservice    /harbor/start.sh                 Up
harbor-log           /bin/sh -c /usr/local/bin/ ...   Up (healthy)   127.0.0.1:1514->10514/tcp
harbor-ui            /harbor/start.sh                 Up (healthy)
nginx                nginx -g daemon off;             Up (healthy)   0.0.0.0:443->443/tcp, 0.0.0.0:4443->4443/tcp, 0.0.0.0:80->80/tcp
notary-server        /bin/server-start.sh             Up
notary-signer        /bin/signer-start.sh             Up
redis                docker-entrypoint.sh redis ...   Up             6379/tcp
registry             /entrypoint.sh /etc/regist ...   Up (healthy)   5000/tcp

```

#### 访问管理页面
如果一切都ok，此时就可以通过浏览器访问管理页面了：https://rosen.me
使用默认用户名/密码(admin/Harbor12345)即可登录，不过浏览器会出现证书不信任的情况。

MacOS下可通过以下命令管理证书：

```
# add
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ca.crt
# delete
sudo security delete-certificate -c "rosen.me"
```


#### Push Image
在push镜像前需要先登录管理页，创建一个新项目。例如`rosen`。

```
# 先登录
docker login rosen.me

# 打Tag
docker tag redis:4.0 rosen.me/rosen/redis:4.0

# push
docker push rosen.me/rosen/redis:4.0
The push refers to repository [rosen.me/rosen/redis]
dbcc6e4e605c: Pushed
7dd0e697055f: Pushed
46bfaddcc6ee: Pushed
e9a42011bbb5: Pushed
3d00edfc2170: Pushed
ba291263b085: Pushed
4.0: digest: sha256:fc13b47aca9b5b53f625efe91bcd5cc44c637e80a81e5b223d5a98a6eac7ceb2 size: 1571
```


#### 启动容器

```bash
docker-compose -f docker-compose.yml -f docker-compose.notary.yml -f docker-compose.clair.yml -f docker-compose.chartmuseum.yml up -d
```

#### 销毁容器

```bash
docker-compose -f docker-compose.yml -f docker-compose.notary.yml -f docker-compose.clair.yml -f docker-compose.chartmuseum.yml down
```

#### 故障排除
- 如果有服务状态不是**UP**状态，查看目录var/log/harbor/对应的日志
- 如果遇到容器内读写本地配置文件有`permission denied`的情况，赋予文件可读权限（644）即可
- **auth_mode** 在第一次设置后，如果需要修改必须先清除db数据，re-install 才能生效


## 参考

- [Harbor仓库](https://github.com/goharbor/harbor)
- [Installation and Configuration Guide](https://github.com/goharbor/harbor/blob/master/docs/installation_guide.md)
- [Configuring Harbor with HTTPS Access](https://github.com/goharbor/harbor/blob/master/docs/configure_https.md)
- [Add a Custom CA Certificates On MacOS](https://docs.docker.com/docker-for-mac/#add-custom-ca-certificates-server-side)

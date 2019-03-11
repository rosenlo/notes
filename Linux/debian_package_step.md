<!-- vim-markdown-toc GFM -->

* [自定义 Debian 安装包](#自定义-debian-安装包)
    * [一、工具安装](#一工具安装)
    * [二、环境准备](#二环境准备)
    * [2.1 配置基础环境](#21-配置基础环境)
            * [环境](#环境)
            * [主机](#主机)
    * [2.2 配置 deb 包](#22-配置-deb-包)
        * [2.2.1 安装路径](#221-安装路径)
        * [2.2.2 修改配置](#222-修改配置)
    * [三、安装辅助脚本](#三安装辅助脚本)
        * [3.1 安装前执行的脚本](#31-安装前执行的脚本)
        * [3.2 安装后执行的脚本](#32-安装后执行的脚本)
        * [3.3 卸载前执行的脚本](#33-卸载前执行的脚本)
        * [3.3 卸载后执行的脚本](#33-卸载后执行的脚本)
    * [四、打包](#四打包)
    * [五、更新仓库](#五更新仓库)
        * [5.1 上传到服务器的仓库目录并更新](#51-上传到服务器的仓库目录并更新)
        * [5.2 配置自定义 deb 源](#52-配置自定义-deb-源)
        * [5.2 安装自定义 deb 包](#52-安装自定义-deb-包)

<!-- vim-markdown-toc -->


## 自定义 Debian 安装包

下面以 ops-updater 为例，安装过程如下:

### 一、工具安装

```bash
yum -y install ruby rubygems ruby-devel rpm-build yum-utils
gem sources -r https://rubygems.org/
gem sources -a http://ruby.sdutlinux.org/
```

### 二、环境准备

### 2.1 配置基础环境

##### 环境

- CentOS: 6.6 2.6.32-431.11.9.el6.ucloud.x86_64
- fpm: 1.9.3

##### 主机

role    | ip       | hostname
--------|----------|---------
打包机  | 10.9.0.1 | fpm09-001
deb仓库 | 10.9.0.2 | deb09-002


### 2.2 配置 deb 包

#### 2.2.1 安装路径

```bash
mkdir /tmp/ops-updater && cd /tmp/ops-updater

# 这里的 opt 是安装路径，对应系统/opt
mkdir -p opt/ops-updater && cd opt/ops-updater

git clone https://github.com/RosenLo/ops-updater.git
```

#### 2.2.2 修改配置

```bash
mv cfg.example.json cfg.json

# 修改配置
vim cfg.json
...
```

### 三、安装辅助脚本

#### 3.1 安装前执行的脚本


vim /data/script/install_ops_updater_before.sh

```bash
#!/bin/bash

# custom
...
```

#### 3.2 安装后执行的脚本


vim /data/script/install_ops_updater_after.sh

```bash
#!/bin/bash

autostart() {
    if [ $(grep "/opt/ops-updater/control restart" /etc/rc.local | wc -l) == 0 ];then
        echo "/opt/ops-updater/control restart" >> /etc/rc.local
    fi
}

host_check() {
    if [ $(grep "git.corp.imdada.cn" /etc/hosts | wc -l) == 0 ];then
        echo "10.10.91.25 git.corp.imdada.cn" >> /etc/hosts
    fi
}

start_server() {
    chown -R app:app /opt/ops/updater
	sudo su - app -c "/opt/ops-updater/control start" >> /tmp/ops-updater.txt 2>&1
    exit 0
}

autostart
host_check
start_server
```

#### 3.3 卸载前执行的脚本


vim /data/script/uninstall_ops_updater_before.sh

```bash
#!/bin/bash

APP_NAME="ops-updater"

stop_agent() {
	sudo su - app -c "/opt/ops-updater/control stop"  >> /tmp/ops-updater.txt 2>&1
}


if [[ `ps aux|grep $APP_NAME |grep -v grep` > 0 ]];then
	stop_agent
    rm -rf /opt/var
	exit 0
else
	echo "ops-updater is stopped"
fi
```

#### 3.3 卸载后执行的脚本

vim /data/script/uninstall_ops_updater_after.sh

```bash
#!/bin/bash
rm -rf /opt/ops-updater
```

### 四、打包

```bash
cd /tmp/ops-updater

fpm -s dir -t deb \
-n ops-updater \
-v 1.0.1 \
-C ops-updater \
--description 'dada agent updater deb package'  \
--before-install /data/script/install_ops_updater_before.sh \
--after-install /data/script/install_ops_updater_after.sh \
--before-remove /data/script/uninstall_ops_updater_before.sh \
--after-remove /data/script/uninstall_ops_updater_after.sh
```


### 五、更新仓库

#### 5.1 上传到服务器的仓库目录并更新

```bash
# root@deb09-002:/data/repo/ubuntu
dpkg-scanpackages . | gzip > ./Packages.gz
```


#### 5.2 配置自定义 deb 源


```bash
echo 'deb http://deb.corp.imdada.cn/ubuntu ./' >> /etc/apt/sources.list && apt-get update
```

#### 5.2 安装自定义 deb 包


```bash
apt-get -y --force-yes install ops-updater
```

## elasticsearch 打包

### 一、工具安装

>\# yum -y install ruby rubygems ruby-devel rpm-build yum-utils  
>\# gem sources -r https://rubygems.org/  
>\# gem sources -a http://ruby.sdutlinux.org/

### 二、环境准备
####2.1 配置官方源

>\# vim /etc/yum.repos.d/elasticsearch.repo

```
[elasticsearch-2.x]
name=Elasticsearch repository for 2.x packages
baseurl=https://packages.elastic.co/elasticsearch/2.x/centos
gpgcheck=1
gpgkey=https://packages.elastic.co/GPG-KEY-elasticsear
enabled=1
```

####2.2 修改基础环境


>\# mkdir /tmp/elasticsearch  
>\# cd /tmp/elasticsearch  
>\# yumdownloader elasticsearch  
>\# rpm2cpio elasticsearch-2.3.3.rpm| cpio -ivd  
>\# mkdir -p {opt/,data,data/log/}  
>\# mv usr/share/elasticsearch/ opt/ && rm -rf usr/share  
>\# mv var/lib/elasticsearch/ data/ && rm -rf var/lib  
>\# mv var/log/elasticsearch/ data/log/ && rm -rf var/log



####2.3 修改基础配置

>\# vim etc/sysconfig/elasticsearch

```
ES_HOME=/opt/elasticsearch
CONF_DIR=/etc/elasticsearch
DATA_DIR=/data/elasticsearch
LOG_DIR=/data/log/elasticsearch
PID_DIR=/var/run/elasticsearch
ES_HEAP_SIZE=256m
ES_USER=elasticsearch
ES_GROUP=elasticsearch
ES_STARTUP_SLEEP_TIME=5
MAX_OPEN_FILES=65535
MAX_LOCKED_MEMORY=unlimited
MAX_MAP_COUNT=262144
```

>\# vim etc/elasticsearch/elasticsearch.yml

```yaml
node.name:
node.master: false
node.data: true
bootstrap.mlockall: true
network.host: 0.0.0.0
```

>\# vim etc/init.d/elasticsearch

```
ES_HOME="/opt/elasticsearch"
LOG_DIR="/data/log/elasticsearch"
DATA_DIR="/data/elasticsearch"
```
>\# vim usr/lib/systemd/system/elasticsearch.service

```
[Service]
Environment=ES_HOME=/opt/elasticsearch
Environment=CONF_DIR=/etc/elasticsearch
Environment=DATA_DIR=/data/elasticsearch
Environment=LOG_DIR=/data/log/elasticsearch
Environment=PID_DIR=/var/run/elasticsearch
EnvironmentFile=-/etc/sysconfig/elasticsearch
WorkingDirectory=/opt/elasticsearch
User=elasticsearch
Group=elasticsearch
ExecStartPre=/opt/elasticsearch/bin/elasticsearch-systemd-pre-exec
ExecStart=/opt/elasticsearch/bin/elasticsearch
```

### 三、安装辅助脚本

####3.1 安装后执行的脚本


>\# vim /tmp/elasticsearch/etc/elasticsearch/scripts/install_es_after.sh

```bash
#!/bin/bash
CLUSTER_NAME=`echo $(hostname) |sed 's/-[0-9]*$//'`
ES_Version='2.3.3'
ES_LOG='/data/log/elasticsearch'
ES_HOME='/opt/elasticsearch'
ES_DATA='/data/elasticsearch'
ES_PID='/var/run/elasticsearch'
Mem_Total=`cat /proc/meminfo |grep -i MemTotal|awk '{print $2}'`
Divi=1000000
Ip=`ifconfig |grep inet|grep -v net6|grep -v 127|awk '{print $2}'`
Size=`[[ -n $Mem_Total ]] && echo  $Mem_Total"/"$Divi"/"2|bc`

[[ -z `getent passwd elasticsearch` ]] && adduser elasticsearch -s /sbin/nologin
if [[ $Size -lt 32 ]];then
	[[ `expr $Size \>= 1 ` -eq 1  ]] && sed -i "s/ES_HEAP_SIZE=.*/ES_HEAP_SIZE=${Size}g/" /etc/sysconfig/elasticsearch
else
	sed -i "s/ES_HEAP_SIZE=.*/ES_HEAP_SIZE=31g/" /etc/sysconfig/elasticsearch
fi

sed -i "s/node.name.*/node.name: $(hostname)/" /etc/elasticsearch/elasticsearch.yml
sed -i "s/cluster.name.*/cluster.name: $CLUSTER_NAME/" /etc/elasticsearch/elasticsearch.yml
grep 'elasticsearch' /etc/security/limits.conf /dev/null 2>&1 || echo 'elasticsearch - memlock unlimited' >/etc/security/limits.conf
ulimit -l unlimited

cd /opt
mv elasticsearch elasticsearch-$ES_Version
ln -s elasticsearch-$ES_Version elasticsearch
chown elasticsearch:elasticsearch $ES_LOG  $ES_HOME $ES_DATA $ES_PID
$ES_HOME/bin/plugin install  mobz/elasticsearch-head

if [[ -z `grep -i 'release 6' /etc/redhat-release` ]];then
	systemctl daemon-reload
	systemctl enable elasticsearch.service
else
	/sbin/chkconfig --add nginx
    /sbin/chkconfig --level 235 nginx on
fi
```

####3.2 卸载前执行的脚本


>\# vim /tmp/elasticsearch/etc/elasticsearch/scripts/uninstall_es_before.sh

```bash
#!/bin/bash
[[ -n `ps aux|grep elasticsearch|grep -v grep` ]] && killall -9 elasticsearch

if [[ -z `grep 'release 6' /etc/centos-release` ]]; then
    systemctl disable elasticsearch.service >/dev/null 2>&1
else
    /sbin/chkconfig --del elasticsearch
fi
```

####3.3 卸载后执行的脚本

>\# vim /tmp/elasticsearch/etc/elasticsearch/scripts/uninstall_es_after.sh

```bash
#!/bin/bash
rm -rf /opt/elasticsearch*
rm -rf /etc/elasticsearch
rm -rf /data/elasticsearch
rm -rf /data/log/elasticsearch
```

### 四、打包

	
>\# fpm -s dir \  
>-t rpm \  
>-n elasticsearch -v 2.3.3 \  
>--iteration 1.ele -C \  
>/tmp/elasticsearch/ \  
>-p /tmp/ \  
>--description 'eleme elasticsearch rpm package' \  
>--url 'https://www.elastic.co/' \  
>--license '(c) 2009' \  
>--after-install /tmp/elasticsearch/etc/elasticsearch/scripts/install_es_after.sh \  
>--after-remove /tmp/elasticsearch/etc/elasticsearch/scripts/uninstall_es_after.sh \  
>--before-remove /tmp/elasticsearch/etc/elasticsearch/scripts/uninstall_es_before.sh


### 五、放到yum仓库
####5.1 上传到服务器的仓库目录，并执行:

>\# createrepo -pdo ./ ./


####5.2 安装命令如下：


>\# yum makecache  
>\# yum -y install elasticsearch

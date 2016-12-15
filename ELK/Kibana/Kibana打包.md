## Kibana 打包

### 一、工具安装


>\# yum -y install ruby rubygems ruby-devel rpm-build yum-utils  
>\# gem sources -r https://rubygems.org/  
>\# gem sources -a http://ruby.sdutlinux.org/


### 二、环境准备

####2.1 配置官方源

>\# vim /etc/yum.repos.d/kibana.repo

```
[kibana-4.5]
name=Kibana repository for 4.5.x packages
baseurl=http://packages.elastic.co/kibana/4.5/centos
gpgcheck=1
gpgkey=http://packages.elastic.co/GPG-KEY-elasticsearch
enabled=1
```

####2.2 修改基础环境

>\# mkdir /tmp/kibana  
>\# cd /tmp/kibana  
>\# yumdownloader kibana  
>\# rpm2cpio kibana-4.5.3-1.x86_64.rpm |cpio -ivd  
>\# sed -i 's/var/data/g' etc/init.d/kibana  
>\# sed -i 's/\/var\/lib/\/data/g' etc/init.d/logstash

####2.3 修改基础配置

>\# vim opt/kibana/config/kibana.yml

```yaml
server.port: 5601
server.host: "0.0.0.0"
elasticsearch.url: "http://localhost:9200"
kibana.index: ".kibana"
logging.dest: /data/log/kibana/kibana.stdout
```

### 三、安装辅助脚本

####3.1 安装后执行的脚本

>\# vim install_kb_after.sh

```bash
#!/bin/bash
Version=4.5.3
cd /opt/
mv kibana kibana-$Version
ln -s kibana-$Version kibana
mkdir -p /data/log/kibana

if [[ -z `grep -i 'release 6' /etc/redhat-release` ]];then
	systemctl daemon-reload
	systemctl enable kibana.service
else
	/sbin/chkconfig --add kibana
	/sbin/chkconfig --level 235 kibana on
fi

[[ -z `getent passwd kibana` ]] && adduser kibana -s /sbin/nologin -G root
chown -R kibana kibana* /data/log/kibana
```

####3.2 卸载前执行的脚本

>\# vim uninstall_kb_before.sh

```bash
#!/bin/bash
[[ -n `ps aux|grep kibana|grep -v grep` ]] && systemctl stop kibana.service || /etc/init.d/kibana stop

if [[ -z `grep 'release 6' /etc/centos-release` ]];then
	systemctl disable kibana.service /dev/null 2>&1
else
	/sbin/chkconfig --del kibana
fi

userdel -r kibana
```

####3.3 卸载后执行的脚本

>\# vim uninstall_kb_after.sh

```bash
#!/bin/bash
rm -rf /opt/kibana*
rm -rf /etc/kibana*
rm -rf /data/log/kibana*
```

### 四、打包

>\# fpm -s dir \  
>-t rpm \  
>-n kibana \  
>-v 4.5.3 \  
>--iteration 1.ele \  
>-C ./kibana/ -p ./ \
>--description 'eleme kibana rpm package' \  
>--url 'https://www.elastic.co' \  
>--license 'Apache 2.0' \  
>--after-install install_kb_after.sh \  
>--before-remove uninstall_kb_before.sh \  
>--after-remove uninstall_kb_after.sh


### 五、放到yum仓库

####5.1 上传到服务器的仓库目录，并执行:

>\# createrepo -pdo ./ ./




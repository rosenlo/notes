## Logstash 打包

### 一、工具安装
```
  # yum -y install ruby rubygems ruby-devel rpm-build yum-utils
  # gem sources -r https://rubygems.org/
  # gem sources -a http://ruby.sdutlinux.org/
```
### 二、环境准备
####2.1 配置官方源


>\# vim /etc/yum.repos.d/logstash.repo

```
[logstash-2.3]
name=Logstash repository for 2.3.x packages
baseurl=https://packages.elastic.co/logstash/2.3/centos
gpgcheck=1
gpgkey=https://packages.elastic.co/GPG-KEY-elasticsearch
enabled=1
```


####2.2 修改基础环境

>\# mkdir /tmp/logstah  
>\# cd /tmp/logstash  
>\# yumdownloader logstash-2.3.3-1  
>\# mv logstash logstash-2.3.3  
>\# rpm2cpio logstash-2.3.3-1.noarch.rpm| cpio -ivd  
>\# mkdir data
>\# sed -i 's/var/data/g' etc/logrotate.d/logstash  
>\# mv var/lib/logstash  data/  
>\# rm -rf var/  
>\# sed -i 's/\/var\/lib/\/data/g' etc/init.d/logstash

### 三、安装辅助脚本

####3.1 安装后执行的脚本

>\# vim install_ls_after.sh

```
#!/bin/bash
Version=2.3.3
cd /opt
mv logstash logstash-$Version
ln -s  logstash-$Version  logstash
```

####3.2 卸载后执行的脚本

>\# vim uninstall_ls_after.sh

```
#!/bin/bash
rm -rf /opt/logstash*
rm -rf /etc/logstash*
rm -rf /data/log/logstash*
rm -rf /data/logstash*
rm -rf /etc/logrotate.d/logstash*
```

### 四、打包

>\# fpm -s dir -t rpm -n logstash -v 2.3.3 --iteration 1.ele -C logstash/ -p ./  --description 'eleme logstash rpm package' --url 'http://www.elasticsearch.org/overview/logstash/' --license 'ASL 2.0' --after-install ./install_ls_after.sh --after-remove ./uninstall_ls_after.sh --before-remove ./uninstall_ls_before.sh  -a noarch --vendor Rosen


### 五、放到yum仓库
####5.1 上传到服务器的仓库目录，并执行:

>\# createrepo -pdo ./ ./

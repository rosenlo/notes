# systemd
### 服务权限
systemd有系统和用户区分；系统目录`/user/lib/systemd/system/`、用户目录`/etc/lib/systemd/user/`、系统管理员自定义目录`/etc/systemd/system/`,一般建议存放在`/etc/systemd/system/`下面。


### 创建服务文件
服务文件命名规则: `server_name.service`

**示例:**

>\# vim es_agent.service 

```
[Unit]
Description=Monitor agent for ElasticSearch
After=network.target

[Service]
Type=forking
WorkingDirectory=/path/es_agent
PIDFile=/var/run/elasticsearch/es_agent.pid
ExecStart=/path/es_agent/es_agent start
ExecReload=/path/es_agent/es_agent restart
ExecStop=/path/es_agent/es_agent stop
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```

### 简单说明
 
```
[Unit]
Description: 服务描述
After: 定义启动顺序

[Service]
Type: forking - 进程开始通过ExecStart产生一个子进程，当父进程启动完成退出后它将成为这个服务的主进程。
WorkingDirectory: 工作目录
PIDFile: 指定PID的存放目录
ExecStart: 指定一个命令或脚本去启动这个服务
ExecReload: 指定一个命令或脚本去重载这个服务
ExecStop: 指定一个命令或脚本去停止这个服务
PrivateTmp: ture - 为该进程新建一块系统空间，挂载到/tmp或/var/tmp目录，不会与其他进程共享

[Install]
WantedBy=multi-user.target - 表示该服务所在的target
```

### 重载服务

**⚠️note：** 每次修改或移除文件时需要重载服务!

重载服务:
> \# systemctl daemon-reload

加入开机启动：
> \# systemctl enable es_agent.service

重启服务：
> \# systemctl restart es_agent.service


---

**参考：**
>1. <https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/System_Administrators_Guide/sect-Managing_Services_with_systemd-Unit_Files.html>

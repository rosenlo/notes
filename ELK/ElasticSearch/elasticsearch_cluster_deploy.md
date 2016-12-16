# ElasticSearch集群部署

## 一、安装

>\# yum makecache  
>\# yum -y install elasticsearch

## 二、集群配置
　　本文集群配置部署基于**ElasticSearch2.3.3的版本**，ES的`CONF_DIR`里面有两个配置文件：`elasticsearch.yml`和`logging.yml`，第一个是ES的基本配置文件，其中大部分默认配置是合理并且实用的；第二个是日志配置文件，ES也是使用log4j来记录日志的，所以`logging.yml`里的设置按普通log4j配置文件来设置就行了。

####2.1 elasticsearch.yml 参考配置：

```yaml
cluster.name: test-search # 默认是主机头
node.name: test-search-1 # 默认是主机名
node.master: true # 允许这个节点被选举为主节点（默认为true）
node.data: true # 允许这个节点存储数据（默认为true）
bootstrap.mlockall: true
network.host: 0.0.0.0
discovery.zen.ping.timeout: 5s
discovery.zen.ping.unicast.hosts: ["test-search-1", "test-search-2", "test-search-3", "test-search-4", "test-search-5"] # 集群节点初始列表
gateway.recover_after_nodes: 3 #一个集群中的3个节点启动后，才允许进行恢复处理
gateway.expected_nodes: 3 # 集群期待节点，一旦这个3个节点启动（并且revoce_after_nodes也符合）立即开始恢复过程（不等待recover_after_time超时）
gateway.recover_after_time: 5m
index.number_of_shards: 5 # 索引分片数量（默认为5）；索引分片分得多一些，可以提高索引的性能，并且把一个大的索引分布到机器中去。
index.number_of_replicas: 2 # 索引副本数量（默认为1）；副本分片分得多一些，可以提高搜索的性能，并且提高集群的可用性。 
discovery.zen.minimum_master_nodes: 2 # 集群大于3个节点，可用设成一个高点的值（2-4）
```

**注意**：

1. `number_of_shards`对一个索引来说只能配置一次。
2. `number_of_replicas`在任何时候都可以增加和减少，通过`update index settings`（更新索引接口）API。


> Reference：
> 
>- https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-node.html#master-node
>- https://www.elastic.co/guide/en/elasticsearch/reference/2.3/index-modules.html#dynamic-index-settings

####2.2 elasticsearch 参考配置
- 在`/etc/sysconfig/`目录下有个**elasticsearch**配置文件，主要是设置一些环境参数和java运行参数，参考配置：

```
ES_HOME=/opt/elasticsearch
CONF_DIR=/etc/elasticsearch
DATA_DIR=/data/elasticsearch/
LOG_DIR=/data/log/elasticsearch
PID_DIR=/var/run/elasticsearch
ES_HEAP_SIZE=16g  # 分配给ES的内存大小，最佳设置为50%RAM，不超过31g
ES_USER=elasticsearch
ES_GROUP=elasticsearch
ES_STARTUP_SLEEP_TIME=5
MAX_OPEN_FILES=65535
MAX_MAP_COUNT=262144
MAX_LOCKED_MEMORY=unlimited 
```

- 要让elasticsearch达到最佳性能，需要设置如下参数：

```
# vi /etc/security/limits.conf

elasticsearch - nofile 65535
elasticsearch - memlock unlimited

# vi /etc/sysconfig/elasticsearch

MAX_LOCKED_MEMORY=unlimited
ES_HEAP_SIZE=''  # 50%可用内存，不超过31g
MAX_OPEN_FILES=65535

# vi /etc/elasticsearch/elasticsearch.yml

bootstrap.mlockall: true
```

> References：
> 
>- https://github.com/grigorescu/Brownian/wiki/ElasticSearch-Configuration

- 更多模块配置请参考：

> https://www.elastic.co/guide/en/elasticsearch/reference/current/modules.html


## 三、集群管理插件安装
####3.1 elasticsearch-head

`head`是一个ElasticSearch的集群管理工具，它是完全由html5编写的独立网页程序

>\# $ES_HOME/bin/plugin -install mobz/elasticsearch-head

访问：https://example.com/search/_plugin/head/


####3.2 elasticsearch-HQ （http://www.elastichq.org/）

`Elastic HQ`提供一个Web应用程序来管理和监控ElasticSearch实例与集群管理和监控。具有良好体验、直观、强大功能的ElasticSearch监控和管理工具。提供实时监控、全集群管理、搜索和查询，无需额外的软件安装


>\# $ES_HOME/bin/plugin -install royrusso/elasticsearch-HQ


访问：https://example.com/search/_plugin/hq/

连接：https://example.com/search/

## 四、zabbix监控

关联`Template App ElasticSearch`模板即可，配置和脚步已封装在rpm包里。

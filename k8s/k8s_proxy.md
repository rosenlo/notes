# Kubernetes代理使用

## Hello Minikube
通过Minikube本地运行Kubernetes。OS X版本的Minikube实际是创建了一个Linux虚拟机，再在虚拟机上支持容器。默认驱动程序为`VirtualBox`，这里我使用了[Hyperkit](https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#hyperkit-driver)。

参考官方教程：https://kubernetes.io/docs/tutorials/hello-minikube/

## Environment
- macOs 10.13.6
- ShadowsocksX-NG 1.8.2
- Docker for Mac 18.03.1-ce
- Minikube v0.25.0

## Proxy Configuration
配置为：`http://0.0.0.0:1087`，需要注意以下问题：
- 一般代理配置默认监听`127.0.0.1:1087`，由于是在虚拟机内创建的容器，因为虚拟机网段的不同并不能使用`127.0.0.1`或`localhost`等地址，对于虚拟机来说`localhost`为虚拟机本身，并不是宿主机。所以须设置为宿主机与虚拟机网段通讯的IP地址。

## Create a Minikube cluster
这里的`HTT_PROXY`和`HTTPS_PROXY`填宿主机地址`192.168.64.1`，这里虚拟机的网段为`192.168.64.0/24`
```
minikube start --vm-driver=hyperkit --docker-env HTTP_PROXY=http://192.168.64.1:1087  --docker-env HTTPS_PROXY=http://192.168.64.1:1087
```
对于已创建过的集群，可以通过配置文件修改`~/.minikube/machines/minikube/config.json`
```
  "Env": [
    "HTTP_PROXY=http://192.168.64.1:1087",
    "HTTPS_PROXY=http://192.168.64.1:1087"
  ],
```

以上设置完后，虚拟机中的容器也可以愉快的访问外面的世界了！

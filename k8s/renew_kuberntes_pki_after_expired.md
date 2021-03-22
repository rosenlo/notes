<!-- vim-markdown-toc GFM -->

* [证书过期后重新生成证书](#证书过期后重新生成证书)
    * [环境](#环境)
    * [操作步骤](#操作步骤)
    * [参考](#参考)

<!-- vim-markdown-toc -->

# 证书过期后重新生成证书

## 环境

- CentOS: 7.2.1511
- kube-apiserver:v1.13.2

## 操作步骤

```
cd /etc/kubernetes/pki/ && mkdir old_cert
mv {apiserver.crt,apiserver-etcd-client.key,apiserver-kubelet-client.crt,front-proxy-ca.crt,front-proxy-client.crt,front-proxy-client.key,front-proxy-ca.key,apiserver-kubelet-client.key,apiserver.key,apiserver-etcd-client.crt} old_cert
kubeadm init phase certs all --apiserver-advertise-address <IP>

cd /etc/kubernetes/ && mkdir old_conf
mv {admin.conf,controller-manager.conf,kubelet.conf,scheduler.conf} old_conf
kubeadm init phase kubeconfig all

reboot

cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
```


## 参考

- [Renew kubernetes pki after expired](https://stackoverflow.com/questions/56320930/renew-kubernetes-pki-after-expired/56334732#56334732)

# How to use Kata Containers and Containerd With Kubernetes (Handbook)


<!-- vim-markdown-toc GFM -->

* [Envrionment](#envrionment)
* [Install And Configure Containerd](#install-and-configure-containerd)
* [Install And Configure Kubernetes](#install-and-configure-kubernetes)
    * [Install and Configure CNI (flannel)](#install-and-configure-cni-flannel)
* [References](#references)

<!-- vim-markdown-toc -->

## Envrionment

- CentOS Linux release 7.8.2003 (Core) 4.19.12-1.el7.elrepo.x86_64
- go version go1.14.4 linux/amd64
- kata-runtime  1.12.0-alpha0
- containerd 1.2.13
- kubeadm-1.16.3
- kubectl-1.16.3
- kubelet-1.16.3




## Install And Configure Containerd

- Install required packages

    ```bash
    yum install -y yum-utils device-mapper-persistent-data lvm2
    ```

- Add docker repository

    ```bash
    yum-config-manager \
        --add-repo \
        https://download.docker.com/linux/centos/docker-ce.repo
    ```

- Install containerd

    ```bash
    yum update -y && yum install -y containerd.io
    ```

- Configure containerd to use Kata Containers

    ```bash
    mkdir -p /etc/containerd
    containerd config default > /etc/containerd/config.toml
    ```

    /etc/containerd/config.toml

    ```
    root = "/data/lib/containerd"
    state = "/data/run/containerd"
    oom_score = 0

    [grpc]
    address = "/run/containerd/containerd.sock"
    uid = 0
    gid = 0
    max_recv_message_size = 16777216
    max_send_message_size = 16777216

    [debug]
    address = ""
    uid = 0
    gid = 0
    level = ""

    [metrics]
    address = ""
    grpc_histogram = false

    [cgroup]
    path = ""

    [plugins]
    [plugins.cgroups]
        no_prometheus = false
    [plugins.cri]
        stream_server_address = "127.0.0.1"
        stream_server_port = "0"
        enable_selinux = false
        sandbox_image = "k8s.gcr.io/pause:3.1"
        stats_collect_period = 10
        systemd_cgroup = true
        enable_tls_streaming = false
        max_container_log_line_size = 16384
        disable_proc_mount = false
        [plugins.cri.containerd]
        snapshotter = "overlayfs"
        no_pivot = false
        [plugins.cri.containerd.runtimes]
        [plugins.cri.containerd.runtimes.runc]
            runtime_type = "io.containerd.runc.v1"
            [plugins.cri.containerd.runtimes.runc.options]
            NoPivotRoot = false
            NoNewKeyring = false
            ShimCgroup = ""
            IoUid = 0
            IoGid = 0
            BinaryName = "runc"
            Root = ""
            CriuPath = ""
            SystemdCgroup = false
        [plugins.cri.containerd.runtimes.kata]
            runtime_type = "io.containerd.kata.v2"
        [plugins.cri.containerd.runtimes.katacli]
            runtime_type = "io.containerd.runc.v1"
            [plugins.cri.containerd.runtimes.katacli.options]
            NoPivotRoot = false
            NoNewKeyring = false
            ShimCgroup = ""
            IoUid = 0
            IoGid = 0
            BinaryName = "/usr/bin/kata-runtime"
            Root = ""
            CriuPath = ""
            SystemdCgroup = false
        [plugins.cri.containerd.default_runtime]
            runtime_type = "io.containerd.runtime.v1.linux"
            runtime_engine = ""
            runtime_root = ""
        [plugins.cri.containerd.untrusted_workload_runtime]
            runtime_type = "io.containerd.kata.v2"
            runtime_engine = ""
            runtime_root = ""
        [plugins.cri.cni]
        bin_dir = "/opt/cni/bin"
        conf_dir = "/etc/cni/net.d"
        conf_template = ""
        [plugins.cri.registry]
        [plugins.cri.registry.mirrors]
            [plugins.cri.registry.mirrors."docker.io"]
            endpoint = ["https://registry-1.docker.io"]
        [plugins.cri.x509_key_pair_streaming]
        tls_cert_file = ""
        tls_key_file = ""
    [plugins.diff-service]
        default = ["walking"]
    [plugins.linux]
        shim = "containerd-shim"
        runtime = "runc"
        runtime_root = ""
        no_shim = false
        shim_debug = false
    [plugins.opt]
        path = "/opt/containerd"
    [plugins.restart]
        interval = "10s"
    [plugins.scheduler]
        pause_threshold = 0.02
        deletion_threshold = 0
        mutation_threshold = 100
        schedule_delay = "0s"
        startup_delay = "100ms"
    ```

-  Restart containerd

    ```bash
    systemctl daemon-reload containerd
    systemctl restart containerd
    ```

-  Modify cgroup driver `systemd`

    /etc/containerd/config.toml

    ```
    systemd_cgroup = true
    ```

- Install and configure crictl

    ```bash
    go get github.com/kubernetes-sigs/cri-tools/cmd/crictl

    cat <<EOF > /etc/crictl.yaml
    runtime-endpoint: unix:///var/run/containerd/containerd.sock
    image-endpoint: unix:///var/run/containerd/containerd.sock
    timeout: 10
    debug: true
    EOF
    ```

## Install And Configure Kubernetes


- Install required packages

    ```bash
    cat <<EOF > /etc/yum.repos.d/kubernetes.repo
    [kubernetes]
    name=Kubernetes
    baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
    enabled=1
    gpgcheck=0
    EOF
    ```

    ```bash
    swapoff -a

    cat <<EOF >  /etc/sysctl.d/k8s.conf
    net.bridge.bridge-nf-call-ip6tables = 1
    net.bridge.bridge-nf-call-iptables = 1
    EOF

    sudo sysctl --system

    ```

- Install

    ```
    yum install kubeadm-1.16.3 kubectl-1.16.3 kubelet-1.16.3
    ```

- Configure kubelet to use Containerd

    ```bash
    cat <<EOF > /etc/sysconfig/kubelet
    KUBELET_EXTRA_ARGS=--container-runtime=remote --runtime-request-timeout=15m --container-runtime-endpoint=unix:///run/containerd/containerd.sock
    EOF
    ```

- Modify cgroup-driver  `systemd`

    add `--cgroup-driver=systemd` to `KUBELET_EXTRA_ARGS`

- Restart kubelet

    ```bash
    systemctl daemon-reload
    systemctl restart kubelet
    ```


- Initialization

    ```bash
    kubeadm init --cri-socket /run/containerd/containerd.sock --pod-network-cidr=10.244.0.0/16 --kubernetes-version=v1.16.3
    ```


### Install and Configure CNI (flannel)

- flannel configurations

    ```bash
    mkdir -p /etc/cni/net.d/
    cat <<EOF> /etc/cni/net.d/10-flannel.conf
    {
        "name": "cbr0",
        "type": "flannel",
        "delegate": {
            "isDefaultGateway": true
        }
    }
    EOF

    mkdir /run/flannel/ -p

    cat <<EOF> /run/flannel/subnet.env
    FLANNEL_NETWORK=10.244.0.0/16
    FLANNEL_SUBNET=10.244.1.0/24
    FLANNEL_MTU=1450
    FLANNEL_IPMASQ=true
    EOF
    ```

- Install pod network

    ```bash
    kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
    ```

- Allow pods to run in master node

    ```bash
    kubectl taint nodes --all node-role.kubernetes.io/master-
    ```

## References

- [Container runtimes](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#containerd)
- [How to use Kata Containers and CRI (containerd plugin) with Kubernetes](https://github.com/kata-containers/documentation/blob/master/how-to/how-to-use-k8s-with-cri-containerd-and-kata.md)
- [How to use Kata Containers and Containerd](https://github.com/kata-containers/documentation/blob/master/how-to/containerd-kata.md#how-to-use-kata-containers-and-containerd)
- [Kata Containers Architecture](https://github.com/kata-containers/kata-containers/blob/2.0-dev/docs/design/architecture.md)

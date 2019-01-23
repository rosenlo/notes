
<!-- vim-markdown-toc GFM -->

* [Install Etcd on Docker with GlusterFS](#install-etcd-on-docker-with-glusterfs)
    * [Environment](#environment)
    * [Prepaering](#prepaering)
    * [Install Docker Volume Plugin](#install-docker-volume-plugin)
    * [TLS](#tls)
        * [Install on Linux](#install-on-linux)
        * [Generate self-signed root CA certificate](#generate-self-signed-root-ca-certificate)
        * [Generate local-issued certificates with private keys](#generate-local-issued-certificates-with-private-keys)
        * [Transferring certs to remote machines](#transferring-certs-to-remote-machines)
    * [Run with Docker](#run-with-docker)
    * [Install etcdctl](#install-etcdctl)
    * [Troubleshooting](#troubleshooting)
    * [Reference](#reference)

<!-- vim-markdown-toc -->

# Install Etcd on Docker with GlusterFS

## Environment

- CentOS: 7.2.1511
- docker: ce-18.06.0.ce
- etcd: k8s.gcr.io/etcd:3.2.24
- cfssl: R1.2
- docker volume plugin: trajano/glusterfs-volume-plugin
- glusterfs: 5.1
- glusterfs-fuse: 5.1.1


## Prepaering

hosts:

 node   | ip        | hostname
--------|-----------|-------------
node1   | 10.9.0.1  | k8s-mgmgt09-001
node2   | 10.9.0.2  | k8s-node09-001
node3   | 10.10.0.1 | k8s-node10-001


## Install Docker Volume Plugin

在每个节点创建对应的 `volume`

```bash
docker plugin install trajano/glusterfs-volume-plugin

docker volume create -d trajano/glusterfs-volume-plugin --opt servers=serverIP volumeName/subName
```

## TLS

### Install on Linux

```bash
rm -f /tmp/cfssl* && rm -rf /tmp/certs && mkdir -p /tmp/certs

curl -L https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 -o /tmp/cfssl
chmod +x /tmp/cfssl
sudo mv /tmp/cfssl /usr/local/bin/cfssl

curl -L https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 -o /tmp/cfssljson
chmod +x /tmp/cfssljson
sudo mv /tmp/cfssljson /usr/local/bin/cfssljson

/usr/local/bin/cfssl version
/usr/local/bin/cfssljson -h

mkdir -p /tmp/certs
```

### Generate self-signed root CA certificate

```bash
mkdir -p /tmp/certs

cat > /tmp/certs/etcd-root-ca-csr.json <<EOF
{
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "O": "etcd",
      "OU": "etcd Security",
      "L": "San Francisco",
      "ST": "California",
      "C": "USA"
    }
  ],
  "CN": "etcd-root-ca"
}
EOF

cfssl gencert --initca=true /tmp/certs/etcd-root-ca-csr.json | cfssljson --bare /tmp/certs/etcd-root-ca

# verify
openssl x509 -in /tmp/certs/etcd-root-ca.pem -text -noout


# cert-generation configuration
cat > /tmp/certs/etcd-gencert.json <<EOF
{
  "signing": {
    "default": {
        "usages": [
          "signing",
          "key encipherment",
          "server auth",
          "client auth"
        ],
        "expiry": "87600h"
    }
  }
}
EOF
```

Results:

```
# CSR configuration
/tmp/certs/etcd-root-ca-csr.json

# CSR
/tmp/certs/etcd-root-ca.csr

# self-signed root CA public key
/tmp/certs/etcd-root-ca.pem

# self-signed root CA private key
/tmp/certs/etcd-root-ca-key.pem

# cert-generation configuration for other TLS assets
/tmp/certs/etcd-gencert.json
```

### Generate local-issued certificates with private keys

k8s-mgmgt09-001:

```bash
mkdir -p /tmp/certs

cat > /tmp/certs/dwf-etcd1-ca-csr.json <<EOF
{
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "O": "etcd",
      "OU": "etcd Security",
      "L": "San Francisco",
      "ST": "California",
      "C": "USA"
    }
  ],
  "CN": "dwf-etcd1",
  "hosts": [
    "127.0.0.1",
    "localhost",
    "10.9.0.1"
  ]
}
EOF

cfssl gencert \
  --ca /tmp/certs/etcd-root-ca.pem \
  --ca-key /tmp/certs/etcd-root-ca-key.pem \
  --config /tmp/certs/etcd-gencert.json \
  /tmp/certs/dwf-etcd1-ca-csr.json | cfssljson --bare /tmp/certs/dwf-etcd1

# verify
openssl x509 -in /tmp/certs/dwf-etcd1.pem -text -noout
```

k8s-node09-001

```bash
mkdir -p /tmp/certs

cat > /tmp/certs/dwf-etcd2-ca-csr.json <<EOF
{
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "O": "etcd",
      "OU": "etcd Security",
      "L": "San Francisco",
      "ST": "California",
      "C": "USA"
    }
  ],
  "CN": "dwf-etcd2",
  "hosts": [
    "127.0.0.1",
    "localhost",
    "10.9.0.2"
  ]
}
EOF
cfssl gencert \
  --ca /tmp/certs/etcd-root-ca.pem \
  --ca-key /tmp/certs/etcd-root-ca-key.pem \
  --config /tmp/certs/etcd-gencert.json \
  /tmp/certs/dwf-etcd2-ca-csr.json | cfssljson --bare /tmp/certs/dwf-etcd2

# verify
openssl x509 -in /tmp/certs/dwf-etcd2.pem -text -noout
```

k8s-node10-001

```bash
mkdir -p /tmp/certs

cat > /tmp/certs/dwf-etcd3-ca-csr.json <<EOF
{
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "O": "etcd",
      "OU": "etcd Security",
      "L": "San Francisco",
      "ST": "California",
      "C": "USA"
    }
  ],
  "CN": "dwf-etcd3",
  "hosts": [
    "127.0.0.1",
    "localhost",
    "10.10.0.1"
  ]
}
EOF

cfssl gencert \
  --ca /tmp/certs/etcd-root-ca.pem \
  --ca-key /tmp/certs/etcd-root-ca-key.pem \
  --config /tmp/certs/etcd-gencert.json \
  /tmp/certs/dwf-etcd3-ca-csr.json | cfssljson --bare /tmp/certs/dwf-etcd3

# verify
openssl x509 -in /tmp/certs/dwf-etcd3.pem -text -noout
```

### Transferring certs to remote machines

生成证书后把证书传到对应的主机上

```bash
mkdir -p /data/etcd/certs
cp /tmp/certs/* /data/etcd/certs
```


## Run with Docker


dwf-k8s-mgmt09-001:

```bash
# after transferring certs to remote machines
# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /data/etcd/dwf-etcd1


# to write service file for etcd with Docker
cat > /tmp/dwf-etcd1.service <<EOF
[Unit]
Description=etcd with Docker
Documentation=https://github.com/coreos/etcd

[Service]
Restart=always
RestartSec=5s
TimeoutStartSec=0
LimitNOFILE=40000

ExecStart=/usr/bin/docker \
  run \
  --rm \
  --net=host \
  --name etcd-3.2.24 \
  --volume=dwf/dwf-k8s-etcd1:/etcd-data \
  --volume=/data/etcd/certs:/etcd-ssl-certs-dir \
  k8s.gcr.io/etcd:3.2.24 \
  /usr/local/bin/etcd \
  --name dwf-etcd1 \
  --data-dir /etcd-data \
  --listen-client-urls https://10.9.0.1:2379 \
  --advertise-client-urls https://10.9.0.1:2379 \
  --listen-peer-urls https://10.9.0.1:2380 \
  --initial-advertise-peer-urls https://10.9.0.1:2380 \
  --initial-cluster dwf-etcd1=https://10.9.0.1:2380,dwf-etcd2=https://10.9.0.2:2380,dwf-etcd3=https://10.10.0.1:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --client-cert-auth \
  --trusted-ca-file /etcd-ssl-certs-dir/etcd-root-ca.pem \
  --cert-file /etcd-ssl-certs-dir/dwf-etcd1.pem \
  --key-file /etcd-ssl-certs-dir/dwf-etcd1-key.pem \
  --peer-client-cert-auth \
  --peer-trusted-ca-file /etcd-ssl-certs-dir/etcd-root-ca.pem \
  --peer-cert-file /etcd-ssl-certs-dir/dwf-etcd1.pem \
  --peer-key-file /etcd-ssl-certs-dir/dwf-etcd1-key.pem

ExecStop=/usr/bin/docker stop etcd-3.2.24

[Install]
WantedBy=multi-user.target
EOF

sudo mv /tmp/dwf-etcd1.service /etc/systemd/system/dwf-etcd1.service


# to start service
sudo systemctl daemon-reload
sudo systemctl cat dwf-etcd1.service
sudo systemctl enable dwf-etcd1.service
sudo systemctl start dwf-etcd1.service

# to get logs from service
sudo systemctl status dwf-etcd1.service -l --no-pager
sudo journalctl -u dwf-etcd1.service -l --no-pager|less
sudo journalctl -f -u dwf-etcd1.service

# to stop service
sudo systemctl stop dwf-etcd1.service
sudo systemctl disable dwf-etcd1.service
```

dwf-k8s-node09-001:

```bash
# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /data/etcd/dwf-etcd2


# to write service file for etcd with Docker
# after transferring certs to remote machines
mkdir -p /data/etcd/certs
cp /tmp/certs/* /data/etcd/certs


# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /data/dwf/etcd/dwf-etcd2


# to write service file for etcd with Docker
cat > /tmp/dwf-etcd2.service <<EOF
[Unit]
Description=etcd with Docker
Documentation=https://github.com/coreos/etcd

[Service]
Restart=always
RestartSec=5s
TimeoutStartSec=0
LimitNOFILE=40000

ExecStart=/usr/bin/docker \
  run \
  --rm \
  --net=host \
  --name etcd-3.2.24 \
  --volume=dwf/dwf-k8s-etcd2:/etcd-data \
  --volume=/data/etcd/certs:/etcd-ssl-certs-dir \
  k8s.gcr.io/etcd:3.2.24 \
  /usr/local/bin/etcd \
  --name dwf-etcd2 \
  --data-dir /etcd-data \
  --listen-client-urls https://10.9.0.2:2379 \
  --advertise-client-urls https://10.9.0.2:2379 \
  --listen-peer-urls https://10.9.0.2:2380 \
  --initial-advertise-peer-urls https://10.9.0.2:2380 \
  --initial-cluster dwf-etcd1=https://10.9.0.1:2380,dwf-etcd2=https://10.9.0.2:2380,dwf-etcd3=https://10.10.0.1:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --client-cert-auth \
  --trusted-ca-file /etcd-ssl-certs-dir/etcd-root-ca.pem \
  --cert-file /etcd-ssl-certs-dir/dwf-etcd2.pem \
  --key-file /etcd-ssl-certs-dir/dwf-etcd2-key.pem \
  --peer-client-cert-auth \
  --peer-trusted-ca-file /etcd-ssl-certs-dir/etcd-root-ca.pem \
  --peer-cert-file /etcd-ssl-certs-dir/dwf-etcd2.pem \
  --peer-key-file /etcd-ssl-certs-dir/dwf-etcd2-key.pem

ExecStop=/usr/bin/docker stop etcd-3.2.24

[Install]
WantedBy=multi-user.target
EOF

sudo mv /tmp/dwf-etcd2.service /etc/systemd/system/dwf-etcd2.service


# to start service
sudo systemctl daemon-reload
sudo systemctl cat dwf-etcd2.service
sudo systemctl enable dwf-etcd2.service
sudo systemctl start dwf-etcd2.service

# to get logs from service
sudo systemctl status dwf-etcd2.service -l --no-pager
sudo journalctl -u dwf-etcd2.service -l --no-pager|less
sudo journalctl -f -u dwf-etcd2.service

# to stop service
sudo systemctl stop dwf-etcd2.service
sudo systemctl disable dwf-etcd2.service
```

dwf-k8s-node10-001:

```bash
# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /data/etcd/dwf-etcd3


# to write service file for etcd with Docker
cat > /tmp/dwf-etcd3.service <<EOF
[Unit]
Description=etcd with Docker
Documentation=https://github.com/coreos/etcd

[Service]
Restart=always
RestartSec=5s
TimeoutStartSec=0
LimitNOFILE=40000

ExecStart=/usr/bin/docker \
  run \
  --rm \
  --net=host \
  --name etcd-3.2.24 \
  --volume=dwf/dwf-k8s-etcd3:/etcd-data \
  --volume=/data/etcd/certs:/etcd-ssl-certs-dir \
  k8s.gcr.io/etcd:3.2.24 \
  /usr/local/bin/etcd \
  --name dwf-etcd3 \
  --data-dir /etcd-data \
  --listen-client-urls https://10.10.0.1:2379 \
  --advertise-client-urls https://10.10.0.1:2379 \
  --listen-peer-urls https://10.10.0.1:2380 \
  --initial-advertise-peer-urls https://10.10.0.1:2380 \
  --initial-cluster dwf-etcd1=https://10.9.0.1:2380,dwf-etcd2=https://10.9.0.2:2380,dwf-etcd3=https://10.10.0.1:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --client-cert-auth \
  --trusted-ca-file /etcd-ssl-certs-dir/etcd-root-ca.pem \
  --cert-file /etcd-ssl-certs-dir/dwf-etcd3.pem \
  --key-file /etcd-ssl-certs-dir/dwf-etcd3-key.pem \
  --peer-client-cert-auth \
  --peer-trusted-ca-file /etcd-ssl-certs-dir/etcd-root-ca.pem \
  --peer-cert-file /etcd-ssl-certs-dir/dwf-etcd3.pem \
  --peer-key-file /etcd-ssl-certs-dir/dwf-etcd3-key.pem

ExecStop=/usr/bin/docker stop etcd-3.2.24

[Install]
WantedBy=multi-user.target
EOF

sudo mv /tmp/dwf-etcd3.service /etc/systemd/system/dwf-etcd3.service


# to start service
sudo systemctl daemon-reload
sudo systemctl cat dwf-etcd3.service
sudo systemctl enable dwf-etcd3.service
sudo systemctl start dwf-etcd3.service

# to get logs from service
sudo systemctl status dwf-etcd3.service -l --no-pager
sudo journalctl -u dwf-etcd3.service -l --no-pager|less
sudo journalctl -f -u dwf-etcd3.service

# to stop service
sudo systemctl stop dwf-etcd3.service
sudo systemctl disable dwf-etcd3.service
```


## Install etcdctl

```bash
ETCD_VER=v3.2.24

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

/tmp/etcd-download-test/etcd --version
ETCDCTL_API=3 /tmp/etcd-download-test/etcdctl version
mv /tmp/etcd-download-test/etcdctl  /usr/local/bin/
```

```bash
export ETCDCTL_API=3
export ETCDCTL_DIAL_TIMEOUT=3s
export ETCDCTL_CACERT=/data/etcd/certs/etcd-root-ca.pem
export ETCDCTL_CERT=/data/etcd/certs/dwf-etcd1.pem
export ETCDCTL_KEY=/data/etcd/certs/dwf-etcd1-key.pem
export ETCDCTL_ENDPOINTS="https://10.9.0.1:2379"
```


## Troubleshooting

- Error:  context deadline exceeded

    有几下几种可能:
    - `initial-advertise-peer-urls` 和 `initial-cluster` 地址不一致
    - `ETCDCTL_ENDPOINTS` 不正确
    - `ETCDCTL_CACERT` `ETCDCTL_KEY` `ETCDCTL_CERT` 没有配置或不正确

## Reference

- [Official Installing Play](http://play.etcd.io/install#top)
- [Glusterfs Volume Plugin](https://github.com/trajano/docker-volume-plugins/tree/master/glusterfs-volume-plugin)

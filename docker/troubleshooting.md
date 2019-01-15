- ERROR: for $image_name b'Driver devicemapper failed to remove root filesystem 890ee25a29bc7b701d625c383850b31ee5a3d4d999e38c06cac94aa9cf175837: remove /var/lib/docker/devicemapper/mnt/d4b2449574d4fe653474a4c665a3c375e644b4568522c60b70176ea22b5ab005: device or resource busy'
    ```
    # find /proc/*/mounts | xargs grep d4b2449574d4fe653474a4c665a3c375e644b4568522c60b70176ea22b5ab005
    /proc/4351/mounts:/dev/mapper/docker-8:1-2149433408-d4b2449574d4fe653474a4c665a3c375e644b4568522c60b70176ea22b5ab005 /var/lib/docker/devicemapper/mnt/d4b2449574d4fe653474a4c665a3c375e644b4568522c60b70176ea22b5ab005 xfs rw,relatime,nouuid,attr2,inode64,logbsize=64k,sunit=128,swidth=128,noquota 0 0

    # ps aux|grep 4351
    ntp       4351  0.0  0.0  29884  2120 ?        Ss   Dec20   0:17 /usr/sbin/ntpd -u ntp:ntp -g

    #systemctl restart ntpd
    ```

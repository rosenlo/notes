# Upgrading to python 2.7 on Centos 6.6

- First, let's download and install python 2.7

	```
	[root@Rosen tmp]# wget --no-check-certificate https://www.python.org/ftp/python/2.7.10/Python-2.7.10.tar.xz
	[root@Rosen tmp]# tar xf Python-2.7.10.tar.xz
	[root@Rosen tmp]# cd Python-2.7.10
	[root@Rosen tmp]# ./configure --prefix=/usr/local
	[root@Rosen tmp]# make && make altinstall
	```
	**⚠️Warning：** It is important to use `make altinstall` instead of `make install` otherwise you will end up with two diffrent versions of Python on your filesystem, both named `python`.


- Create a soft link

	```
	[root@Rosen tmp]# ln -s /usr/local/bin/python2.7 /usr/local/bin/python
	[root@Rosen tmp]# ls -ltr /usr/local/bin/python*
	-rwxr-xr-x 1 root root 6225045 Oct 25 03:35 /usr/local/bin/python2.7
	-rwxr-xr-x 1 root root    1687 Oct 25 03:35 /usr/local/bin/python2.7-config
	lrwxrwxrwx 1 root root      24 Oct 25 03:36 /usr/local/bin/python -> /usr/local/bin/python2.7
	```
- Check python version

	```
	[root@Rosen tmp]# /usr/local/bin/python2.7 -V
	Python 2.7.10
	```

- Install easy_install

	```
	[root@Rosen tmp]# wget --no-check-certificate https://bootstrap.pypa.io/ez_setup.py
	[root@Rosen tmp]# /usr/local/bin/python2.7 ez_setup.py
	```
		
That's it! Enjoy.

---

**Reference：**
>1. [Difference in details between “make install” and “make altinstall”](https://stackoverflow.com/questions/16018463/difference-in-details-between-make-install-and-make-altinstall?answertab=votes#tab-top)


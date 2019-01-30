# Python 标准库学习

<!-- vim-markdown-toc GFM -->

* [os](#os)
    * [os.walk](#oswalk)
    * [os.listdir](#oslistdir)
    * [os.path.normcase](#ospathnormcase)
    * [os.stat](#osstat)
* [fnmatch](#fnmatch)
    * [fnmatch.fnmatch](#fnmatchfnmatch)

<!-- vim-markdown-toc -->


## os

### os.walk

os.walk(top, topdown=True, onerror=None, followlinks=False)

默认从上之下遍历所有文件并生成一个目录树，返回一个元组包含了三个元素：**根路径**、**子目录**、**文件名**

```python
In [1]: import os

In [2]: for root, dirs, files in os.walk('/tmp/testdir/'):
    ...:     print("{} {} {}".format(root, dirs, files))
    ...:
/tmp/testdir/ ['dir2', 'dir3', 'dir1'] ['file3', 'file4', 'file5', 'file2', 'file1']
/tmp/testdir/dir2 [] []
/tmp/testdir/dir3 [] []
/tmp/testdir/dir1 [] ['dir1_file1', 'dir1_file2', 'dir1_file3']
```

### os.listdir

os.listdir(path='.')

返回一个指定路径下所有文件名的列表，包括目录，不包括特殊目录 `.` `..`

```python
In [1]: import os

In [2]: os.listdir('/tmp/testdir')
Out[2]: ['file3', 'file4', 'file5', 'file2', 'dir2', 'dir3', 'file1', 'dir1']
```

### os.path.normcase

os.path.normcase(path)

- 标准化路径名
- 在 Unix 和 Mac OS X 系统，原封不动返回，在大小写不区分的系统，会转成小写
- 在 Windows 系统，它将正斜杠转换成反斜杠
- 如果 path 的类型不是 `str` 或 `bytes` 会引发 `TypeError` 异常

### os.stat

os.stat(path)

运行系统层级的 stat() 调用， 返回 `posix.stat_result` 对象

- st_mode - protection bits
- st_ino - inode number
- st_dev - device
- st_nlink - number of hard links
- st_uid - user id of owner
- st_gid - group id of owner
- st_size - size of file, in bytes
- st_atime - time of most recent access
- st_mtime - time of most recent content modification
- st_ctime - platform dependent; time of most recent metadata change on Unix, or the time of creation on Windows
- st_blocks - number of 512-byte blocks allocated for file
- st_blksize - filesystem blocksize for efficient file system I/O
- st_rdev - type of device if an inode device
- st_flags - user defined flags for file
- st_gen - file generation number
- st_birthtime - time of file creation
- st_ftype (file type)
- st_attrs (attributes)
- st_obtype (object type)


## fnmatch

### fnmatch.fnmatch

fnmatch.fnmatch(filename, pattern)

- 测试文件名是否匹配 pattern ，返回 `True` 和 `False`
- 这两个参数都使用 [os.path.normcase](#ospathnormcase) 大小写标准

```python
In [1]: import os, fnmatch

In [2]: pattern = "*.log"

In [3]: for name in os.listdir('/tmp/testdir/dir2'):
    ...:     if fnmatch.fnmatch(name, pattern):
    ...:         print(name)
    ...:
access1.log
access2.log
access3.log
```

# Linux

## 一、基础概念

#### iowait
先看下官方解释`man iostat`：

```
%iowait
    Show the percentage of time that the CPU or CPUs were idle during which the system had an outstanding disk I/O request.
```

意思是说CPUs处于**空闲时**等待系统磁盘I/O时等待的时间以百分数表示。


对%iowait的值升高有两个常见的误解：
- CPU不能工作
- I/O有瓶颈

正确认识：
- 第一个误解很显然是错的，因为只有CPU处于空闲时%iowait才会升高，%iowait是%idle的子类(sub-category)，此时CPU还是可以继续工作。
- 第二个误解很常见，认为%iowait升高，其他进程都在休眠，等待I/O的时间也更长了。听上去挺有道理，但不看全局只看一个指标就有点耍流氓的意思了，首先%iowait升高的前提是CPU在idle的情况下磁盘I/O，那需要看结合其他指标如：应用的I/O并发、系统的I/O请求量，可以通过iostat命令查看（avgrq-sz、avgqu-sz、await）等指标有没有明显增大。


参考：
- https://linux.die.net/man/1/iostat
- http://man7.org/linux/man-pages/man5/proc.5.html

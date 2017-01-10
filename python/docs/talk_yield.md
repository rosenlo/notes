## 谈谈Python的生成器
> 来源：思诚之道  
> 链接：www.bjhee.com/python-yield.html

第一次看到Python代码中出现yield关键字时，一脸懵逼，完全理解不了这个。网上查下解释，函数中出现了yield关键字，则调用该函数时会返回一个生成器。那到底什么是生成器呢？我们经常看到类似下面的代码

```python
def count(n):
    x = 0
    while x < n:
        yield x
        x += 1
 
for i in count(5):
    print i
```
这段代码执行后打印序列0到4，所以我一开始以为这个生成器就是生成一个序列呀。那这跟迭代器有什么区别呢？我们来看下迭代器的例子：

```python
class CountIter:
    def __init__(self, n):
        self.n = n
 
    def __iter__(self):
        self.x = -1
        return self
 
    def next(self):  # For Python 2.x
        self.x += 1
        if self.x < self.n:
            return self.x
        else:
            raise StopIteration
 
for i in CountIter(5):
    print i
```
CountIter类就是一个迭代器，它的`__iter__()`方法返回可迭代对象，next()方法则执行下一轮迭代（注：在Python 3.x里是`__next__()`方法）。上面的代码执行后也会打印序列0到4，看上去跟之前的生成器效果一样，就是代码长一点。不仅如此，生成器自带next()方法，而且在越界时也会抛出`StopIteration`异常。

```python
gen = count(2)
print gen.next() # 0
print gen.next() # 1
print gen.next() # StopIteration
```
那区别到底是什么，在何种情况下，我们应该使用生成器呢？

每次执行迭代器的next()方法并返回后，该方法的上下文环境即消失了，也就是所有在next()方法中定义的局部变量就无法被访问了。而对于生成器，每次执行next()方法后，代码会执行到yield关键字处，并将yield后的参数值返回，同时当前生成器函数的上下文会被保留下来。也就是函数内所有变量的状态会被保留，同时函数代码执行到的位置会被保留，感觉就像函数被暂停了一样。当再一次调用next()方法时，代码会从yield关键字的下一行开始执行。很神奇吧！如果执行next()时没有遇到yield关键字即退出（或返回），则抛出`StopIteration`异常。

本文的第一个例子是使用生成器函数来构造生成器，Python也提供了生成器表达式，下面的例子也可以打印序列0到4。

```python
gen = (x for x in range(5))  # 注意这里是()，而不是[]
for i in gen:
    print i
```

到目前为止，我们了解了生成器同迭代器在实现机制上的不同，但似乎功能是一样的，那生成器的存在有什么价值呢？我们先来看看除了next()方法外，生成器还提供了哪些方法。

1. `close()`方法  
	顾名思义，close()方法就是关闭生成器。生成器被关闭后，再次调用next()方法，不管能否遇到yield关键字，都会立即抛出`StopIteration`异常。
	
	```python
	gen = (x for x in range(5))
	gen.close()
	gen.next()  # StopIteration
	```
2. `send()`方法  
	- 这是我认为生成器最重要的功能，我们可以通过send()方法，向生成器内部传递参数。我们来看个例子：
	
	```python
	def count(n):
    x = 0
    while x < n:
        value = yield x
        if value is not None:
            print 'Received value: %s' %value
        x += 1
	```
	- 还是之前的count函数，唯一的区别是我们将**”yield x”**的值赋给了变量`value`，并将其打印出来。如何给`value`传值呢？
	
	```python
	gen = count(5)
	print gen.next()  # print 0
	print gen.send('Hello')  # Received value: Hello, then print 1

	```
	- 我们先调用next()方法，让代码执行到yield关键字（这步必须要），当前打印出0。然后当我们调用`”gen.send(‘Hello’)”`时，字符串’Hello’就被传入生成器中，并作为yield关键字的执行结果赋给变量”value”，所以控制台会打印出**”Received value: Hello”**。然后代码继续执行，直到下一次遇到yield关键字后暂定，此时生成器返回的是**1**。
	- 简单的说，send()就是next()的功能，加上传值给yield。如果你有兴趣看下Python的源码，你会发现，其实next()的实现，就是send(None)。
3. `throw()`方法  
	- 除了向生成器函数内部传递参数，我们还可以传递异常。还是先看例子：
	
	```python
	def throw_gen():
    try:
        yield 'Normal'
    except ValueError:
        yield 'Error'
    finally:
        print 'Finally'
        
	gen = throw_gen()
	print gen.next()  # Normal
	print gen.next()  # Finally, then StopIteration

	```
	- 如果像往常一样调用next()方法，会返回’Normal’。再次调用next()，会进入finally语句，打印’Finally’，同时由于函数退出，生成器会抛出`StopIteration`异常。我们换个方式，在第一次调用next()方法后，调用throw()方法，情况会怎样？
	
	```python
	gen = throw_gen()
	print gen.next()  # Normal
	print gen.throw(ValueError)    # Error
	print gen.next()  # Finally, then StopIteration
	```
	- 我们会看到，throw()方法向生成器函数内部传递了`ValueError`异常，代码进入`except ValueError`语句，当遇到下一个yield时才暂停并退出，此时生成器返回的是’Error’字符串。简单的说，throw()就是next()的功能，加上传异常给yield。

	- 聊到这里，相信大家对生成器的功能已经有了一个很好的理解。生成器不但可以逐步生成序列，不用像列表一样初始化时就要开辟所有的空间。它更大的价值，我个人认为，就是模拟并发。很多朋友可能已经知道，Python虽然可以支持多线程，但由于`GIL`（全局解释锁，Global Interpreter Lock）的存在，同一个时间，只能有一个线程在运行，所以无法实现真正的并发。我们暂且不讨论GIL存在的意义，这里我们提出了一个新的概念，就是**协程**（`Coroutine`）。

	- Python实现协程最简单的方法，就是使用yield。当一个函数在执行过程中被阻塞时，就用yield挂起，然后执行另一个函数。当阻塞结束后，可以用next()或者send()唤醒。相比多线程，协程的好处是它在一个线程内执行，避免线程之间切换带来的额外开销，而且多线程中使用共享资源，往往需要加锁，而协程不需要，因为代码的执行顺序是你完全可以预见的，不存在多个线程同时写某个共享变量而导致出错的情况。
	- 我们来使用协程写一个生产者消费者的例子：
	
	```python
	def consumer():
    last = ''
    while True:
        receival = yield last
        if receival is not None:
            print 'Consume %s' % receival
            last = receival
            
	def producer(gen, n):
	    gen.next()
	    x = 0
	    while x < n:
	        x += 1
	        print 'Produce %s' % x
	        last = gen.send(x)
	    gen.close()
	    
	gen = consumer()
	producer(gen, 5)
	```
	- 执行下例子，你会看到控制台交替打印出生产和消费的结果。消费者consumer()函数是一个生成器函数，每次执行到yield时即挂起，并返回上一次的结果给生产者。生产者producer()接收到生成器的返回，并生成一个新的值，通过send()方法发送给消费者。至此，我们成功实现了一个（伪）并发。

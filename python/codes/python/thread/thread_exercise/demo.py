#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: demo.py
@time: 10/9/16 5:32 PM
"""

from threading import Thread
import time

Thread.run()

def Foo(arg):
    for item in range(10):
        print(item)
        time.sleep(1)


print('before')
t1 = Thread(target=Foo, args=('aaa',))
#t1.setDaemon(True)
t1.start()
t1.join()

print('after')

time.sleep(5)

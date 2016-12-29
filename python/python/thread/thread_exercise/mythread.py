#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: test.py
@time: 10/9/16 6:31 PM
"""

from threading import Thread
import time


class MyThread(Thread):
    def run(self):
        Thread.run(self)


def Foo():
    print('Foo')


t1 = MyThread(target=Foo)
t1.start()
print('end')

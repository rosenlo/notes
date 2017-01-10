#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: mtl5.py
@time: 10/24/16 8:32 PM
"""
import time
from multiprocessing import Process, Lock


def f(l, i):
	l.acquire()
	time.sleep(1)
	print('hello world', i)
	l.release()


if __name__ == '__main__':
	lock = Lock()
	for num in range(10):
		Process(target=f, args=(lock, num)).start()

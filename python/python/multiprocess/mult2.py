#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: mtprocess2.py
@time: 10/19/16 9:50 PM
"""

from multiprocessing import Process
import time
import os


def info(title):
	print(title)
	print("module name:", __name__)
	print("parent process id:", os.getppid())
	time.sleep(5)
	print("process id:", os.getpid())


def f(name):
	info("function f")
	print("hello ", name)


if __name__ == '__main__':
	info("main line")
	print("--------------")
	p = Process(target=f, args=("rosen",))
	p.start()
	p.join()

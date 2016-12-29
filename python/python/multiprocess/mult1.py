#!/usr/bin/env python
# -*- coding=utf-8 -*-


from multiprocessing import Pool
import time


def f(n):
	print(time.ctime())
	time.sleep(0.5)
	return n * n


p = Pool(2)

"""
第一种写法
"""
print(p.map(f, [1, 2, 3]))
print map(f, [1, 2, 3])

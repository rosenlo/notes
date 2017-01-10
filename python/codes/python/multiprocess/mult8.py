#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: mult8.py
@time: 10/25/16 10:13 AM
"""
from multiprocessing import Pool
import time


def f(x):
	# print(x * x)
	time.sleep(1)
	return x * x


pool = Pool(processes=4)
res_list = []

for i in range(10):
	"""
	第三种写法
	"""
	res = pool.apply_async(f, args=(i,))
	res_list.append(res)

for i in res_list:
	print(i.get())

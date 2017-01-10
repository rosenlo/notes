#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: mlt6.py
@time: 10/25/16 8:25 AM
"""
from multiprocessing import Process, Value, Array


def f(n, a):
	n.value = 3.1415
	for i in range(len(a)):
		a[i] = -a[i]


if __name__ == '__main__':
	num = Value('d', 0.0)
	arr = Array('i', range(10))

	p = Process(target=f, args=(num, arr))
	p.start()
	p.join()

	print(num.value)
	print(arr[:])

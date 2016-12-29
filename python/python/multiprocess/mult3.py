#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: mtp3.py
@time: 10/20/16 8:19 AM
"""

from multiprocessing import Process, Queue


def f(q, n):
	q.put([n, u"hello"])


if __name__ == '__main__':
	q = Queue()
	q.put("aaa")
	for i in range(5):
		"""
		第二种写法
		"""
		p = Process(target=f, args=(q, i))
		p.start()
		print(q.get())
		p.join()

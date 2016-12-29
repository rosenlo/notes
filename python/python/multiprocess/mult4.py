#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: mtl4.py
@time: 10/20/16 6:46 PM
"""
from multiprocessing import Process, Pipe


def f(conn):
	conn.send([42, None, "hello"])
	conn.close()


if __name__ == '__main__':
	parent_conn, child_conn = Pipe()
	p = Process(target=f, args=(child_conn,))
	p.start()
	print parent_conn.recv()
	p.join()

#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: mythread5.py
@time: 10/13/16 10:34 AM
"""
import threading
import time

num = 0


def run():
    time.sleep(1)
    global num
    lock.acquire()
    time.sleep(0.001)
    num += 1
    print(num)
    lock.release()


lock = threading.BoundedSemaphore(4)

for i in range(200):
    t = threading.Thread(target=run)
    t.start()

if __name__ == '__main__':
    pass

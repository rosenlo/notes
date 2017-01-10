#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: thread_event.py
@time: 10/13/16 9:05 PM
"""
import threading
import time


def producer():
    print(u'等人来买包子')
    event.wait()
    print(u'chef: sb is coming for baozi...')
    print(u'chef: making a baozi for sb...')
    time.sleep(5)
    print(u'chef: your baozi is done')
    event.set()


def consumer():
    time.sleep(2)
    print(u'去买包子')
    event.set()
    event.clear()
    time.sleep(2)
    print(u'rosen: waiting for baozi to be ready...')
    while True:
        if event.isSet():
            print(u'thanks...')
            break
        else:
            print(u'so slow')
            time.sleep(1)


event = threading.Event()
p = threading.Thread(target=producer)
c = threading.Thread(target=consumer)

p.start()
c.start()

if __name__ == '__main__':
    pass

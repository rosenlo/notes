#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: producer.py
@time: 10/11/16 10:45 AM
"""

from threading import Thread
from Queue import Queue
import time


class Producer(Thread):
    def __init__(self, name, queue):
        self.__name = name
        self.__queue = queue
        super(Producer, self).__init__()

    def run(self):
        try:
            while True:
                if self.__queue.full():
                    time.sleep(1)
                else:
                    self.__queue.put('test')
                    time.sleep(0.5)
                    print('{}生产了一条消息'.format(self.__name))
        except KeyboardInterrupt as e:
            print(e)
            exit()


class Consumer(Thread):
    def __init__(self, name, queue):
        self.__name = name
        self.__queue = queue
        super(Consumer, self).__init__()

    def run(self):
        try:
            while True:
                if self.__queue.empty():
                    time.sleep(1)
                else:
                    self.__queue.get()
                    #print self.__queue.qsize()
                    time.sleep(1)
                    print('{}消费了一条消息'.format(self.__name))
        except KeyboardInterrupt as e:
            print(e)
            exit()


def main():
    q = Queue(maxsize=100)

    for i in range(1):
        name = 'Producer-%d' % i
        work = Producer(name, q)
        work.start()

    for i in range(20):
        name = 'Consumer-%d' % i
        work = Consumer(name, q)
        work.start()


if __name__ == '__main__':
    main()

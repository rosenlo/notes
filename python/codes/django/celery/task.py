#!/usr/bin/env python
# encoding: utf-8

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: task.py
Created Time: 3/22/17 12:17
"""
from tasks import add

if __name__ == '__main__':
    res = add.delay(4, 10)
    print(dir(res))
    print(res.state)

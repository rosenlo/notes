#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: greenlet_client.py
Created Time: Thu Jan 12 17:42:18 2017
"""

from socket import *

addr = ('localhost', 22222)
client = socket(AF_INET, SOCK_STREAM)
client.connect(addr)

while 1:
    cmd = raw_input(">>:").strip()
    if len(cmd) == 0:
        continue
    client.send(str(cmd))
    # data = client.recv(1024)
    # print(data)
    if cmd == 'exit':
        client.close()
        break

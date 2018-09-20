#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: greenlet_server.py
Created Time: Thu Jan 12 17:45:15 2017
"""

import socket

import gevent
from gevent import monkey

monkey.patch_all()


def server(port):
    sock = socket.socket()
    sock.bind(('127.0.0.1', port))
    sock.listen(50)
    while 1:
        conn, addr = sock.accept()
        gevent.spawn(handle_request, conn)


def handle_request(conn):
    try:
        while 1:
            data = conn.recv(1024)
            if not data:
                break
            print('recv:', data)
            conn.send(data)
            if data == 'exit':
                conn.close()
                break
    except Exception as e:
        print(e)
    finally:
        conn.close()


if __name__ == '__main__':
    server(22222)

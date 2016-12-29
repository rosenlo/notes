#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: client.py
@time: 8/28/16 12:00 PM
"""

import socket


def main():
    client = socket.socket()
    ip_port = ('127.0.0.1', 9999)
    client.connect(ip_port)

    while True:
        data = client.recv(1024)
        print(data)
        inp = raw_input('client:')
        client.send(inp)
        if inp == 'exit':
            break


if __name__ == '__main__':
    main()

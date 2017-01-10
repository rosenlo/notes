#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: server.py
@time: 8/28/16 12:00 PM
"""
import socket


def main():
    sk = socket.socket()
    ip_port = ('127.0.0.1', 9999)
    sk.bind(ip_port)
    sk.listen(5)

    while True:
        conn, address = sk.accept()
        Flag = True
        conn.send('hello.')
        while Flag:
            data = conn.recv(1024)
            print(data)
            if data == 'exit':
                Flag = False
            conn.send('sb')
        conn.close()


if __name__ == '__main__':
    main()

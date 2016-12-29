#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: server.py
@time: 9/11/16 6:17 PM
"""

import SocketServer


class MyServer(SocketServer.BaseRequestHandler):
    def setup(self):
        pass

    def handle(self):
        print(self.request, self.client_address, self.server)
        conn = self.request
        conn.send('hello')
        Flag = True
        while Flag:
            data = conn.recv(1024)
            print(data)
            if data == 'exit':
                Flag = False
            conn.send('sb')
        conn.close()

    def finish(self):
        pass


if __name__ == '__main__':
    server = SocketServer.ThreadingTCPServer(('127.0.0.1', 9999), MyServer)
    server.serve_forever()

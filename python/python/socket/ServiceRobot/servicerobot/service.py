#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: main.py
@time: 9/20/16 10:51 AM
"""
import SocketServer
from admin.db import RobotDB


class ServiceRobot(SocketServer.BaseRequestHandler):
    def setup(self):
        pass

    def handle(self):
        print(self.request, self.server, self.client_address)
        name = 'Service Robot'
        sender = name
        # sql_data = []
        db = RobotDB()
        conn = self.request
        user_name = conn.recv(1024)
        xxx = conn.recv(1024)
        recevier = user_name
        text = '{0}: 你好{1}, 我是{0}'.format(name, user_name)
        conn.send(text)
        # sql_data.append((name, sender, recevier, text))
        sql_data = (name, sender, recevier, text)
        db.Insert(*sql_data)
        Flag = True
        while Flag:
            data = conn.recv(1024)
            #db.Insert(name, user_name, name, data)
            # sql_data.append((name, sender, recevier, text))
            print(data)
            if data == 'exit':
                Flag = False
            else:
                text = 'sb'
                conn.send('{}: '.format(name) + text)
                db.Insert(name, sender, recevier, text)
        conn.close()

    def finish(self):
        pass


if __name__ == '__main__':
    service = SocketServer.ThreadingTCPServer(('127.0.0.1', 9999), ServiceRobot)
    service.serve_forever()

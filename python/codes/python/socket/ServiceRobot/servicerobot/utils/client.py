#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: client.py
@time: 8/28/16 12:00 PM
"""

from admin.db import RobotDB
import socket


class Client(object):
    def __init__(self, ip, port):
        self.client = socket.socket()
        self.ip_port = (ip, port)
        self.client.connect(self.ip_port)
        self.user_name, self.get_records = None, None
        self.choose = None
        try:
            self.handle()
        finally:
            self.finish()

    def setup(self):
        print('''
        1) 开始聊天
        2) 查看聊天记录
        ''')
        while True:
            self.choose = raw_input('选择: ')
            if self.choose in ['1', '2']:
                return self.choose

    def handle(self):
        self.user_name = raw_input('请输入用户名:')

        choose = self.setup()
        db = RobotDB()
        if choose == '2':
            self.get_records = db.GetRecords(self.user_name, 'Server Robot')
        elif choose == '1':
            self.client.send(self.user_name)
            self.client.send('aaaa')
            self.client.send('bbb')
        # self.client.send(self.start_chat)
        while True:
            data = self.client.recv(1024)
            print(data)
            inp = ''
            while not inp:
                inp = raw_input('{}:'.format(self.user_name))
            sql_data = (self.user_name, self.user_name, 'Service Robot', inp)
            self.client.send(inp)
            db.Insert(*sql_data)
            if inp == 'exit':
                break

    def finish(self):
        pass

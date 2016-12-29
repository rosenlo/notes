#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: min.py
@time: 9/22/16 4:44 PM
"""
from utils.mysqlhelper import MySQLHelper


class RobotDB(object):
    def __init__(self):
        self.__helper = MySQLHelper()

    def Insert(self, username, sender, recevier, text):
        sql = 'INSERT INTO chat_records (username, sender, recevier, text) VALUES (%s, %s, %s, %s)'
        params = (username, sender, recevier, text)
        self.__helper.Insert(sql, params)

    def GetRecords(self, username, name):
        sql = 'SELECT * FROM chat_records WHERE sender in ({}, {})'.format(username, name)


#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: mysqlhelper.py
@time: 8/27/16 5:44 PM
"""

import MySQLdb
from conf import settings


class MySQLHelper(object):
	def __init__(self):
		self.host = settings.HOST
		self.db = settings.DB
		self.user = settings.USER
		self.passwd = settings.PASSWD

	def insert_and_update(self, sql, params):
		conn = MySQLdb.connect(self.host, self.user, self.passwd, self.db)
		cur = conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)

		res = cur.execute(sql, params)
		conn.commit()

		cur.close()
		conn.close()
		return res

	def get_data(self, sql, params, size):
		conn = MySQLdb.connect(self.host, self.user, self.passwd, self.db)
		cur = conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)

		cur.execute(sql, params)
		data = cur.fetchall() if size == 'all' else cur.fetchone()

		cur.close()
		conn.close()

		return data

	def update(self, sql, params):
		conn = MySQLdb.connect(self.host, self.user, self.passwd, self.db)
		cur = conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)

		cur.execute(sql, params)
		conn.commit()

		cur.close()
		conn.close()

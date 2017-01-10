#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: login.py
@time: 10/12/16 4:05 PM
"""

from admin.db import FortressesDB


class LoginAuth(object):
	def __init__(self):
		self.name = self.func_input('username')
		self.password = self.func_input('password')
		self.__db = FortressesDB()
		self.host_list = {}

	@staticmethod
	def func_input(input_name='user_name'):
		return raw_input("Please input your %s:" % input_name).strip()

	def user_check(self):
		res = self.__db.user_auth(self.name, self.password)
		if res:
			print(self.__db.update(username=self.name))
		return res

	def show_hosts(self):
		res = self.__db.display_hosts(self.name)
		host_list = {}
		for i in res:
			host_list[i['hostname']] = i['host_ip']

		num = 0
		for t, host in host_list.items():
			num += 1
			self.host_list[str(num)] = host
			print("""
        {}: {}({})""".format(num, t, host))

		return str(num)

	def choose_num(self):
		return raw_input("Please choose your hosts(if you want to exit the enter 'exit'): ")


if __name__ == '__main__':
	pass

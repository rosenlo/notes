#!/usr/bin/env python
# encoding: utf-8

"""
@file: common.py
@time: 10/31/16 20:09
@author: Rosen 
"""

from admin.db import FortressDB


def display_menu():
	while True:
		print("""
		1) display your hosts
		2) quit
		""")
		return raw_input('Please input your choose: ')


def choose_num():
	return raw_input("Please choose your hosts(if you want to exit the enter 'q' or 'quit'): ")


class ShowHostList(object):
	def __init__(self, username):
		self.name = username
		self.__db = FortressDB()
		self.host_list = {}

	def show_hosts(self):
		res = self.__db.display_hosts(self.name)
		print(res)
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

	def get_host_user(self, host_ip):
		

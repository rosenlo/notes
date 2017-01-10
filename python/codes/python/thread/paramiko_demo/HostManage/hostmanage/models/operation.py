#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: operation.py
@time: 10/14/16 5:24 PM
"""
import paramiko


class operation(object):
	def __init__(self, name, private_key_path):
		self.__ssh = paramiko.SSHClient()
		self.__key = paramiko.RSAKey.from_private_key_file(private_key_path)
		self.username = name
		self.port = 22

	def output(self, cmd):
		stdin, stdout, stderr = self.__ssh.exec_command(cmd)
		for i in [stdout, stderr]:
			if i:
				print(i.read())

	def once(self, host):
		self.__ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
		self.__ssh.connect(host, self.port, self.username, pkey=self.__key)
		Flag = True
		while Flag:
			cmd = self.loop_cmd()
			self.output(cmd)
			if cmd == 'exit':
				self.__ssh.close()
				Flag = False

	def batch(self, host, cmd):
		self.__ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
		self.__ssh.connect(host, self.port, self.username, pkey=self.__key)
		self.output(cmd)
		self.__ssh.close()

	@staticmethod
	def loop_cmd():
		while True:
			inp = raw_input('Please input the command to be executed(if you want to exit the enter "exit"): ')
			if inp == 'exit':
				return inp
			return inp


if __name__ == '__main__':
	pass

#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: main.py
@time: 10/12/16 4:49 PM
"""

from models import login
from models.operation import operation
import threading


def display_host():
	while True:
		print("""
		1) display your hosts
		2) quit
		""")
		return raw_input('Please input your choose: ')


def thread_lock(threads):
	return threading.BoundedSemaphore(threads)


def input_threads():
	while True:
		try:
			return input('Please specify how many threads: ')
		except (NameError, SyntaxError):
			pass


def alone_host(choose_num, login_auth, private_key_path):
	op = operation('root', private_key_path)
	host = login_auth.host_list["".join(choose_num)]
	op.once(host)


def multiple_host(choose_num, login_auth, private_key_path, host_num):
	cmd = raw_input('Please input the command to be executed(if you want to exit the enter "exit"): ')
	threads = input_threads()
	for num in choose_num:
		lock = thread_lock(threads)
		if num <= host_num:
			lock.acquire()

			host = login_auth.host_list[num]
			op = operation('root', private_key_path)
			t = threading.Thread(target=op.batch, args=(host, cmd))
			t.start()
			t.join(1)

			lock.release()
		else:
			exit('The host number is incorrect.')


def main():
	# retry = 0
	# while retry < 3:
	private_key_path = '/Users/Rosen/.ssh/id_rsa'
	login_auth = login.LoginAuth()
	res = login_auth.user_check()
	while True:
		if res:
			opt = display_host()
			if opt == '1':
				host_num = login_auth.show_hosts()
				choose_num = login_auth.choose_num().split(',')

				if host_num == 'exit' or "".join(choose_num) == 'exit':
					exit()
				elif choose_num == ['']:
					continue

				if len(choose_num) == 1 and choose_num <= host_num:
					alone_host(choose_num, login_auth, private_key_path)
				elif len(choose_num) > 1 and choose_num <= host_num:
					multiple_host(choose_num, login_auth, private_key_path, host_num)
			elif opt == '2':
				exit()
		else:
			exit('Authenticate failed.')


if __name__ == '__main__':
	main()

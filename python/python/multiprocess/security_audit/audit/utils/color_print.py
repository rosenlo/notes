#!/usr/bin/env python
# encoding: utf-8

"""
@file: color_print.py
@time: 10/31/16 10:39
@author: Rosen 
"""


def color_print(msg, color='red'):
	color_msg = {'blue': '\033[1;36m%s\033[0m',
				 'green': '\033[1;32m%s\033[0m',
				 'yellow': '\033[1;33m%s\033[0m',
				 'red': '\033[1;31m%s\033[0m',
				 'title': '\033[30;42m%s\033[0m',
				 'info': '\033[32m%s\033[0m'}
	msg = color_msg.get(color, 'red') % msg
	print msg

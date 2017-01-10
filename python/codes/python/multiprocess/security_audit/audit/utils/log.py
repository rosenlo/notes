#!/usr/bin/env python
# encoding: utf-8

"""
@file: log.py
@time: 10/30/16 08:55
@author: Rosen 
"""

import logging

from settings import DEBUG


def log_to_file(filename, level=DEBUG):
	log = logging.getLogger('paramiko')
	if len(log.handlers) > 0:
		return
	log.setLevel(level)
	f = open(filename, 'ab+')
	log_hdlr = logging.StreamHandler(f)
	log_hdlr.setFormatter(
		logging.Formatter('%(levelname)s [%(asctime)s.%(msecs)03d] thr=%(_threadid)-3d %(name)s: %(message)s',
						  '%Y%m%d-%H:%M:%S'))
	log.addHandler(log_hdlr)

#!/usr/bin/env python
# encoding: utf-8

"""
@file: session.py
@time: 10/30/16 14:56
@author: Rosen 
"""

from audit import interactive


def session(t):
    chan = t.open_session()
    chan.get_pty()
    chan.invoke_shell()
    print('*** Here we go!\n')
    interactive.interactive_shell(chan)
    chan.close()
    t.close()

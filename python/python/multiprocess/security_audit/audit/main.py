#!/usr/bin/env python
# encoding: utf-8

"""
@file: main.py
@time: 10/29/16 21:26
@author: Rosen
"""

import getpass
import sys

from audit.module import connect
from audit.settings import PORT
from audit.utils import common
from audit.utils.color_print import color_print
from audit.utils.log import log_to_file

# welcome site
color_print('\n                ***      Welcome to the fortress machine!     ***\n')
# setup logging
username = getpass.getuser()
log_to_file('{}_history.log'.format(username))

"""
username = ''
if len(sys.argv) > 1:
	hostname = sys.argv[1]
else:
	hostname = raw_input('Hostname: ')
if len(hostname) == 0:
	print('*** hostname required.')
	sys.exit(1)

port = PORT
# now connect
sock = connect.create_sock(hostname, port)

connect.ssh_connect(sock, hostname, username)
"""


def main():
    opt = common.display_menu()
    if opt == '1':
        comm = common.ShowHostList(username)
        host_num = comm.show_hosts()
        choose_num = common.choose_num()
        if 'q' in choose_num or 'quit' in choose_num:
            sys.exit(1)
        if choose_num <= host_num:
            host = comm.host_list[choose_num]
            host_user = comm.get_host_user()
            sk = connect.create_sock(host, PORT)
            connect.ssh_connect(sk, host, )

    elif opt == 'q':
        sys.exit(1)


if __name__ == '__main__':
    main()

#!/usr/bin/env python
# encoding: utf-8

"""
@file: connect.py
@time: 10/30/16 14:18
@author: Rosen
"""
import getpass
import os
import socket
import sys
import time
import traceback

import paramiko
from paramiko.py3compat import u

try:
    import termios
    import tty

    has_termios = True
except ImportError:
    print '\033[1;31m仅支持类Unix系统 Only unix like supported.\033[0m'
    has_termios = False


def manual_auth(t, username, hostname):
    default_auth = 'p'
    auth = raw_input('Auth by (p)assword, (r)sa key, or (d)ss key? [%s] ' % default_auth)
    if len(auth) == 0:
        auth = default_auth

    if auth == 'r':
        default_path = os.path.join(os.environ['HOME'], '.ssh', 'id_rsa')
        path = raw_input('RSA key [%s]: ' % default_path)
        if len(path) == 0:
            path = default_path
        try:
            key = paramiko.RSAKey.from_private_key_file(path)
        except paramiko.PasswordRequiredException:
            password = getpass.getpass('RSA key password: ')
            key = paramiko.RSAKey.from_private_key_file(path, password)
        t.auth_publickey(username, key)
    elif auth == 'd':
        default_path = os.path.join(os.environ['HOME'], '.ssh', 'id_dsa')
        path = raw_input('DSS key [%s]: ' % default_path)
        if len(path) == 0:
            path = default_path
        try:
            key = paramiko.DSSKey.from_private_key_file(path)
        except paramiko.PasswordRequiredException:
            password = getpass.getpass('DSS key password: ')
            key = paramiko.DSSKey.from_private_key_file(path, password)
        t.auth_publickey(username, key)
    else:
        pw = getpass.getpass('Password for %s@%s: ' % (username, hostname))
        t.auth_password(username, pw)


def agent_auth(transport, username):
    """
    Attempt to authenticate to the given transport using any of the private
    keys available from an SSH agent.
    """

    agent = paramiko.Agent()
    agent_keys = agent.get_keys()
    if len(agent_keys) == 0:
        return

    for key in agent_keys:
        print('Trying ssh-agent key %s' % hexlify(key.get_fingerprint()))
        try:
            transport.auth_publickey(username, key)
            print('... success!')
            return
        except paramiko.SSHException:
            print('... nope.')


def create_sock(hostname, port):
    try:
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect((hostname, port))
        return sock
    except Exception as e:
        print('*** Connect failed: ' + str(e))
        traceback.print_exc()
        sys.exit(1)


def ssh_connect(sock, hostname, username):
    try:
        t = paramiko.Transport(sock)
        try:
            t.start_client()
        except paramiko.SSHException:
            print('*** SSH negotiation failed.')
            sys.exit(1)

        try:
            keys = paramiko.util.load_host_keys(os.path.expanduser('~/.ssh/known_hosts'))
        except IOError:
            print('*** Unable to open host keys file')
            keys = {}

        key = t.get_remote_server_key()
        if hostname not in keys:
            print('*** WARNING: Unknown host key!')
        elif key.get_name() not in keys[hostname]:
            print('*** WARNING: Unknown host key!')
        elif keys[hostname][key.get_name()] != key:
            print('*** WARNING: Host key has changed!!!')
            sys.exit(1)
        else:
            print('*** Host key OK.')

        if username == '':
            default_username = getpass.getuser()
            username = raw_input('Username [{}]: '.format(default_username))
            if len(username) == 0:
                username = default_username

        # starting authentication.
        agent_auth(t, username)
        if not t.is_authenticated():
            manual_auth(t, username, hostname)
        if not t.is_authenticated():
            print('*** Authentication failed. :(')
            t.close()
            sys.exit(1)

        # starting session.
        chan = t.open_session()
        chan.get_pty()
        chan.invoke_shell()
        print('*** Here we go!\n')

        check_terminal(chan, username)
        # interactive.interactive_shell(chan)
        chan.close()
        t.close()

    except Exception as e:
        print('*** Caught exception ' + str(e.__class__) + ': ' + str(e))
        traceback.print_exc()
        sys.exit(1)


def check_terminal(chan, username):
    if has_termios:
        posix_shell(chan, username)


def posix_shell(chan, username):
    import select

    oldtty = termios.tcgetattr(sys.stdin)
    f = open('{}.log'.format(username), 'a+')
    try:
        tty.setraw(sys.stdin.fileno())
        # tty.setcbreak(sys.stdin.fileno())
        chan.settimeout(0.0)

        order = []
        flag = False
        while True:
            r, w, e = select.select([chan, sys.stdin], [], [])
            if chan in r:
                try:
                    x = u(chan.recv(1024))
                    if len(x) == 0:
                        sys.stdout.write('\r\n*** EOF\r\n')
                        break
                    if flag:
                        if x.startswith('\r\n'):
                            pass
                        else:
                            order.append(x)
                        flag = False
                    sys.stdout.write(x)
                    sys.stdout.flush()
                except socket.timeout:
                    pass
            if sys.stdin in r:
                x = u(sys.stdin.read(1))
                if x == '\t' or x == '\b':
                    flag = True
                else:
                    order.append(x)

                    if x == '\r':
                        # strip = [u'\t', u'\x7f', u'\x07', u'vim']
                        c_time = time.strftime('%Y-%m-%d %H:%M:%S')
                        # print(cmd)
                        cmd = ''.join(order).replace(u'\r', u'\n')
                        # cmd = ''.join(order).replace(u'\r', u'\n').strip(u'\t').strip(u'\x7f').strip(u'\x07')
                        if u'vi' in cmd[:3] and cmd[-1] == '\r':
                            f.write('{}     {}'.format(c_time, cmd))
                            order = []

                        else:
                            f.write('{}     {}'.format(c_time, cmd))
                            # f.flush()
                            order = []

                if len(x) == 0:
                    break
                chan.send(x)
    finally:
        termios.tcsetattr(sys.stdin, termios.TCSADRAIN, oldtty)
        f.close()

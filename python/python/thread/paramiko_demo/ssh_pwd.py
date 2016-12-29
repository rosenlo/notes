#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: ssh_pwd.py
@time: 10/12/16 10:17 AM
"""
import paramiko

ssh = paramiko.SSHClient()
ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
ssh.connect('192.168.186.130', 22, 'root', 'redhat')
stdin, stdout, stderr = ssh.exec_command('df -h')
print(stdout.read())

ssh.close()

if __name__ == '__main__':
    pass

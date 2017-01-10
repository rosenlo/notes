#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: ssh_key.py
@time: 10/12/16 11:01 AM
"""
import paramiko

private_key_path = '/Users/Rosen/.ssh/id_rsa'
key = paramiko.RSAKey.from_private_key_file(private_key_path)

ssh = paramiko.SSHClient()
ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
ssh.connect('192.168.186.130', 22, 'root', pkey=key)
stdin, stdout, stderr = ssh.exec_command('df -h')
print(stdout.read())
stdin, stdout, stderr = ssh.exec_command('cd /tmp && pwd')
print(stdout.read())
stdin, stdout, stderr = ssh.exec_command('pwd')
print(stdout.read())

ssh.close()

if __name__ == '__main__':
	pass

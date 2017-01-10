#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: db.py
@time: 10/12/16 3:35 PM
"""
from audit.utils.sqlhelper import MySQLHelper


class FortressDB(object):
    def __init__(self):
        self.__helper = MySQLHelper()

    def insert_users(self, username, password, user_group):
        sql = 'INSERT INTO users (username, password, user_group) VALUE (%s, %s, %s)'
        params = (username, password, user_group)
        self.__helper.insert_and_update(sql, params)

    def insert_group(self, host_ip, hostname, host_group, user_group):
        sql = 'INSERT INTO users (host_ip, hostname, host_group, user_group) VALUE (%s, %s, %s, %s)'
        params = (host_ip, hostname, host_group, user_group)
        self.__helper.insert_and_update(sql, params)

    def update(self, username, after_login=None):
        sql = 'UPDATE users SET after_login=%s WHERE username=%s'
        params = (after_login, username)
        self.__helper.insert_and_update(sql, params)

    def user_auth(self, username, password):
        sql = 'SELECT * from users WHERE username=%s AND password=%s'
        params = (username, password)
        return self.__helper.get_data(sql, params, size='all')

    def display_hosts(self, username):
        sql = "SELECT hostname,host_ip FROM host_list c, group_relation b, user2group a " \
              "WHERE a.username=%s " \
              "AND a.user_group=b.user_group " \
              "AND b.host_group=c.host_group"
        params = (username,)
        return self.__helper.get_data(sql, params, size='all')

    def get_user(self, host_ip):
        sql = "SELECT hostuser, hostpwd FROM host_users a, host_list b WHERE b.host_ip=%s AND b.hostname=a.hostname"
        params = (host_ip,)
        return self.__helper.get_data(sql, params, size='all')


if __name__ == '__main__':
    pass

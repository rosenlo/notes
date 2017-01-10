#!/usr/bin/env python
# -*- coding=utf-8 -*-

import MySQLdb
import sys

reload(sys)
sys.setdefaultencoding('utf-8')

INFO = ('localhost', 'root', 'redhat', 'simplefortresses')
username = 'lzh'
password = '123'
user_group = 'admin'

conn = MySQLdb.connect(*INFO, charset='utf8')
cur = conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)

sql1 = 'INSERT INTO users (username, password, user_group) VALUE (%s, %s, %s)'
sql2 = 'INSERT INTO users (host_ip, hostname, host_group, user_group) VALUE (%s, %s, %s, %s)'
sql3 = 'UPDATE users SET after_login=%s WHERE username=%s'
sql4 = 'SELECT * from users WHERE username=%s and password=%s'
sql5 = "SELECT host_ip FROM host_list c, group_relation b, users a " \
	   "WHERE a.username=%s " \
	   "AND a.user_group=b.user_group " \
	   "AND b.host_group=c.host_group "

params1 = (username, password, user_group)
params4 = (username, password)
params3 = (username, 'NULL')
params5 = (username,)
cur.execute(sql3, params3)
conn.commit()
print cur.fetchall()

cur.close()
conn.close()

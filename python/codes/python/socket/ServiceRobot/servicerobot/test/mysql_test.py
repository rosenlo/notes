#!/usr/bin/env python
# -*- coding=utf-8 -*-

import MySQLdb
import sys
reload(sys)
sys.setdefaultencoding('utf-8')

INFO = ('localhost', 'service_robot', '123456', 'service_robot')

conn = MySQLdb.connect(*INFO, charset='utf8')
cur = conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)

sql = 'select * from chat_records'
#sql = "insert into chat_records (username, sender, recevier, text) values ('a', 'b', 'c', 'd')"

cur.execute(sql)
conn.commit()
print cur.fetchall()

cur.close()
conn.close()

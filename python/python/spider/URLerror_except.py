#!/usr/bin/env python
# -*- coding=utf-8 -*-

# Author: Rosen
# Mail: rosenluov@gmail.com
# Created Time: Sat Dec 17 18:03:13 2016

import urllib2

#req = urllib2.Request('http://www.ffff.com')

req = urllib2.Request('http://bbs.csdn.net/callmewhy') 
try:
    print urllib2.urlopen(req)
except urllib2.URLError as e:
    print e.code
    print e.read()

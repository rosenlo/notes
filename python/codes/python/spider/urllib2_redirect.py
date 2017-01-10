#!/usr/bin/env python
# -*- coding=utf-8 -*-

# Author: Rosen
# Mail: rosenluov@gmail.com
# Created Time: Sat Dec 17 18:23:56 2016

import urllib2

old_url = 'http://cm.g.doubleclick.net/pixel?google_nid=baidu&google_cm'
req = urllib2.Request(old_url)

res = urllib2.urlopen(req)

print(u'Old url : ' + old_url)
print(u'Real url : ' + res.geturl())


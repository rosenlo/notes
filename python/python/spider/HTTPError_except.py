#!/usr/bin/env python
# -*- coding=utf-8 -*-

# Author: Rosen
# Mail: rosenluov@gmail.com
# Created Time: Sat Dec 17 18:09:37 2016

import urllib2

req = urllib2.Request('http://bbs.csdn.net/callmewhy')

try:
    res = urllib2.urlopen(req)

except urllib2.HTTPError as e:
    print(u'The server couldn\'t fulfill the request.')
    print(u'Error code: ', e.code)

except urllib2.URLError as e:
    print(u'We failed to reach a server.')
    print(u'Reason: ', e.reason)

else:
    print(u'No exception was raised.')

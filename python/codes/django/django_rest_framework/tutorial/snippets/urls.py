#!/usr/bin/env python
# encoding: utf-8

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: urls.py
Created Time: 2/8/17 21:06
"""

from django.conf.urls import url, include
from snippets.views import *
from rest_framework.urlpatterns import format_suffix_patterns

urlpatterns = [
    url(r'^snippets/$', SnippetList.as_view()),
    url(r'^snippets/(?P<pk>[0-9]+)/$', SnippetDetail.as_view()),
    url(r'^users/$', UserList.as_view()),
    url(r'^users/(?P<pk>[0-9]+)/$', UserDetail.as_view()),
]

urlpatterns = format_suffix_patterns(urlpatterns)

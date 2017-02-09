#!/usr/bin/env python
# encoding: utf-8

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: urls.py
Created Time: 2/8/17 21:06
"""

from django.conf.urls import url
from snippets import views

urlpatterns = [
    url(r'^snippets/$', views.snippet_list),
    url(r'^snippets/(?P<pk>[0-9]+)/$', views.snippet_detail),
]

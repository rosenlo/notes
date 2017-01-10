#!/usr/bin/env python
# encoding: utf-8

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: urls.py
Created Time: 12/18/16 22:53
"""
from django.conf.urls import url
from host_manage import views

urlpatterns = [
    url(r'^register/', views.register),
    url(r'^login/', views.login),
    url(r'^create_host/', views.create_host),
    url(r'^hosts/', views.show_hosts),
]

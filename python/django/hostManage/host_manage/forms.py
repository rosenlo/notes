#!/usr/bin/env python
# encoding: utf-8

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: forms.py
Created Time: 12/18/16 22:56
"""

from django import forms


class Register(forms.Form):
    username = forms.CharField(required=True)
    password = forms.CharField()
    email = forms.EmailField()


class Login(forms.Form):
    username = forms.CharField(required=True)
    password = forms.CharField(required=True)

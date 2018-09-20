#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: first_app.py
Created Time: Thu Apr 26 17:02:12 2018
"""

from flask import Flask

print(__name__)
app = Flask(__name__)


@app.route('/')
def hello_world():
    return "hello, world!"


@app.route('/params/<name>')
def hello_params(name):
    return "hello, {}!".format(name)


@app.route('/projects/')
def projects():
    return 'The project page'


@app.route('/about')
def about():
    return 'The about page'

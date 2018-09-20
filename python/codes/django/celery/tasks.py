#!/usr/bin/env python
# encoding: utf-8

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: tasks.py
Created Time: 3/22/17 11:52
"""

from celery import Celery

app = Celery('tasks', broker='redis://localhost:6379/0')


@app.task
def add(x, y):
    return x + y



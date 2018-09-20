#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: setup.py
Created Time: Wed May  2 14:23:30 2018
"""


try:
    with open('requirements.txt', 'r') as f:
        requirements = [
            item for item in f.readlines()
            if not item.strip().startswith('#')
        ]
except IOError:
    requirements = []

from setuptools import find_packages, setup
from flaskr import __version__

setup(
    name='flaskr',
    version=__version__,
    packages=find_packages(),
    include_package_data=True,
    zip_safe=False,
    install_requires=requirements,
)

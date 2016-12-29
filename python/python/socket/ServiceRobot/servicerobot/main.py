#!/usr/bin/env python
# encoding: utf-8

"""
@author: Rosen
@file: main.py
@time: 9/20/16 10:51 AM
"""
from utils.client import Client


def CreateConnect():
    Client('127.0.0.1', 9999)


def main():
    CreateConnect()


if __name__ == '__main__':
    CreateConnect()

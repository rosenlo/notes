#!/usr/bin/env python
# encoding: utf-8

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: serializers.py
Created Time: 2/8/17 12:08
"""

from django.contrib.auth.models import User, Group
from rest_framework import serializers


class UserSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = User
        fields = ('url', 'username', 'email', 'groups')


class GroupSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Group
        fields = ('url', 'name')

#!/usr/bin/env python
# encoding: utf-8

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: permissions.py
Created Time: 3/7/17 11:17
"""

from rest_framework import permissions


class IsOwnerOrReadOnly(permissions.BasePermission):
    """
    Create permission to only allow owners of an object to edit it.
    """

    def has_object_permission(self, request, view, obj):
        # Read permissions are allowed to any request.
        # so we'll always allow GET, HEAD or Options requests.
        if request.method in permissions.SAFE_METHODS:
            return True

        # write permissions are only allowed to the owner of the snippet.
        return obj.owner == request.user

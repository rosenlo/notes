from __future__ import unicode_literals

from django.db import models


# Create your models here.


class UserGroup(models.Model):
    name = models.CharField(max_length=20)

    def __unicode__(self):
        return self.name


class UserInfo(models.Model):
    name = models.CharField(max_length=20)
    password = models.CharField(max_length=50)
    createTime = models.DateTimeField(auto_now_add=True)
    updateTime = models.DateTimeField(auto_now=True)
    userGroup = models.ManyToManyField(UserGroup)

    def __unicode__(self):
        return self.name


class Hosts(models.Model):
    hostname = models.CharField(max_length=50)
    ip = models.GenericIPAddressField()
    userGroup = models.ForeignKey(UserGroup)

    def __unicode__(self):
        return self.hostname

#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: test_db.py
Created Time: Wed May  2 15:43:17 2018
"""


import sqlite3
import pytest
from flaskr.db import get_db


def test_get_close_db(app):
    with app.app_context():
        db = get_db()
        assert db is get_db()

    with pytest.raises(sqlite3.ProgrammingError) as e:
        db.execute('SELECT 1')

    assert 'closed ' in str(e)


def test_init_db_command(runner, monkeypatch):
    class Record(object):
        called = False

    def fake_init_db():
        Record.called = True

    monkeypatch.setattr('flaskr.db.init_db', fake_init_db)
    result = runner.invoke(args=['init-db'])
    assert 'Initalized' in result.output
    assert Record.called

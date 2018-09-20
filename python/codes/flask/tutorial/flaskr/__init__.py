#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
Author: Rosen
Mail: rosenluov@gmail.com
File: __init__.py
Created Time: Sat Apr 28 14:53:41 2018
"""

# Semantic version: Major.Minor.Patch
__version__ = '1.0.0'


import os
from flask import Flask


def create_app(test_config=None):
    # create and configure the app
    app = Flask(__name__, instance_relative_config=True)
    app.config.from_mapping(
        SECRET_KEY='dev',
        DATABASE=os.path.join(app.instance_path, 'flaskr.sqlite')
    )

    if test_config is None:
        # Load the instance config, if it exists, when not testing
        app.config.from_pyfile('config.py', silent=True)
    else:
        # Load the test config if passed in
        app.config.from_mapping(test_config)

    # ensure the instance folder exists
    try:
        os.makedirs(app.instance_path)
    except OSError:
        pass

    # a simle page that say hello
    @app.route('/hello')
    def hello():
        return 'Hello, There!'

    from . import db
    db.init_app(app)

    from . import auth
    app.register_blueprint(auth.bp)

    from . import blog
    app.register_blueprint(blog.bp)
    app.add_url_rule('/', endpoint='index')

    return app

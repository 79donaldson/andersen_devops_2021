#!/usr/bin/env python
# -*- coding: utf-8 -*-
from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, this is my first application! Первое минималистическое приложение в Docker.'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=False)
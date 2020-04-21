#!/usr/bin/python3

import json
from typing import Dict, TextIO

from flask import request
from flask import Flask



LOG_FILE: TextIO = open('morty_ci.log', 'w')
app: Flask = Flask('morty-ci')



class Handlers:
    @staticmethod
    def push(github_event: str, info: Dict):
        full_name: str = info['repository']['full_name']
        print(f'{github_event.upper()}: {full_name}')

    @staticmethod
    def pull_request(github_event: str, info: Dict):
        action: str = info['action']
        full_name: str = info['repository']['full_name']
        print(f"{github_event.upper()} ({action}): {full_name}")

MUX: Dict = {
    'push': Handlers.push,
    'pull_request': Handlers.pull_request,
}



@app.route('/github', methods=['POST'])
def github() -> str:
    if request.headers['Content-Type'] != 'application/json':
        return '415 - unsupported media type'

    github_event: str = request.headers['X-GitHub-Event']
    info: Dict = request.json
    MUX[github_event](github_event, info)
    return '200 - OK'



if __name__ == '__main__':
    app.run(debug=True)


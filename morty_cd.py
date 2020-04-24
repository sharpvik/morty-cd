#!/usr/bin/python3

# Python Standard Library Imports
import json
from sys import argv as args
import subprocess as sh
from typing import Dict, TextIO

# External Imports
from flask import Flask, request
import waitress # type: ignore

# Local Imports
import config



"""-----------------------------------------------------------------------------
SCRIPT is a string representing a script that you want to run on Push and Pull
Request. It must reside withing a directory that is available to the $PATH
evironment variable.
-----------------------------------------------------------------------------"""
SCRIPT: str = args[1]

LOG_FILE: TextIO = open(args[2], 'w')

app: Flask = Flask('morty-cd')



"""-----------------------------------------------------------------------------
Handlers class is a tiny namespace where static handler functions are defined.
These handlers are used to handle specific GitHub webhook events.
More about GitHub webhook events here:
    https://developer.github.com/v3/activity/events/types
-----------------------------------------------------------------------------"""
class Handlers:
    @staticmethod
    def ping(event: str, info: Dict):
        pass

    @staticmethod
    def push(event: str, info: Dict):
        sh.run(SCRIPT, stdout=LOG_FILE)

    @staticmethod
    def pull_request(event: str, info: Dict):
        action: str = info['action']
        merged: bool = info['pull_request']['merged']

        # Only run SCRIPT when Pull Request was closed
        # with a successful merge.
        if action == 'closed' and merged:
            sh.run(SCRIPT, stdout=LOG_FILE)



"""-----------------------------------------------------------------------------
MUX is an alternative to a switch case. It maps GitHub webhook event strings to
a handler function that's used to handle that event.
-----------------------------------------------------------------------------"""
MUX: Dict = {
    'ping': Handlers.ping,
    'push': Handlers.push,
    'pull_request': Handlers.pull_request,
}



"""-----------------------------------------------------------------------------
Use this route to test your server before doing anything crazy. Just see if it's
working or not.
-----------------------------------------------------------------------------"""
@app.route('/ping', methods=['GET'])
def ping() -> str:
    return '200 - OK'



"""-----------------------------------------------------------------------------
When setting up a GitHub webhook for this app, don't forget to append '/github'
to the end of your server's link.

For example, if this script runs under domain name 'http://mycoolserver.org',
give the following link to GitHub:
    http://mycoolserver.org/github
-----------------------------------------------------------------------------"""
@app.route('/github', methods=['POST'])
def github() -> str:
    # On the same note, please make sure that the webhook sends you data in JSON
    # format only. Otherwise, feel free to rewrite this function your way.
    if request.headers['Content-Type'] != 'application/json':
        return '415 - unsupported media type'

    event: str = request.headers['X-GitHub-Event']
    info: Dict = request.json
    MUX[event](event, info)
    return '200 - OK'



if __name__ == '__main__':
    if config.STATUS == 'dev':
        app.run(debug=True)
    else:
        waitress.serve(app, port=config.PORT)


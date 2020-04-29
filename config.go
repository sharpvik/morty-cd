package main

import "os"

// port is a string that represents the port at which server will be listening
// for incoming requests.
const port = ":5050"

// logWriter implements the io.Writer interface and is used by the logr to
// output the data it logs.
var logWriter = os.Stdout
// var logWriter, _ = os.Create("/home/viktor/Public/Lisn/logs/morty-cd-lisn-web-app.log")

// logPrefix is a string that will be inserted before every log message.
const logPrefix = ""

# Morty CD

Simple Go tool for Continuous Integration and Continuous Deployment utilizing
the power of GitHub webhooks.



## Prerequisites

- Install *Go* programming language.



## User Guide

### Setup On Machine

#### Get Source Code

```bash
go get github.com/sharpvik/morty-cd
```

#### Add Config File

This is how my `config.go` looks (you can just `Copy + Paste`):

```go
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
```

#### Installation

To install this globally, you can do two things:

1. Add your `$(go env GOPATH)/bin` folder to `$PATH` and run `go install`
2. Run `go build` and then `mv morty-cd /usr/local/bin/` or to any other folder
registered within `$PATH`


### GitHub Webhooks Setup

Locate the GitHub repo you'd like to bind your `script` to. From there go
**Settings > Webhooks > Add webhook**. For **Payload URL** use your server's IP
address and port like so: `http://155.13.12.5:5050/github`. You must include the
`/github` at the end!

Set **Content type** to `application/json` and select the events you want to
trigger webhook notifications. Confirm.

> If you don't see some of the events you want included in the `handlers.go`
> file, it means you'll have to include them yourself and then modify the switch
> statement in the main handler function called `github`.
> If you do create some additional handlers, please share them with us!


### Run It

When you've properly installed the app and setup webhooks, go to the project you
want to monitor and run `morty-cd ./deploy.sh`. This step *assumes* that some
deployment script called `deploy.sh` is present within that folder and has
executive privileges.

> To give executive privileges to any script, specify its interpreter on the
> very first line by adding `#!/usr/bin/interpreter` add then run
> `chmod +x deploy.sh`. For example, if your script is a `bash` script, put
> `#!/usr/bin/bash` as your shebang.

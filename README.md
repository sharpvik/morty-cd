# Morty CI

Simple Python Flask tool for Continuous Integration and Continuous Deployment
utilizing the power of GitHub webhooks.



## Prerequisites

- Install *Go* programming language.



## User Guide

### Setup On Machine

```bash
go get github.com/sharpvik/morty-cd

cd morty-cd

echo 'echo Executing the dummy build script...' > ./script

go build
```


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

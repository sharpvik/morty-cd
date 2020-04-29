package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

//! 
//! Go, like Flask, only actually responds to the client on function return.
//! However, Go's concurrency is much easier to understand. Though, there are a
//! few things to watch out for.
//!
//! GitHub's webhook events carry a lot of JSON data, so the pullRequestEvent
//! struct comes out extremely large. It is extremely important that it is passed
//! by &reference for efficiency reasons. Especially, knowing that it is never
//! mutated anyways.
//! 

var logr *log.Logger

// script is a string that represents a script that you want to run on GitHub
// webhook event.
var script string

func main() {
	logr = log.New(logWriter, logPrefix, log.Ltime)

	if len(os.Args) < 2 {
		logr.Fatal("Specify the script name!")
	}

	script = os.Args[1]

	http.HandleFunc("/ping", onPing)   // Local testing.
	http.HandleFunc("/github", github) // URL utilized by GitHub.

	logr.Printf("Serving at port %s", port)
	http.ListenAndServe(port, nil)
}

// runScript is a function we use to safely execute the script.
func runScript() {
	logr.Printf("Executing %s...", script)

	// This part is responsible for running the script concurrently. Otherwise,
	// the client will time out and GitHub will mark payload as undelivered.
	go func() {
		bts, err := exec.Command(script).Output()

		if err != nil {
			logr.Print(err)
			return
		}

		logr.Print(string(bts))
	}()
}

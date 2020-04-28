package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

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

	bts, err := exec.Command(script).Output()

	if err != nil {
		logr.Print(err)
		return
	}

	logr.Println(bts)
}

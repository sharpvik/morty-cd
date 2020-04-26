package main

import (
    "net/http"
    "io"
    "os/exec"
    "os"
    "encoding/json"
    "log"
)

var logr *log.Logger
var script string

func main() {
    logr = log.New(logWriter, logPrefix, log.Ltime)
    
    if (len(os.Args) < 2) {
        logr.Fatal("Specify the script name!")
    }

    script = os.Args[1]

    http.HandleFunc("/ping", ping)
    http.HandleFunc("/github", github)

    logr.Printf("Serving at port %s", port)
    logr.Fatal(http.ListenAndServe(port, nil))
}

// Dummy ping handler function to check whether everything's working.
func ping(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "200 - OK")
    logr.Printf("Received ping request")
}

// Main handler function.
func github(w http.ResponseWriter, r *http.Request) {
    // Must reject all requests except POST!
    if r.Method != "POST" {
        w.WriteHeader(http.StatusBadRequest)
        io.WriteString(w, "400 - bad request")
        return
    }

    io.WriteString(w, "200 - OK")

    event := r.Header.Get("X-GitHub-Event")

    switch event {
    case "ping":
        ping(w, r)

    case "push":
        onPush()

    case "pull_request":
        var info PullRequestEvent
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&info)

        if err != nil {
            logr.Print("Cannot decode JSON")
            return
        }

        onPullRequest(info)
    }   
}

func onPush() {
    logr.Print("Push detected")
    logr.Printf("Executing %s...", script)
    exec.Command(script)
}

func onPullRequest(info PullRequestEvent) {
    logr.Print("Pull Request detected")

    action := info.Action
    merged := info.PullRequest.Merged

    // Only run script when pull request has been closed and merged.
    if action == "closed" && merged {
        logr.Print("Pull request has been merged")
        logr.Printf("Executing %s...", script)
        exec.Command(script)
        return
    }

    logr.Print("Pull request failed the 'closed with merge' test")
}

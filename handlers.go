package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// All handler functions are defined in this file.
// These handlers are used to handle specific GitHub webhook events.
// More about GitHub webhook events here:
//     https://developer.github.com/v3/activity/events/types

// Dummy ping handler function to check whether everything's working.
func onPing(w http.ResponseWriter, r *http.Request) {
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

	// Responding first, so that client doesn't time out.
	io.WriteString(w, "200 - OK")

	event := r.Header.Get("X-GitHub-Event")

	switch event {
	case "ping":
		onPing(w, r)

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

// The following handlers respond to specific event types.
func onPush() {
	logr.Print("Push detected")
	runScript()
}

func onPullRequest(info PullRequestEvent) {
	logr.Print("Pull Request detected")

	action := info.Action
	merged := info.PullRequest.Merged

	// Only run script when pull request has been closed and merged.
	if action == "closed" && merged {
		logr.Print("Pull request has been merged")
		runScript()
		return
	}

	logr.Print("Pull request failed the 'closed with merge' test")
}

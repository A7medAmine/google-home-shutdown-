package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

const secret = "CHANGE_ME_TO_A_LONG_RANDOM_STRING"

func shutdown(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("key") != secret {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	exec.Command("shutdown", "/s", "/f", "/t", "0").Run()
	fmt.Fprintln(w, "Shutting down...")
}

func main() {
	http.HandleFunc("/shutdown", shutdown)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
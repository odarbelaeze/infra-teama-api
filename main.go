package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
	Path    string `json:"path"`
	Host    string `json:"host"`
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "set the `MESSAGE` env var to customize"
	}
	vr := Response{Message: message, Path: r.URL.Path, Host: host}
	body, _ := json.Marshal(vr)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", body)
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

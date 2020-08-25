package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()
	nodeName := os.Getenv("NODE_NAME")
	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		listenPort = "80"
	}

	tw := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fwd := r.Header.Get("X-Forwarded-For")
		fmt.Fprintf(w, "%s[ff: %s] --> %s[node: %s]\n", r.RemoteAddr, fwd, hostname, nodeName)
		log.Printf("%s[ff: %s] --> %s[node: %s]\n", r.RemoteAddr, fwd, hostname, nodeName)

	}
	http.HandleFunc("/", tw)

	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}

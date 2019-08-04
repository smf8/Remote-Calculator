package main

import (
	"Remote-Calculator/pkg/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/solve", server.HandleProblem)
	log.Fatal(http.ListenAndServe(":1234", nil))
}

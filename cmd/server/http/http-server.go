package main

import (
	"log"
	"net/http"

	"github.com/smf8/Remote-Calculator/pkg/server"
)

func main() {
	http.HandleFunc("/solve", server.HandleProblem)
	log.Fatal(http.ListenAndServe(":1234", nil))
}

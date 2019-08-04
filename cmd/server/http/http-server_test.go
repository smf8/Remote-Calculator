package main

import (
	"Remote-Calculator/pkg/model"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestServerSetup(t *testing.T) {
	go main()
	p := model.NewProblem("4+1.3")
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)
	res, err := http.Post("http://localhost:1234/solve", "application/json; charset=utf-8", b)
	if err != nil {
		t.Error("Failed in sending POST request")
	}
	json.NewDecoder(res.Body).Decode(&p)
	if p.Result != 5.3 {
		t.Error("Failed in computation")
	}
}

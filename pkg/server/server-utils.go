package server

import (
	"Remote-Calculator/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
)

//HandleProblem function to handle problems and calculate them
func HandleProblem(wr http.ResponseWriter, r *http.Request) {
	p := new(model.Problem)
	if r.Body == nil {
		http.Error(wr, "Request Body is empty, send something", 400)
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err, " On json parsing")
	}
	err = p.Compute()
	if err != nil {
		fmt.Println(err, " On computation")
	}
	json.NewEncoder(wr).Encode(p)
	fmt.Println(p.Result, " is the result")
}

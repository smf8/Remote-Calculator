package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/smf8/Remote-Calculator/pkg/model"
)

func BenchmarkCalculation(b *testing.B) {
	for i := 1; i < b.N; i++ {
		p := model.Problem{
			FirstOperand:  float64(i),
			SecondOperand: float64((b.N - i) / 3),
		}
		if p.SecondOperand == 0 {
			p.SecondOperand = 1
		}
		switch i % 5 {
		case 0:
			p.Operator = "+"
		case 1:
			p.Operator = "-"
		case 2:
			p.Operator = "*"
		case 3:
			p.Operator = "/"
		case 4:
			p.Operator = "%"
		}
		p.Compute()
	}
}
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

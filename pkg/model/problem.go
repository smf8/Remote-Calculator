package model

import (
	"errors"
	"regexp"
	"strconv"
)

//Problem is a model of a simple mathematical problem
type Problem struct {
	FirstOperand  float64 `json:"first_operand"`
	SecondOperand float64 `json:"second_operand"`
	Operator      string  `json:"operator"`
	Result        float64 `json:"result"`
}

func (p *Problem) Compute() error {
	switch p.Operator {
	case "+":
		p.Result = p.FirstOperand + p.SecondOperand
	case "-":
		p.Result = p.FirstOperand - p.SecondOperand
	case "*":
		p.Result = p.FirstOperand * p.SecondOperand
	case "/":
		p.Result = p.FirstOperand / p.SecondOperand
	case "%":
		p.Result = float64(int(p.FirstOperand) % int(p.SecondOperand))
	default:
		p.Result = 0.0
		return errors.New("Inavlid operation")
	}
	return nil
}

//NewProblem creates a new problem from text
func NewProblem(s string) *Problem {
	p := new(Problem)

	// don't know what this is so don't ask just a regex to find * + - / % operator
	regx := regexp.MustCompile("[%-+*/\\/]")
	p.Operator = string(regx.Find([]byte(s)))
	spl := regx.Split(s, -1)
	p.FirstOperand, _ = strconv.ParseFloat(spl[0], 64)
	p.SecondOperand, _ = strconv.ParseFloat(spl[1], 64)
	return p
}

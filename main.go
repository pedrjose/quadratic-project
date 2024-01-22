package main

import (
	"fmt"

	"github.com/pedrjose/quadratic-project/quadratic"
)

func main() {
	a := 1.0
	b := 6.0
	c := 9.0

	delta := quadratic.Discriminant(a, b, c)
	bhaskara := quadratic.Bhaskara(delta, a, b)

	fmt.Println(bhaskara)
}


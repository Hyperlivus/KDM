package main

import (
	"os"
	//"kdm/set"
	"fmt"
	pars "kdm/parser"
	"kdm/solver"
)

func main() {
	path := "target/eq.set"
	code, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("File not found")
		return
	}

	parser := pars.NewParser(code)
	parser.Parse()

	for _, eq := range parser.Equations() {
		solutions, field := solver.SolveEq(eq)
		for _, sol := range solutions {
			fmt.Println(sol.ToString(field))
		}
	}

}

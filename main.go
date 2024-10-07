package main

import (
	"os"
	//"kdm/set"
	"fmt"
	pars "kdm/parser"
	"kdm/solver"
)

func LoadResToFile(path string, equations []pars.Equation) error {
	res := ""
	for _, eq := range equations {
		solutions, field := solver.SolveEq(eq)

		for _, sol := range solutions {
			res += sol.ToString(field)
			res += "\n"
		}
	}

	err := os.WriteFile(path, []byte(res), 0644)
	return err
}

func main() {
	path := "target/eq.set"
	code, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("File not found")
		return
	}

	parser := pars.NewParser(code)
	parser.Parse()
	equations := parser.Equations()

	err = LoadResToFile("target/eq.res", equations)
	if err != nil {
		fmt.Println(err.Error())
	}

}

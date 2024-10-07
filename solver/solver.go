package solver

import (
	"fmt"
	pars "kdm/parser"
	"kdm/set"

	slices "golang.org/x/exp/slices"
)

func SolveEq(eq pars.Equation) ([]EqSolution, *set.Field) {
	leftIdents := FindIdentsName(eq.Left)
	rightIdents := FindIdentsName(eq.Right)

	var idents []string = append(leftIdents, rightIdents...)
	idents = slices.Compact(idents)
	field := set.NewField(idents)

	for name, set := range field.Sets() {
		fmt.Printf("%s : %5b \n", name, set)
	}

	expr := pars.NewBinary(pars.SYM_SUB, eq.Left, eq.Right)
	res := Eval(field, expr)

	if res.Tag() == set.VARIABLE {
		variable := res.(set.Variable)
		sets, values := GetAllVariants(field, variable)
		fmt.Printf("%9b \n", variable)

		solutions := make([]EqSolution, 0)
		for i := 0; i < len(sets); i++ {
			s := sets[i]
			value := values[i]
			fmt.Printf("%5b \n", value)
			fmt.Println(SetToString(field, value))
			if value == field.None() {

				solution := NewPrmSolution(s)
				solutions = append(solutions, solution)
			} else {
				rule := NewRule([][2]set.Set{
					{value, field.None()},
				})
				solution := NewSolution(s, rule)
				solutions = append(solutions, solution)
			}
		}

		return solutions, field
	} else if res.Tag() == set.SET {
		set := res.(set.Set)
		if set == field.None() {
			return make([]EqSolution, 0), field
		} else {
			return make([]EqSolution, 0), field
		}

	}
	panic("wrong tag")
}

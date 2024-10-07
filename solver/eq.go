package solver

import (
	"fmt"
	"kdm/set"
)

type Rule struct {
	Equivalents [][2]set.Set
}

func (rule Rule) ToString(field *set.Field) string {
	res := "["
	for i, eq := range rule.Equivalents {
		lStr := SetToString(field, eq[0])
		rStr := SetToString(field, eq[1])
		res += fmt.Sprint(lStr, " = ", rStr)
		if i != len(rule.Equivalents)-1 {
			res += ";"
		}
	}
	res += "]"

	return res
}
func NewRule(equivalents [][2]set.Set) Rule {
	return Rule{
		Equivalents: equivalents,
	}
}

var NONE_RULE = Rule{
	Equivalents: make([][2]set.Set, 0),
}

type EqSolution struct {
	ans  set.Set
	rule Rule
}

func (sol EqSolution) ToString(field *set.Field) string {
	return fmt.Sprint("x = ", SetToString(field, sol.ans), ", if ", sol.rule.ToString(field))
}
func NewSolution(ans set.Set, rule Rule) EqSolution {
	return EqSolution{
		ans,
		rule,
	}
}
func NewPrmSolution(ans set.Set) EqSolution {
	return EqSolution{
		ans:  ans,
		rule: NONE_RULE,
	}
}

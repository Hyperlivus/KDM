package parser

type Equation struct {
	Left Expression
	Right Expression
}

func NewEquation(left Expression, right Expression) Equation {
	return Equation{
		left, 
		right,
	}
}
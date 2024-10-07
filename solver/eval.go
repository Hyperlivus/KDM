package solver

import (
	pars "kdm/parser"
	set "kdm/set"

	slices "golang.org/x/exp/slices"
)

func Eval(field *set.Field, expr pars.Expression) set.Expression {
	univ := field.Univ()
	if expr.Tag() == pars.UNARIAN {
		unarian := expr.(pars.Unarian)

		res := Eval(field, unarian.Value)

		if unarian.Op == pars.BR {
			return res
		} else if unarian.Op == pars.NOT {
			return set.UseOperand(univ, res, func(a uint, b uint) uint {
				return a ^ b
			})
		}
	} else if expr.Tag() == pars.BINARY {
		binary := expr.(pars.Binary)
		lValue := Eval(field, binary.Lhs)
		rValue := Eval(field, binary.Rhs)

		var op set.Operand

		if binary.Op == pars.UNIT {
			op = func(a uint, b uint) uint {
				return a | b
			}
		} else if binary.Op == pars.INTER {
			op = func(a uint, b uint) uint {
				return a & b
			}
		} else if binary.Op == pars.SYM_SUB {
			op = func(a uint, b uint) uint {
				return a ^ b
			}
		} else if binary.Op == pars.MUL {
			op = func(a uint, b uint) uint {
				notA := univ.Base() ^ a
				notB := univ.Base() ^ b

				return notA | notB
			}
		} else if binary.Op == pars.SUB {
			op = func(a uint, b uint) uint {
				notB := univ.Base() ^ b
				return a & notB
			}
		}
		return set.UseOperand(lValue, rValue, op)
	} else if expr.Tag() == pars.IDENT {
		ident := expr.(pars.Ident)
		name := ident.Name
		if slices.Contains(VariableNames, name) {
			return set.NewVariable(univ, field.PartsCount())
		} else {
			set := field.Sets()[name]
			return set
		}
	}
	panic("wrong tag")
}

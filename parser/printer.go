package parser

import (
	"fmt"
)

var BinaryStr = map[BinOperand]string{
	UNIT:    "+",
	INTER:   "^",
	MUL:     "*",
	SUB:     "/",
	SYM_SUB: "-",
}

func SprintExpr(expr Expression) string {
	if expr.Tag() == IDENT {
		ident := expr.(Ident)
		return ident.Name
	} else if expr.Tag() == BINARY {

		binary := expr.(Binary)
		lStr := SprintExpr(binary.Lhs)
		rStr := SprintExpr(binary.Rhs)
		return fmt.Sprint(lStr, " ", BinaryStr[binary.Op], " ", rStr)
	} else if expr.Tag() == UNARIAN {
		unarian := expr.(Unarian)

		if unarian.Op == NOT {
			return fmt.Sprint("!(", SprintExpr(unarian.Value), ")")
		} else if unarian.Op == BR {
			return fmt.Sprint("(", SprintExpr(unarian.Value), ")")
		}
	}
	panic("undefined tag")
}

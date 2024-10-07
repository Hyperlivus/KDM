package parser

import (
	_ "fmt"
	"unicode"
)

const LINE_TERMINATOR = ';'

var binOpChar = map[byte]BinOperand{
	'+': UNIT,
	'^': INTER,
	'*': MUL,
	'/': SUB,
	'-': SYM_SUB,
}

type Parser struct {
	cursor    *Cursor
	exprs     []Expression
	equations []Equation
}

func (parser *Parser) Exprs() []Expression {
	return parser.exprs
}
func (parser *Parser) Equations() []Equation {
	return parser.equations
}
func (parser *Parser) Parse() {
	cursor := parser.cursor

	for !cursor.IsEnd() {
		cursor.EatSpaces()
		expr := parser.ParseExpr()
		curr := cursor.Current()

		if curr == LINE_TERMINATOR {
			cursor.RequireNext(LINE_TERMINATOR)
			parser.exprs = append(parser.exprs, expr)
		} else if curr == '=' {
			cursor.RequireNext('=')
			next := parser.ParseExpr()
			equation := NewEquation(expr, next)
			parser.equations = append(parser.equations, equation)
			cursor.RequireNext(';')
		} else {
			panic("something went wrong")
		}
	}

}
func (parser *Parser) ParseExpr() Expression {
	cursor := parser.cursor
	var prev Expression = nil

	for {
		curr := cursor.Current()
		if curr == LINE_TERMINATOR || curr == '=' {
			break
		}
		if prev == nil {
			if unicode.IsLetter(rune(curr)) {
				prev = parser.ParseIdent()
			} else if curr == '!' {
				prev = parser.ParseNot()
			} else if curr == '(' {
				prev = parser.ParseBr()
			}
		} else if binOp, ok := binOpChar[curr]; ok {
			cursor.RequireNext(curr)
			next := parser.ParseExpr()
			prev = NewBinary(binOp, prev, next)

		} else {
			break
		}
	}
	return prev
}
func (parser *Parser) ParseBr() Unarian {
	parser.cursor.RequireNext('(')
	value := parser.ParseExpr()
	parser.cursor.RequireNext(')')
	return NewUnarian(BR, value)
}
func (parser *Parser) ParseNot() Unarian {
	parser.cursor.RequireNext('!')
	parser.cursor.RequireNext('(')
	value := parser.ParseExpr()
	parser.cursor.RequireNext(')')

	return NewUnarian(NOT, value)
}
func (parser *Parser) ParseIdent() Ident {
	name := string(parser.cursor.Current())

	ident := NewIdent(name)
	parser.cursor.index++

	return ident
}
func NewParser(code []byte) *Parser {
	return &Parser{
		cursor: NewCursor(code),
		exprs:  make([]Expression, 0),
	}
}

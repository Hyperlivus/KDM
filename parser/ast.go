package parser

type ExprTag int

const WRONG_LEN = "wrong length of children"
const (
	IDENT ExprTag = iota
	BINARY
	UNARIAN
)

type Expression interface {
	Tag() ExprTag
	Children() []Expression
	SetChildren(chidren []Expression) Expression
}
type Ident struct {
	Name string
}

func (Ident) Tag() ExprTag {
	return IDENT
}
func (ident Ident) Children() []Expression {
	return make([]Expression, 0)
}
func (ident Ident) SetChildren(children []Expression) Expression {
	return ident
}
func NewIdent(name string) Ident {
	return Ident{
		Name: name,
	}
}

type BinOperand int

const (
	UNIT BinOperand = iota
	INTER
	MUL
	SUB
	SYM_SUB
)

var AllBinOperands = []BinOperand{
	UNIT,
	INTER,
	MUL,
	SUB,
	SYM_SUB,
}

type Binary struct {
	Op  BinOperand
	Lhs Expression
	Rhs Expression
}

func (bin Binary) Children() []Expression {
	return []Expression{bin.Lhs, bin.Rhs}
}
func (bin Binary) Tag() ExprTag {
	return BINARY
}
func (bin Binary) SetChildren(children []Expression) Expression {
	if len(children) != 2 {
		panic(WRONG_LEN)
	}
	return Binary{
		Op:  bin.Op,
		Lhs: children[0],
		Rhs: children[1],
	}
}
func NewBinary(op BinOperand, lhs Expression, rhs Expression) Binary {
	return Binary{
		Op:  op,
		Lhs: lhs,
		Rhs: rhs,
	}
}

type UnOperand int

const (
	NOT UnOperand = iota
	BR
)

var AllUnarians = []UnOperand{
	NOT, BR,
}

type Unarian struct {
	Op    UnOperand
	Value Expression
}

func (un Unarian) SetChildren(children []Expression) Expression {
	if len(children) != 1 {
		panic(WRONG_LEN)
	}
	return Unarian{
		Op:    un.Op,
		Value: children[0],
	}
}
func (un Unarian) Children() []Expression {
	return []Expression{un.Value}
}
func (Unarian) Tag() ExprTag {
	return UNARIAN
}
func NewUnarian(op UnOperand, value Expression) Unarian {
	return Unarian{
		Op:    op,
		Value: value,
	}
}

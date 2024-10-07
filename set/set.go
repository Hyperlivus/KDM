package set

//	"kdm/matf"

type ExprTag int

const (
	SET ExprTag = iota
	VARIABLE
)

type Expression interface {
	Tag() ExprTag
	Base() uint
	Size() int
}

type Set uint

func (Set) Tag() ExprTag {
	return SET
}
func (set Set) Base() uint {
	size := set.Size()
	var x uint = 1 << (size)
	return uint(set) ^ x
}
func (set Set) Size() int {
	x := uint(1)
	n := 0
	for x <= uint(set) {
		x = x << 1
		n++
	}
	return n - 1
}
func NewSet(size int) Set {
	res := 1 << size
	return Set(res)
}
func (set Set) SetIndex(index int) Set {
	var x uint = 1 << index
	return Set(uint(set) | x)
}
func (set Set) Index(index int) uint {
	var x uint = uint(set) >> uint(index)
	return x & uint(0b1)
}

type Variable uint

func (Variable) Tag() ExprTag {
	return VARIABLE
}
func (v Variable) Size() int {
	x := uint(1)
	n := 0
	for x < uint(v) {
		x = x << 1
		n++
	}

	return n - 1
}
func (v Variable) Base() uint {
	size := v.Size()
	var x uint = 1 << (size)
	return uint(v) ^ x
}
func (v Variable) Get() [2]uint {
	size := v.Size()
	n := int(size / 2)
	base := v.Base()
	var flag uint = 1 << n

	b1 := base % flag
	b0 := base >> n
	return [2]uint{b0, b1}
}
func (v Variable) GetSets() []Set {
	x := v.Get()

	return []Set{Set(x[0]), Set(x[1])}
}
func NewVariable(univ Set, size int) Variable {
	x := uint(1)
	x = x << (size * 2)
	x = x | univ.Base()
	return Variable(x)
}

type Operand func(x uint, y uint) uint

func UseOperand(a Expression, b Expression, op Operand) Expression {
	if a.Tag() == b.Tag() {
		size := a.Size()
		var flag uint = 1 << (size)

		aBase := a.Base()
		bBase := b.Base()

		res := op(aBase, bBase)

		res = flag ^ res

		if a.Tag() == SET {
			return Set(res)
		} else if a.Tag() == VARIABLE {
			return Variable(res)
		}
	} else if a.Tag() == VARIABLE && b.Tag() == SET {
		v := a.(Variable)
		s := b.(Set)

		size := a.Size()
		base := s.Base()

		bases := v.Get()
		b0, b1 := bases[0], bases[1]
		x0, x1 := op(b0, base), op(b1, base)
		var flag uint = 1 << (size)
		res := (flag ^ (x0 << (size / 2))) ^ x1
		return Variable(res)
	} else if b.Tag() == VARIABLE && a.Tag() == SET {
		v := b.(Variable)
		s := a.(Set)
		size := b.Size()
		base := s.Base()
		bases := v.Get()
		b0, b1 := bases[0], bases[1]
		x0, x1 := op(base, b0), op(base, b1)
		var flag uint = 1 << (size)
		res := (flag ^ (x0 << (size / 2))) ^ x1
		return Variable(res)
	}
	panic("wrong tag of expression")
}

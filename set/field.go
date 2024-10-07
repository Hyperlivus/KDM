package set

import (
	"kdm/matf"
)

type Field struct {
	names      []string
	size       int
	parts      []int
	partsCount int

	univ Set
	none Set
	sets map[string]Set
}

func (field *Field) None() Set {
	return field.none
}
func (field *Field) Names() []string {
	return field.names
}
func (field *Field) PartsCount() int {
	return field.partsCount
}
func (field *Field) Sets() map[string]Set {
	return field.sets
}
func (field *Field) Size() int {
	return field.size
}
func (field *Field) Univ() Set {
	return field.univ
}
func (field *Field) initParts() {

	field.partsCount = 0
	field.parts = make([]int, field.size+1)

	for i := 0; i <= field.size; i++ {
		c := matf.C(field.size, i)
		field.partsCount += c
		field.parts[i] = c
	}

	field.univ = NewSet(field.partsCount)
	for i := 0; i <= field.size+1; i++ {
		field.univ = field.univ.SetIndex(i)
	}
	field.none = NewSet(field.partsCount)
}
func (field *Field) initSets() {
	sets := make(map[string]Set)

	for i := 0; i < field.size; i++ {
		set := NewSet(field.partsCount)
		index := 0
		for j := 1; j < len(field.parts); j++ {
			partSize := field.parts[j]
			count := matf.C(field.size-1, j-1)
			for k := 0; k < count; k++ {
				offset := (k*2 + i) % partSize
				set = set.SetIndex(offset + index)
			}
			index += partSize
		}
		sets[field.names[i]] = set
	}
	field.sets = sets
}
func NewField(names []string) *Field {
	field := &Field{
		names: names,
	}
	field.size = len(names)
	field.initParts()
	field.initSets()

	return field
}

package solver

import (
	"kdm/set"
)

func Set2ToString(s set.Set) string {
	b := s.Base()

	if b == 0 {
		return "null"
	} else if b == 1 {
		return "a / b"
	} else if b == 2 {
		return "b / a"
	} else if b == 3 {
		return "a - b"
	} else if b == 4 {
		return "a^b"
	} else if b == 5 {
		return "a"
	} else if b == 6 {
		return "b"
	} else if b == 7 {
		return "a + b"
	} else if b == 8 {
		return "!(a + b)"
	} else if b == 9 {
		return "!b"
	} else if b == 10 {
		return "!a"
	} else if b == 11 {
		return "!(a^b)"
	} else if b == 12 {
		return "!(a - b)"
	} else if b == 13 {
		return "!(b / a)"
	} else if b == 14 {
		return "!(a / b)"
	} else if b == 15 {
		return "U"
	}
	return "uWu"
}
func SetToString(field *set.Field, s set.Set) string {
	if field.Size() == 2 {
		return Set2ToString(s)
	}
	return "sosi yaytsa ya zaebalsa"
}

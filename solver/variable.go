package solver

import (
	"kdm/set"

	"golang.org/x/exp/slices"
)

func GetAllVariants(field *set.Field, variable set.Variable) ([]set.Set, []set.Set) {

	univ := field.Univ()
	sets := make([]set.Set, 0)
	values := make([]set.Set, 0)

	bases := variable.GetSets()
	curr := set.NewSet(field.PartsCount())

	for curr <= univ {

		newSet := set.NewSet(field.PartsCount())
		for i := 0; i < field.PartsCount(); i++ {
			index := curr.Index(i)
			el := bases[index].Index(i)

			if el == 1 {
				newSet = newSet.SetIndex(i)
			}
		}
		if !slices.Contains(sets, newSet) {
			values = append(values, curr)
			sets = append(sets, newSet)
		}
		x := set.Set(uint(curr) + 1)
		curr = set.Set(x)
	}

	return sets, values
}

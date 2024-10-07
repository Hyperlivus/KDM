package solver

import (
	pars "kdm/parser"

	slices "golang.org/x/exp/slices"
)

var VariableNames = []string{
	"X", "Y", "Z",
}

func FindIdentsName(expr pars.Expression) []string {
	names := make([]string, 0)
	children := expr.Children()

	if expr.Tag() == pars.IDENT {
		ident := expr.(pars.Ident)
		if !slices.Contains(VariableNames, ident.Name) {
			return []string{ident.Name}
		} else {
			return make([]string, 0)
		}
	}

	for _, child := range children {
		if child.Tag() == pars.IDENT {
			ident := child.(pars.Ident)
			name := ident.Name
			if !slices.Contains(names, name) &&
				!slices.Contains(VariableNames, name) {
				names = append(names, name)

			}
		} else {
			next := FindIdentsName(child)
			for _, name := range next {
				if !slices.Contains(names, name) &&
					!slices.Contains(VariableNames, name) {
					names = append(names, name)
				}
			}
		}
	}

	return names
}

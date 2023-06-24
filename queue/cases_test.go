package queue

import "github.com/marciodojr/datastructures/element"

var qCases = []struct {
	desc   string
	values []element.Element
}{
	{
		desc:   "Add10Remove2Peek1Remove1Peek2IsEmptyFalse",
		values: []element.Element{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

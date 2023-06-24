package stack

import "github.com/marciodojr/datastructures/element"

var pushCases = []struct {
	desc    string
	pushArr []element.Element
	popArr  []element.Element
}{
	{
		desc:    "Push10Pop10",
		pushArr: []element.Element{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		popArr:  []element.Element{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	},
}

var isEmptyCases = []struct {
	desc    string
	pushArr []element.Element
	popArr  []element.Element
	isEmpty bool
}{
	{
		desc:    "NoPushIsEmptyTrue",
		pushArr: []element.Element{},
		popArr:  []element.Element{},
		isEmpty: true,
	},
	{
		desc:    "PushPopIsEmptyTrue",
		pushArr: []element.Element{1, 2},
		popArr:  []element.Element{2, 1},
		isEmpty: true,
	},
	{
		desc:    "Push2Pop1IsEmptyFalse",
		pushArr: []element.Element{1, 2},
		popArr:  []element.Element{2},
		isEmpty: false,
	},
}

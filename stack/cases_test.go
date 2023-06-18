package stack

var pushCases = []struct {
	desc    string
	pushArr []Element
	popArr  []Element
}{
	{
		desc:    "Push10Pop10",
		pushArr: []Element{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		popArr:  []Element{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	},
}

var isEmptyCases = []struct {
	desc    string
	pushArr []Element
	popArr  []Element
	isEmpty bool
}{
	{
		desc:    "NoPushIsEmptyTrue",
		pushArr: []Element{},
		popArr:  []Element{},
		isEmpty: true,
	},
	{
		desc:    "PushPopIsEmptyTrue",
		pushArr: []Element{1, 2},
		popArr:  []Element{2, 1},
		isEmpty: true,
	},
	{
		desc:    "Push2Pop1IsEmptyFalse",
		pushArr: []Element{1, 2},
		popArr:  []Element{2},
		isEmpty: false,
	},
}

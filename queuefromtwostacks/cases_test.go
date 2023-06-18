package queuefromtwostacks

import "github.com/marciodojr/datastructures/stack"

var qCases = []struct {
	desc       string
	enqueueArr []stack.Element
	dequeueArr []stack.Element
}{
	{
		desc:       "enqueueAndPeek10DequeueAndPeek10",
		enqueueArr: []stack.Element{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		dequeueArr: []stack.Element{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

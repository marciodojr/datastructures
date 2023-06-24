package queuefromtwostacks

import "github.com/marciodojr/datastructures/element"

var qCases = []struct {
	desc       string
	enqueueArr []element.Element
	dequeueArr []element.Element
}{
	{
		desc:       "enqueueAndPeek10DequeueAndPeek10",
		enqueueArr: []element.Element{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		dequeueArr: []element.Element{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

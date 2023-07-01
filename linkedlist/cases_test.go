package linkedlist

import "github.com/marciodojr/datastructures/element"

var lCases = []struct {
	desc   string
	values []element.Element
}{
	{
		desc:   "Append10AndPeek10",
		values: []element.Element{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

var cycleCases = []struct {
	desc     string
	list     *List
	hasCycle bool
}{
	{
		desc: "EmptyList",
		list: (func() *List {
			return New()
		})(),
		hasCycle: false,
	},
	{
		desc: "OneNodeLinearList",
		list: (func() *List {
			return New()
		})(),
		hasCycle: false,
	},
	{
		desc: "TwoNodesLinearList",
		list: (func() *List {
			return New()
		})(),
		hasCycle: false,
	},
	{
		desc: "ManyNodesLinearList",
		list: (func() *List {
			l := New()

			l.Append(1)
			l.Append(2)
			l.Append(3)
			l.Append(4)
			l.Append(5)
			l.Append(6)

			return l
		})(),
		hasCycle: false,
	},
	{
		desc: "SingleNodeCycle",
		list: (func() *List {
			l := New()

			l.Append(1)

			l.head.Next = l.head

			return l
		})(),
		hasCycle: true,
	},
	{
		desc: "TwoNodesCycle",
		list: (func() *List {
			l := New()

			l.Append(1)
			l.Append(2)

			n1, _ := l.prevNodeOfPos(1)
			n1.Next = l.head

			return l
		})(),
		hasCycle: true,
	},
	{
		desc: "ThreeNodesCycle",
		list: (func() *List {
			l := New()

			l.Append(1)
			l.Append(2)
			l.Append(3)

			n1, _ := l.prevNodeOfPos(2)
			n1.Next = l.head

			return l
		})(),
		hasCycle: true,
	},
	{
		desc: "FourNodesCycleStartingAtTheFirstNode",
		list: (func() *List {
			l := New()

			l.Append(1)
			l.Append(2)
			l.Append(3)
			l.Append(4)

			n1, _ := l.prevNodeOfPos(3)
			n1.Next = l.head

			return l
		})(),
		hasCycle: true,
	},
	{
		desc: "FourNodesCycleStartingAtTheSecondNode",
		list: (func() *List {
			l := New()

			l.Append(1)
			l.Append(2)
			l.Append(3)
			l.Append(4)

			n1, _ := l.prevNodeOfPos(4)
			n2, _ := l.prevNodeOfPos(1)
			n1.Next = n2.Next

			return l
		})(),
		hasCycle: true,
	},
}

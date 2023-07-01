package element

type Element int

const ZeroVal = 0

type Node struct {
	E    Element
	Next *Node
}

func NewNode(e Element) *Node {
	return &Node{e, nil}
}

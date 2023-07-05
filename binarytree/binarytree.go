package binarytree

import "github.com/marciodojr/datastructures/element"

type Node struct {
	e     element.Element
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func New(e element.Element) *Tree {
	return &Tree{
		NewNode(e),
	}
}

func NewNode(e element.Element) *Node {
	return &Node{e, nil, nil}
}

func (t *Tree) Add(e element.Element) {
	c := t.root

	if c == nil {
		t.root = NewNode(e)

		return
	}

	for {
		if c.e < e {
			if c.right == nil {
				c.right = NewNode(e)
				break
			}

			c = c.right
			continue
		}

		if c.left == nil {
			c.left = NewNode(e)
			break
		}

		c = c.left
	}
}

func (t *Tree) Print() []element.Element {
	c := t.root
	v := []element.Element{}
	if c.left != nil {
		nt := &Tree{c.left}
		v = nt.Print()
	}

	v = append(v, c.e)

	if c.right != nil {
		nt := &Tree{c.right}
		v = append(v, nt.Print()...)
	}

	return v
}

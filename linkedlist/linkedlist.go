package linkedlist

import (
	"errors"
	"fmt"
	"sync"

	"github.com/marciodojr/datastructures/element"
)

type List struct {
	n  *element.Node
	mx sync.RWMutex
}

var errPrevIndexInvalid = errors.New("invalid index")
var errPrevIndexOutOfRange = errors.New("index out of range")

func New() *List {
	return &List{}
}

func (l *List) IsEmpty() bool {
	l.mx.RLock()
	defer l.mx.RUnlock()

	return l.n == nil
}

func (l *List) Peek(i int) (element.Element, error) {
	l.mx.RLock()
	defer l.mx.RUnlock()

	if i == 0 {
		return l.n.E, nil
	}

	p, err := l.prevNodeOfPos(i)
	if err != nil {
		return element.ZeroVal, fmt.Errorf("impossible to peek: %w", err)
	}

	return p.Next.E, nil
}

func (l *List) Append(e element.Element) {
	if l.IsEmpty() {
		l.mx.Lock()
		defer l.mx.Unlock()

		l.n = element.NewNode(e)

		return
	}
	l.mx.Lock()
	defer l.mx.Unlock()

	var p, n *element.Node
	var c int

	n = l.n
	for n != nil {
		c++
		p = n
		n = n.Next
	}

	p.Next = element.NewNode(e)
}

func (l *List) Remove(i int) (element.Element, error) {
	if l.IsEmpty() {
		return element.ZeroVal, errors.New("impossible to remove: List is empty")
	}

	l.mx.Lock()
	defer l.mx.Unlock()

	n := l.n
	if i == 0 {
		l.n = l.n.Next
	} else {
		p, err := l.prevNodeOfPos(i)
		if err != nil {
			return element.ZeroVal, fmt.Errorf("impossible to remove: %w", err)
		}

		n = p.Next
		p.Next = n.Next
	}

	return n.E, nil
}

func (l *List) prevNodeOfPos(i int) (*element.Node, error) {
	if i < 1 {
		return nil, errPrevIndexInvalid
	}

	n := l.n
	p := n
	c := 0

	for n != nil && c < i {
		c++
		p = n
		n = n.Next
	}

	if c != i || n == nil {
		return nil, errPrevIndexOutOfRange
	}

	return p, nil
}

func (l *List) HasCycle() bool {
	if l.IsEmpty() {
		return false
	}

	l.mx.RLock()
	defer l.mx.RUnlock()

	s := l.n
	f := s.Next

	for f != nil && f.Next != nil {
		if s == f {
			return true
		}

		f = f.Next.Next
		s = s.Next
	}

	return false
}

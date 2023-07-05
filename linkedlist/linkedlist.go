package linkedlist

import (
	"errors"
	"fmt"
	"sync"

	"github.com/marciodojr/datastructures/element"
)

type List struct {
	head *element.Node
	mx   sync.RWMutex
}

var errPrevIndexInvalid = errors.New("invalid index")
var errPrevIndexOutOfRange = errors.New("index out of range")

func New() *List {
	return &List{}
}

func (l *List) IsEmpty() bool {
	l.mx.RLock()
	defer l.mx.RUnlock()

	return l.head == nil
}

func (l *List) Peek(i int) (element.Element, error) {
	l.mx.RLock()
	defer l.mx.RUnlock()

	if i == 0 {
		return l.head.E, nil
	}

	p, err := l.prevNodeOfPos(i)
	if err != nil || p.Next == nil {
		return element.ZeroVal, fmt.Errorf("impossible to peek: index out of range")
	}

	return p.Next.E, nil
}

func (l *List) Append(e element.Element) {
	if l.IsEmpty() {
		l.mx.Lock()
		defer l.mx.Unlock()

		l.head = element.NewNode(e)

		return
	}
	l.mx.Lock()
	defer l.mx.Unlock()

	n := l.head
	for n.Next != nil {
		n = n.Next
	}

	n.Next = element.NewNode(e)
}

func (l *List) Prepend(e element.Element) {
	if l.IsEmpty() {
		l.Append(e)
	}

	l.mx.Lock()
	defer l.mx.Unlock()

	newHead := element.NewNode(e)
	newHead.Next = l.head
	l.head = newHead
}

func (l *List) RemoveElement(e element.Element) {
	n := l.head
	if n.E == e {
		l.head = n.Next

		return
	}

	for n.Next != nil {
		if n.Next.E == e {
			n.Next = n.Next.Next

			return
		}

		n = n.Next
	}
}

func (l *List) RemoveAt(i int) (element.Element, error) {
	if l.IsEmpty() {
		return element.ZeroVal, errors.New("impossible to remove: List is empty")
	}

	l.mx.Lock()
	defer l.mx.Unlock()

	n := l.head
	if i == 0 {
		l.head = l.head.Next

		return n.E, nil
	}

	p, err := l.prevNodeOfPos(i)
	if err != nil || p.Next == nil {
		return element.ZeroVal, fmt.Errorf("impossible to remove: index out of range")
	}

	n = p.Next
	p.Next = n.Next

	return n.E, nil
}

func (l *List) prevNodeOfPos(i int) (*element.Node, error) {
	if i < 1 {
		return nil, errPrevIndexInvalid
	}

	n := l.head
	c := 0

	for n.Next != nil && c < i-1 {
		c++
		n = n.Next
	}

	if c != i-1 {
		return nil, errPrevIndexOutOfRange
	}

	return n, nil
}

func (l *List) HasCycle() bool {
	if l.IsEmpty() {
		return false
	}

	l.mx.RLock()
	defer l.mx.RUnlock()

	s := l.head
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

package stack

import (
	"errors"
	"sync"

	"github.com/marciodojr/datastructures/element"
)

type Stack struct {
	head *element.Node
	mx   sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{nil, sync.RWMutex{}}
}

func (s *Stack) Push(e element.Element) {
	n := element.NewNode(e)
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.head == nil {
		s.head = n
	} else {
		h := s.head
		n.Next = h
		s.head = n
	}
}

func (s *Stack) Pop() (element.Element, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	s.mx.Lock()
	defer s.mx.Unlock()

	h := s.head
	s.head = s.head.Next

	return h.E, nil
}

func (s *Stack) Peek() (element.Element, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	s.mx.RLock()
	s.mx.RUnlock()

	return s.head.E, nil
}

func (s *Stack) IsEmpty() bool {
	s.mx.RLock()
	defer s.mx.RUnlock()

	return s.head == nil
}

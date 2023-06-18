package stack

import (
	"errors"
	"sync"
)

type Element int

type Stack struct {
	c  []Element
	mx sync.RWMutex
}

func NewStack() *Stack {
	c := []Element{}

	return &Stack{c, sync.RWMutex{}}
}

func (s *Stack) Push(e Element) {
	s.mx.Lock()
	s.c = append(s.c, e)
	s.mx.Unlock()
}

func (s *Stack) Pop() (Element, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	s.mx.Lock()

	last := len(s.c) - 1
	e := s.c[last]
	s.c = s.c[:last]

	s.mx.Unlock()

	return e, nil
}

func (s *Stack) Top() (Element, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	s.mx.RLock()
	e := s.c[len(s.c)-1]
	s.mx.RUnlock()

	return e, nil
}

func (s *Stack) IsEmpty() bool {
	s.mx.RLock()
	l := len(s.c)
	s.mx.RUnlock()

	return l == 0
}

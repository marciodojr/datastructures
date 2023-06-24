package queue

import (
	"errors"
	"fmt"

	"github.com/marciodojr/datastructures/element"
)

type node struct {
	e    element.Element
	next *node
}

type Queue struct {
	head *node
	tail *node
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Peek() (element.Element, error) {
	if q.IsEmpty() {
		return element.ZeroVal, errors.New("impossible to peek. Queue is empty")
	}

	return q.head.e, nil
}

func (q *Queue) Add(e element.Element) {
	n := &node{e, nil}

	if q.IsEmpty() {
		q.tail = n
		q.head = n

		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *Queue) Remove() (element.Element, error) {
	if q.IsEmpty() {
		return element.ZeroVal, errors.New("impossible to remove. Queue is empty")
	}

	n := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return n.e, nil
}

func (q *Queue) String() string {
	s := []element.Element{}

	n := q.head
	for n != nil {
		s = append(s, n.e)
		n = n.next
	}

	return fmt.Sprint(s)
}

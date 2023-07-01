package queue

import (
	"errors"
	"fmt"
	"sync"

	"github.com/marciodojr/datastructures/element"
)

type Queue struct {
	head *element.Node
	tail *element.Node
	mx   sync.RWMutex
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	q.mx.RLock()
	defer q.mx.RUnlock()

	return q.head == nil
}

func (q *Queue) Peek() (element.Element, error) {
	if q.IsEmpty() {
		return element.ZeroVal, errors.New("impossible to peek. Queue is empty")
	}

	q.mx.RLock()
	defer q.mx.RUnlock()

	return q.head.E, nil
}

func (q *Queue) Add(e element.Element) {
	n := element.NewNode(e)

	if q.IsEmpty() {
		q.mx.Lock()
		defer q.mx.Unlock()

		q.tail = n
		q.head = n

		return
	}

	q.mx.Lock()
	defer q.mx.Unlock()

	q.tail.Next = n
	q.tail = n
}

func (q *Queue) Remove() (element.Element, error) {
	if q.IsEmpty() {
		return element.ZeroVal, errors.New("impossible to remove. Queue is empty")
	}
	q.mx.Lock()
	defer q.mx.Unlock()

	n := q.head
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}

	return n.E, nil
}

func (q *Queue) String() string {
	s := []element.Element{}

	n := q.head
	for n != nil {
		s = append(s, n.E)
		n = n.Next
	}

	return fmt.Sprint(s)
}

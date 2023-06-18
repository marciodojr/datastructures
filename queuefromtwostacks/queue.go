package queuefromtwostacks

import (
	"fmt"
	"sync"

	"github.com/marciodojr/datastructures/stack"
)

type QueueFromTwoStacks struct {
	stackNewestOnTop *stack.Stack
	stackOldestOnTop *stack.Stack
	mx               sync.RWMutex
}

func NewQueueFromTwoStacks() *QueueFromTwoStacks {
	return &QueueFromTwoStacks{
		stack.NewStack(),
		stack.NewStack(),
		sync.RWMutex{},
	}
}

// Insert a new element
func (q *QueueFromTwoStacks) Enqueue(e stack.Element) {
	q.mx.Lock()
	defer q.mx.Unlock()

	q.stackNewestOnTop.Push(e)
}

// Return the oldest element
func (q *QueueFromTwoStacks) Peek() (stack.Element, error) {
	q.mx.RLock()
	defer q.mx.RUnlock()

	if q.stackOldestOnTop.IsEmpty() {
		q.shiftFromNewestToOldest()
	}

	e, err := q.stackOldestOnTop.Top()

	if err != nil {
		return 0, fmt.Errorf("no element left to peek: %w", err)
	}

	return e, nil
}

// Remove the oldest element
func (q *QueueFromTwoStacks) Dequeue() (stack.Element, error) {
	q.mx.Lock()
	defer q.mx.Unlock()

	if q.stackOldestOnTop.IsEmpty() {
		q.shiftFromNewestToOldest()
	}

	e, err := q.stackOldestOnTop.Pop()

	if err != nil {
		return 0, fmt.Errorf("no element left to dequeue: %w", err)
	}

	return e, nil
}

func (q *QueueFromTwoStacks) shiftFromNewestToOldest() {
	n, o := q.stackNewestOnTop, q.stackOldestOnTop

	for !n.IsEmpty() {
		e, _ := n.Pop()
		o.Push(e)
	}
}

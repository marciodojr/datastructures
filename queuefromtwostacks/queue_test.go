package queuefromtwostacks

import "testing"

func TestEnqueuePeekDequeue(t *testing.T) {
	q := NewQueueFromTwoStacks()

	for _, tc := range qCases {
		t.Run(tc.desc, func(t *testing.T) {
			for i, e := range tc.enqueueArr {
				q.Enqueue(e)
				oldest, err := q.Peek()
				if err != nil {
					t.Errorf("want nil receive %v on peek after enqueue %d", err, i)
				}

				if oldest != tc.enqueueArr[0] {
					t.Errorf("want %v receive %v on peek after enqueue %d", e, oldest, i)
				}
			}

			for i, e := range tc.dequeueArr {
				oldest, err := q.Peek()
				if err != nil {
					t.Errorf("want nil receive %v on top before dequeue %d", err, i)
				}

				if oldest != e {
					t.Errorf("want %v receive %v on top before dequeue %d", e, oldest, i)
				}

				r, err := q.Dequeue()

				if err != nil {
					t.Errorf("want nil receive %v on dequeue %d", err, i)
				}

				if r != e {
					t.Errorf("want %v receive %v on dequeue %d", e, r, i)
				}
			}
		})
	}
}

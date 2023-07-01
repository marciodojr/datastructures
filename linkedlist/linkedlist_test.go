package linkedlist

import (
	"testing"
)

func TestAppendPeekRemovePeek(t *testing.T) {
	l := New()
	for _, tc := range lCases {
		t.Run(tc.desc, func(t *testing.T) {
			// add 10
			for i, v := range tc.values {
				l.Append(v)
				e, err := l.Peek(i)

				if err != nil {
					t.Errorf("peek at position %d, found error %v want error nil", i, err)
				}

				if e != v {
					t.Errorf("peek at position %d, found %v want %v", i, e, v)
				}
			}

			e, _ := l.Remove(3)
			v := tc.values[3]
			if e != v {
				t.Errorf("peek at position %d, found %v want %v", 3, e, v)
			}
			e, _ = l.Remove(5)
			v = tc.values[6]
			if e != v {
				t.Errorf("peek at position %d, found %v want %v", 5, e, v)
			}

			e, _ = l.Remove(7)
			v = tc.values[9]
			if e != v {
				t.Errorf("peek at position %d, found %v want %v", 7, e, v)
			}
		})
	}
}

func TestHasCycle(t *testing.T) {
	for _, tc := range cycleCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := tc.list.HasCycle()
			w := tc.hasCycle

			if w != res {
				t.Errorf("HasCycle: found %v want %v", res, w)
			}
		})
	}
}

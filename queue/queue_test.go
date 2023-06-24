package queue

import (
	"testing"
)

func TestPeekAddRemoveIsEmpty(t *testing.T) {
	q := New()
	// Add10Remove2Peek1Remove1Peek2IsEmptyFalse
	for _, tc := range qCases {
		t.Run(tc.desc, func(t *testing.T) {
			// add 10
			for _, v := range tc.values {
				q.Add(v)
			}
			// remove 2
			for i := 0; i < 2; i++ {
				want := tc.values[i]
				e, err := q.Remove()
				if err != nil {
					t.Errorf("remove at position %d, found error %v want error nil", i, err)
				}
				if e != want {
					t.Errorf("remove at position %d, found %v want %v", i, e, want)
				}
			}

			// peek 1
			want := tc.values[2]
			e, err := q.Peek()
			if err != nil {
				t.Errorf("peek at position %d, found error %v want error nil", 2, err)
			}
			if e != want {
				t.Errorf("remove at position %d, found %v want %v", 2, e, want)
			}

			// remove 1
			want = tc.values[2]
			e, err = q.Remove()
			if err != nil {
				t.Errorf("remove at position %d, found error %v want error nil", 2, err)
			}
			if e != want {
				t.Errorf("remove at position %d, found %v want %v", 2, e, want)
			}

			// peek 2
			for i := 0; i < 2; i++ {
				want = tc.values[3]
				e, err = q.Peek()
				if err != nil {
					t.Errorf("peek at position %d, found error %v want error nil", 3, err)
				}
				if e != want {
					t.Errorf("peek at position %d, found %v want %v", 3, e, want)
				}
			}

			if q.IsEmpty() {
				t.Errorf("check emptiness, found %v want %v", true, false)
			}
		})
	}
}

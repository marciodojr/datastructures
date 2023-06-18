package stack

import "testing"

func TestPushPop(t *testing.T) {
	s := NewStack()

	for _, tc := range pushCases {
		t.Run(tc.desc, func(t *testing.T) {
			for i, e := range tc.pushArr {
				s.Push(e)
				currentTop, err := s.Top()
				if err != nil {
					t.Errorf("want nil receive %v on top after push %d", err, i)
				}

				if currentTop != e {
					t.Errorf("want %v receive %v on top after push %d", e, currentTop, i)
				}
			}

			for i, e := range tc.popArr {
				currentTop, err := s.Top()
				if err != nil {
					t.Errorf("want nil receive %v on top before pop %d", err, i)
				}

				if currentTop != e {
					t.Errorf("want %v receive %v on top before top %d", e, currentTop, i)
				}

				r, err := s.Pop()
				if err != nil {
					t.Errorf("want nil receive %v on pop %d", err, i)
				}

				if r != e {
					t.Errorf("want %v receive %v on pop %d", e, r, i)
				}
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack()

	for _, tc := range isEmptyCases {
		t.Run(tc.desc, func(t *testing.T) {
			for _, e := range tc.pushArr {
				s.Push(e)
			}

			for range tc.popArr {
				s.Pop()
			}

			isEmpty := s.IsEmpty()
			if isEmpty != tc.isEmpty {
				t.Errorf("want %v receive %v", tc.isEmpty, isEmpty)
			}
		})
	}
}

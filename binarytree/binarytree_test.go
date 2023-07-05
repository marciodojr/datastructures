package binarytree

import (
	"testing"

	"github.com/marciodojr/datastructures/element"
)

func TestAdd(t *testing.T) {
	tr := New(0)

	for _, tc := range tCases {
		t.Run(tc.desc, func(t *testing.T) {
			for _, v := range tc.values {
				tr.Add(v)
			}

			res := tr.Print()

			for i, w := range []element.Element{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
				r := res[i]
				if r != w {
					t.Errorf("found %v want %v", r, w)
				}
			}
		})
	}
}

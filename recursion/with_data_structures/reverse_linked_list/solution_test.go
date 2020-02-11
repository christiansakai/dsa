package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	l := &Node{
		Value: 4,
		Next: &Node{
			Value: 3,
			Next: &Node{
				Value: 11,
				Next: &Node{
					Value: 7,
					Next:  nil,
				},
			},
		},
	}

	got := Solve(l)

	want := &Node{
		Value: 7,
		Next: &Node{
			Value: 11,
			Next: &Node{
				Value: 3,
				Next: &Node{
					Value: 4,
					Next:  nil,
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

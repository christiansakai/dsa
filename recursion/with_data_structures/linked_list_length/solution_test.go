package solution

import "testing"

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
	want := 4

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

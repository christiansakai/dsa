package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input  [][]int
		output bool
	}{
		{[][]int{
			[]int{1, 3},
			[]int{0, 2},
			[]int{1, 3},
			[]int{0, 2},
		}, true},
		{[][]int{
			[]int{1, 2, 3},
			[]int{0, 2},
			[]int{0, 1, 3},
			[]int{0, 2},
		}, false},
	}

	for _, tt := range tests {
		got := SolveIterative(tt.input)

		if got != tt.output {
			t.Errorf("got %t, want %t", got, tt.output)
		}
	}

	for _, tt := range tests {
		got := SolveRecursive(tt.input)

		if got != tt.output {
			t.Errorf("got %t, want %t", got, tt.output)
		}
	}
}

package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
	}

	for _, tt := range tests {
		got := Solve(tt.input)

		if got != tt.output {
			t.Errorf("got %d, want %d", got, tt.output)
		}
	}
}

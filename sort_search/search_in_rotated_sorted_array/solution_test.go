package solution

import "testing"

func TestSolution(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}

	tests := []struct {
		input  input
		output int
	}{
		{input{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{input{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
		{input{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8}, 4},
	}

	for _, tt := range tests {
		got := Solve(tt.input.nums, tt.input.target)

		if got != tt.output {
			t.Errorf("got %d, want %d", got, tt.output)
		}
	}
}

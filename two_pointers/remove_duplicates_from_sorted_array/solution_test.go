package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []int
		Output int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)
		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}

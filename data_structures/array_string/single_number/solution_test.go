package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []int
		Output int
	}{
    {[]int{2, 2, 1}, 1},
    {[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, tt := range tests {
    got := Solve(tt.Input)
		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}

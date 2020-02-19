package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []int
		Output bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)
		if got != tt.Output {
			t.Errorf("got %t, want %t", got, tt.Output)
		}
	}
}

package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input  []int
		output bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0, 2, 3}, false},
	}

	t.Run("Bottom-up Dynamic Programming", func(t *testing.T) {
		for _, tt := range tests {
			got := BottomUp(tt.input)

			if got != tt.output {
				t.Errorf("got %t, want %t", got, tt.output)
			}
		}
	})
}

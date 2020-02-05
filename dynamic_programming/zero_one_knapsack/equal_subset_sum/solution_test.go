package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []int
		Output bool
	}{
		{[]int{1, 2, 3, 4}, true},
		{[]int{1, 1, 3, 4, 7}, true},
		{[]int{2, 3, 4, 6}, false},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := TopDown(tt.Input)
			if got != tt.Output {
				t.Errorf("got %t, want %t", got, tt.Output)
			}
		}
	})

	t.Run("Bottom-up Dynamic Programming", func(t *testing.T) {
		for _, tt := range tests {
			got := BottomUp(tt.Input)
			if got != tt.Output {
				t.Errorf("got %t, want %t", got, tt.Output)
			}
		}
	})

	t.Run("Bottom-up Dynamic Programming with Minimum Space", func(t *testing.T) {
		for _, tt := range tests {
			got := BottomUpMinSpace(tt.Input)
			if got != tt.Output {
				t.Errorf("got %t, want %t", got, tt.Output)
			}
		}
	})
}

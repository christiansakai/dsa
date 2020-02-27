package solution

import "testing"

func TestSolution(t *testing.T) {
	type input struct {
		m int
		n int
	}

	tests := []struct {
		input  input
		output int
	}{
		{input{3, 2}, 3},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := Solve(tt.input.m, tt.input.n)

			if got != tt.output {
				t.Errorf("got %d, want %d", got, tt.output)
			}
		}
	})
}

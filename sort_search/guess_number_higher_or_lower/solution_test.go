package solution

import "testing"

func createGuessFn(target int) func(int) int {
	return func(guess int) int {
		if target == guess {
			return 0
		} else if target > guess {
			return 1
		} else {
			return -1
		}
	}
}

func TestSolution(t *testing.T) {
	type input struct {
		limit  int
		target int
	}

	tests := []struct {
		input  input
		output int
	}{
		{input{10, 5}, 5},
	}

	for _, tt := range tests {
		guessFn := createGuessFn(tt.input.target)
		got := Solve(tt.input.limit, guessFn)

		if got != tt.output {
			t.Errorf("got %d, want %d", got, tt.output)
		}
	}
}

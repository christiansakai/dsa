package solution

import "testing"

func createIsBadVersion(badVersion int) func(int) bool {
	return func(check int) bool {
		return check >= badVersion
	}
}

func TestSolution(t *testing.T) {
	type input struct {
		n       int
		version int
	}

	tests := []struct {
		input  input
		output int
	}{
		{input{2, 2}, 2},
		{input{5, 4}, 4},
		{input{10, 7}, 7},
		{input{8, 1}, 1},
		{input{6, 6}, 6},
	}

	for _, tt := range tests {
		isBadVersionFunc := createIsBadVersion(tt.input.version)
		solveFunc := Solve(isBadVersionFunc)
		got := solveFunc(tt.input.n)

		if got != tt.output {
			t.Errorf("got %d, want %d", got, tt.output)
		}
	}
}

package solution

import "testing"

func TestSolution(t *testing.T) {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}

	tests := []struct {
		input  string
		output bool
	}{
		{"ABCCED", true},
		{"SEE", true},
		{"ABCB", false},
	}

	for _, tt := range tests {
		got := Solve(board, tt.input)

		if got != tt.output {
			t.Errorf("got %t, want %t", got, tt.output)
		}
	}
}

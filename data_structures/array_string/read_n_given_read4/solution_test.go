package solution

import "testing"

func TestSolution(t *testing.T) {
	t.Skip()
	tests := []struct {
		Input struct {
			File string
			N    int
		}
		Output int
	}{
		{struct {
			File string
			N    int
		}{"abc", 4}, 3},
		{struct {
			File string
			N    int
		}{"abcde", 5}, 5},
		{struct {
			File string
			N    int
		}{"abcdABCD1234", 12}, 12},
	}

	for _, tt := range tests {
		read4 := createRead4(tt.Input.File)
		read := Solve(read4)

		buf := make([]byte, len(tt.Input.File))

		got := read(buf, tt.Input.N)
		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}

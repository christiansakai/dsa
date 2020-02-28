package solution

import "testing"

func createKnowFunc(graph [][]int) func(int, int) bool {
	return func(i, j int) bool {
		if graph[i][j] == 1 {
			return true
		}

		return false
	}
}

func TestSolution(t *testing.T) {
	type input struct {
		graph [][]int
		n     int
	}

	tests := []struct {
		input  input
		output int
	}{
		{input{
			graph: [][]int{
				[]int{1, 1, 0},
				[]int{0, 1, 0},
				[]int{1, 1, 1},
			},
			n: 3,
		}, 1},
		{input{
			graph: [][]int{
				[]int{1, 0, 1},
				[]int{1, 1, 0},
				[]int{0, 1, 1},
			},
			n: 3,
		}, -1},
		{input{
			graph: [][]int{
				[]int{1, 1},
				[]int{1, 1},
			},
			n: 2,
		}, -1},
	}

	for _, tt := range tests {
		knowFn := createKnowFunc(tt.input.graph)
		solutionFn := Solve(knowFn)
		got := solutionFn(tt.input.n)

		if got != tt.output {
			t.Errorf("got %d, want %d", got, tt.output)
		}
	}
}

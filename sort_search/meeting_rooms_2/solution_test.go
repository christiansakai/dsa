package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input  [][]int
		output int
	}{
		{[][]int{
			[]int{0, 30},
			[]int{5, 10},
			[]int{15, 20},
		}, 2},
		{[][]int{
			[]int{7, 10},
			[]int{2, 4},
		}, 1},
		{[][]int{
			[]int{13, 15},
			[]int{1, 13},
		}, 1},
		{[][]int{
			[]int{1, 5},
			[]int{8, 9},
			[]int{8, 9},
		}, 2},
		{[][]int{
			[]int{1293, 2986},
			[]int{848, 3846},
			[]int{4284, 5907},
			[]int{4466, 4781},
			[]int{518, 2918},
			[]int{300, 5870},
		}, 4},
	}

	for _, tt := range tests {
		got := Solve(tt.input)

		if got != tt.output {
			t.Errorf("got %d, want %d", got, tt.output)
		}
	}
}

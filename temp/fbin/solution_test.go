package solution

import "testing"

func testsolution(t *testing.T) {
	input := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 1},
	}

	output := 2

	got := solve(input)

	if got != output {
		t.Errorf("got %d, want %d", got, output)
	}
}

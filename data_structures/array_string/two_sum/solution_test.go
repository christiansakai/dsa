package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Nums   []int
			Target int
		}
		Output []int
	}{
		{struct {
			Nums   []int
			Target int
		}{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.Nums, tt.Input.Target)
		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}

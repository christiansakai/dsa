package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  struct{
      Arr []int
      K int
    }
		Output []int
	}{
    {struct{
      Arr []int
      K int
    }{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
    {struct{
      Arr []int
      K int
    }{[]int{-1, -100, 3, 99}, 2}, []int{3, 99, -1, -100}},
	}

	for _, tt := range tests {
		Solve(tt.Input.Arr,tt.Input.K)

		if !reflect.DeepEqual(tt.Input, tt.Output) {
			t.Errorf("got %v, want %v", tt.Input, tt.Output)
		}
	}
}

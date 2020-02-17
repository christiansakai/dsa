package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output []string
	}{
		{"()())()", []string{
			"(())()",
			"()()()",
		}},
		{"(a)())()", []string{
			"(a())()",
			"(a)()()",
		}},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}

package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []byte
		Output []byte
	}{
		{
			[]byte{'h', 'e', 'l', 'l', 'o'},
			[]byte{'o', 'l', 'l', 'e', 'h'},
		},
	}

	for _, tt := range tests {
		Solve(tt.Input)

		if !reflect.DeepEqual(tt.Input, tt.Output) {
			t.Errorf("got %v, want %v", tt.Input, tt.Output)
		}
	}
}

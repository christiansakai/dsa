package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	input := [][]string{
		[]string{"John", "johnsmith@mail.com", "john00@mail.com"},
		[]string{"John", "johhnybravo@mail.com"},
		[]string{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		[]string{"Mary", "mary@mail.com"},
	}

	output := [][]string{
		[]string{"John", "john00@mail", "john_newyork@mail", "johnsmith@mail"},
		[]string{"John", "johhnybravo@mail"},
		[]string{"Mary", "mary@mail.com"},
	}

	got := Solve(tt.input)

	if !reflect.DeepEqual(got, output) {
		t.Errorf("got %v, want %v", got, output)
	}
}

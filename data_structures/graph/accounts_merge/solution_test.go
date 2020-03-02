package solution

import (
	"reflect"
	"testing"
)

func testsolution(t *testing.t) {
	input := [][]string{
		[]string{"john", "johnsmith@mail.com", "john00@mail.com"},
		[]string{"john", "johhnybravo@mail.com"},
		[]string{"john", "johnsmith@mail.com", "john_newyork@mail.com"},
		[]string{"mary", "mary@mail.com"},
	}

	output := [][]string{
		[]string{"john", "john00@mail", "john_newyork@mail", "johnsmith@mail"},
		[]string{"john", "johhnybravo@mail"},
		[]string{"mary", "mary@mail.com"},
	}

	got := solve(tt.input)

	if !reflect.deepequal(got, output) {
		t.errorf("got %v, want %v", got, output)
	}
}

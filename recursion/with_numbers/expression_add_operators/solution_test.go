package solution

import "testing"

func TestSolution(t *testing.T) {
	type input struct {
		num    string
		target int
	}

	tests := []struct {
		input  input
		output []string
	}{
		{input{
			num:    "123",
			target: 6,
		}, []string{"1+2+3", "1*2*3"}},
		{input{
			num:    "232",
			target: 8,
		}, []string{"2*3+2", "2+3*2"}},
		{input{
			num:    "105",
			target: 5,
		}, []string{"1*0+5", "10-5"}},
		{input{
			num:    "00",
			target: 0,
		}, []string{"0+0", "0-0", "0*0"}},
		{input{
			num:    "3456237490",
			target: 9191,
		}, []string{}},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %q, want %q", got, tt.Output)
		}
	}
}

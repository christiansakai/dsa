package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			N1 string
			N2 string
		}
		Output string
	}{
		// {struct{
		//   N1 string
		//   N2 string
		// }{"0", "0"}, "0"},
		// {struct{
		//   N1 string
		//   N2 string
		// }{"1", "0"}, "0"},
		// {struct{
		//   N1 string
		//   N2 string
		// }{"0", "1"}, "0"},
		// {struct{
		//   N1 string
		//   N2 string
		// }{"12", "3"}, "36"},
		{struct {
			N1 string
			N2 string
		}{"3", "12"}, "36"},
		// {struct{
		//   N1 string
		//   N2 string
		// }{"12", "12"}, "144"},
		// {struct{
		//   N1 string
		//   N2 string
		// }{"25", "5"}, "125"},
		// {struct{
		//   N1 string
		//   N2 string
		// }{"9", "99"}, "891"},
		// {struct{
		//   N1 string
		//   N2 string
		// }{"80", "502"}, "40160"},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.N1, tt.Input.N2)

		if got != tt.Output {
			t.Errorf("got %q, want %q", got, tt.Output)
		}
	}
}

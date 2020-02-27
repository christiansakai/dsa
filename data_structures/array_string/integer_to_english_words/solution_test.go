package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  int
		Output string
	}{
		{0, "Zero"},
		{1, "One"},
		{10, "Ten"},
		{20, "Twenty"},
		{100, "One Hundred"},
		{123, "One Hundred Twenty Three"},
		{1000, "One Thousand"},
		{10000, "Ten Thousand"},
		{12345, "Twelve Thousand Three Hundred Forty Five"},
		{100000, "One Hundred Thousand"},
		{123456, "One Hundred Twenty Three Thousand Four Hundred Fifty Six"},
		{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{1234567891, "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)
		if got != tt.Output {
			t.Errorf("got %q, want %q", got, tt.Output)
		}
	}
}

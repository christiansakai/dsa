package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  *ListNode
		Output bool
	}{
		{&ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		}, false},
		{&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		}, false},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %t, want %t", got, tt.Output)
		}
	}
}

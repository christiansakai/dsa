package solution

import "testing"

func TestSolution(t *testing.T) {
	n1 := &ListNode{
		Val:  3,
		Next: nil,
	}

	n2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	n3 := &ListNode{
		Val:  0,
		Next: nil,
	}

	n4 := &ListNode{
		Val:  0,
		Next: nil,
	}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	got := Solve(n1)
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}

	n1 = &ListNode{
		Val:  3,
		Next: nil,
	}

	n2 = &ListNode{
		Val:  2,
		Next: nil,
	}

	n1.Next = n2

	got = Solve(n1)
	want = false

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

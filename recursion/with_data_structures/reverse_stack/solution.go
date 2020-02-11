package solution

func Solve(s *Stack) *Stack {
	if !s.IsEmpty() {
		topMost, _ := s.Pop()
		recurse(s)
		insertAtBottom(s, topMost)
	}

	return s
}

func recurse(s *Stack) {
	if s.IsEmpty() {
		return
	}

	topMost, _ := s.Pop()
	recurse(s)
	insertAtBottom(s, topMost)
}

func insertAtBottom(s *Stack, item int) {
	if s.IsEmpty() {
		s.Push(item)
		return
	}

	topMost, _ := s.Pop()
	insertAtBottom(s, item)
	s.Push(topMost)
}

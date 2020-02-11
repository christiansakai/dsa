package solution

import (
	"errors"
	"math"
)

type Stack struct {
	Arr []int
}

func New() *Stack {
	return &Stack{[]int{}}
}

func (s *Stack) Push(value int) {
	s.Arr = append(s.Arr, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Arr) == 0
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return int(math.Inf(1)), errors.New("Stack is empty")
	}

	value := s.Arr[len(s.Arr)-1]
	s.Arr = s.Arr[:len(s.Arr)-1]

	return value, nil
}

func (s *Stack) ToSlice() []int {
	slice := []int{}
	for _, el := range s.Arr {
		slice = append(slice, el)
	}

	return slice
}

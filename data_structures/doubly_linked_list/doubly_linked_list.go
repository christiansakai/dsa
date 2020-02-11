package doubly_linked_list

type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func New() *LinkedList {
	headDummy := Node{
		Value: nil,
		Prev:  nil,
		Next:  nil,
	}

	tailDummy := Node{
		Value: nil,
		Prev:  nil,
		Next:  nil,
	}

	return &LinkedList{
		Head:   &headDummy,
		Tail:   &tailDummy,
		Length: 0,
	}
}

func (l *LinkedList) PushHead(value interface{}) {
	newNode := Node{Value: value}

	if l.Length == 0 {
		l.Head.Next = &newNode
		l.Tail.Prev = &newNode
	} else {
		headNext := l.Head.Next
		headNext.Prev = &newNode

		l.Head.Next = &newNode
	}

	l.Length += 1
}

func (l *LinkedList) PushTail(value interface{}) {
	newNode := Node{Value: value}

	if l.Length == 0 {
		l.Head.Next = &newNode
		l.Tail.Prev = &newNode
	} else {
		tailPrev := l.Tail.Prev
		tailPrev.Next = &newNode

		l.Tail.Prev = &newNode
	}

	l.Length += 1
}

func (l *LinkedList) ToSlice() []interface{} {
	result := []interface{}{}

	node := l.Head.Next
	for {
		if node == nil {
			break
		}

		result = append(result, node.Value)
		node = node.Next
	}

	return result
}

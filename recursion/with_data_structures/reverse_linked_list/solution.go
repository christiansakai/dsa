package solution

func Solve(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := Solve(head.Next)

	// Change references for middle chain
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// func recurse(curr, next *Node) *Node {
//   if next == nil {
//     return curr
//   }

//   subProb := recurse(curr.Next, curr.Next.Next)
//   subProb.Next = curr

//   return subProb
// }

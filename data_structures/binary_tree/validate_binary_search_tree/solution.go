package solution

import "math"

func Solve(root *TreeNode) bool {
  if root == nil {
    return true
  }

  isValid, _, _ := recurse(root)
  return isValid
}

func recurse(root *TreeNode) (bool, int, int) {
  if root.Left == nil && root.Right == nil {
    return true, root.Val, root.Val
  }

  isValid := true
  smallest := root.Val
  biggest := root.Val

  if root.Left != nil {
    isValid = isValid && (root.Left.Val < root.Val)

    leftIsValid, leftSmallest, leftBiggest := recurse(root.Left)
    
    isValid = isValid && leftIsValid && (leftBiggest < root.Val)
    smallest = int(math.Min(float64(smallest), float64(leftSmallest)))
    biggest = int(math.Max(float64(biggest), float64(root.Val)))
  }

  if root.Right != nil {
    isValid = isValid && (root.Val < root.Right.Val) 

    rightIsValid, rightSmallest, rightBiggest := recurse(root.Right)

    isValid = isValid && rightIsValid && (root.Val < rightSmallest)
    smallest = int(math.Min(float64(smallest), float64(root.Val)))
    biggest = int(math.Max(float64(biggest), float64(rightBiggest)))
  }

  return isValid, smallest, biggest
}

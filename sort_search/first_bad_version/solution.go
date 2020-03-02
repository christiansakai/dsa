package solution

func Solve(isBadVersion func (int) bool) func(int) int {
  return func(n int) int {
    left := 1
    right := n

    for left < right {
      mid := left + (right - left) / 2
      if isBadVersion(mid) {
        right = mid
      } else {
        left = mid + 1
      }
    }

    // terminating condition is left == mid
    return left
  }
}

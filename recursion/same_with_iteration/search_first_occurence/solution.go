package solution

func Solve(arr []int, n int) int {
	if len(arr) == 0 {
		return -1
	}

	return recurse(arr, n, 0)
}

func recurse(arr []int, n, index int) int {
	if index == len(arr) {
		return -1
	}

	if arr[index] == n {
		return index
	}

	return recurse(arr, n, index+1)
}

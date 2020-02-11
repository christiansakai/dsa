package solution

func Solve(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	return recurse(arr, 0, 0)
}

func recurse(arr []int, index, sum int) int {
	if index == len(arr) {
		return sum / (len(arr))
	}

	return recurse(arr, index+1, sum+arr[index])
}

package solution

func Solve(arr []int, key int) int {
	if len(arr) == 0 {
		return 0
	}

	return recurse(arr, key, 0)
}

func recurse(arr []int, key, index int) int {
	if index == len(arr) {
		return 0
	}

	if arr[index] == key {
		return 1 + recurse(arr, key, index+1)
	}

	return recurse(arr, key, index+1)
}

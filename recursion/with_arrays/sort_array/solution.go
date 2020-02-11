package solution

func Solve(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	recurse(arr, 0, 1)

	return arr
}

func recurse(arr []int, i, j int) {
	if i == len(arr)-1 {
		return
	}

	if j == len(arr) {
		nextI := i + 1
		nextJ := nextI + 1

		recurse(arr, nextI, nextJ)
		return
	}

	if arr[i] > arr[j] {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}

	recurse(arr, i, j+1)
}

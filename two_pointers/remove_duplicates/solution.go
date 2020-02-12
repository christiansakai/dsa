package solution

func Solve(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	lastAvailablePos := 0
	i := 0
	for ; i < len(arr); i++ {
		j := i + 1
		for ; j < len(arr) && arr[i] == arr[j]; j++ {
			arr[j] = 0
		}

		arr[lastAvailablePos] = arr[i]
		lastAvailablePos += 1
		i = j
	}

	return 0
}

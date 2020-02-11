package solution

func Solve(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	return recurse(arr, 0)
}

func recurse(arr []int, index int) []int {
	if index == len(arr) {
		return []int{}
	}

	if arr[index] < 0 {
		newArr := []int{0}
		subProb := recurse(arr, index+1)

		newArr = append(newArr, subProb...)

		return newArr
	}

	newArr := []int{arr[index]}
	subProb := recurse(arr, index+1)

	newArr = append(newArr, subProb...)

	return newArr
}

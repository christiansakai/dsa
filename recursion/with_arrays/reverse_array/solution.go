package solution

func Solve(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	if len(arr) == 1 {
		return []int{arr[0]}
	}

	subArr := arr[0 : len(arr)-1]
	subProb := Solve(subArr)

	newArr := []int{arr[len(arr)-1]}
	newArr = append(newArr, subProb...)

	return newArr
}

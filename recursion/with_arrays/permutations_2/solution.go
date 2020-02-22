package solution

func Solve(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	result := [][]int{}
	visited := map[string]bool{}
	set := []int{}

	recurse(nums, &set, visited, &result)

	return result
}

func recurse(nums []int, set *[]int, visited map[string]bool, result *[][]int) {
	if len(nums) == 0 {
		copied := copySet(set)
		stringified := stringify(copied)

		if _, ok := visited[stringified]; !ok {
			*result = append(*result, copied)
		}

		visited[stringified] = true

		return
	}

	for i, el := range nums {
		newNums := createNewNumsWithout(nums, i)

		*set = append(*set, el)
		recurse(newNums, set, visited, result)
		*set = (*set)[:len(*set)-1]
	}
}

func copySet(set *[]int) []int {
	newSet := make([]int, len(*set))
	for i, el := range *set {
		newSet[i] = el
	}

	return newSet
}

func stringify(set []int) string {
	b := make([]byte, len(set))

	for i, el := range set {
		b[i] = byte(el + 48)
	}

	return string(b)
}

func createNewNumsWithout(nums []int, index int) []int {
	newNums := []int{}
	for i, n := range nums {
		if i != index {
			newNums = append(newNums, n)
		}
	}

	return newNums
}

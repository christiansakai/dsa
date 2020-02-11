package solution

func Solve(str string) bool {
	if len(str)%2 != 0 {
		return false
	}

	return recurse([]byte(str), 0, 0)
}

func recurse(arr []byte, index, par int) bool {
	if index == len(arr) {
		return par == 0
	}

	if par < 0 {
		return false
	}

	if arr[index] == byte('(') {
		return recurse(arr, index+1, par+1)
	}

	return recurse(arr, index+1, par-1)
}

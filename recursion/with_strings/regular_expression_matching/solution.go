package solution

// import "fmt"

func Solve(s, p string) bool {
	if len(s) == 0 || len(p) == 0 {
		return false
	}

	return recurse(s, p, 0, 0)
}

func recurse(s, p string, sPointer, pPointer int) bool {
	// fmt.Println(sPointer, pPointer, len(s), len(p))

	if pPointer == len(p) {
		if sPointer == len(s) {
			return true
		}

		return false
	}

	if sPointer == len(s) {
		if pPointer == len(p) {
			return true
		}

		return false
	}

	currPChar := p[pPointer]

	// Check if next char is a '*'
	if pPointer+1 < len(p) && p[pPointer+1] == '*' {

		for ; sPointer < len(s); sPointer++ {
			if !(currPChar == '.' || s[sPointer] == currPChar) {
				break
			}
		}

		return recurse(s, p, sPointer, pPointer+2)
	}

	if p[pPointer] == '.' || s[sPointer] == p[pPointer] {
		return recurse(s, p, sPointer+1, pPointer+1)
	}

	return false
}

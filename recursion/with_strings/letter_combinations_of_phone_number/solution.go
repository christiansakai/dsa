package solution

func Solve(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := &[]string{}
	set := &[]byte{}

	recurse(digits, 0, set, result)

	return *result
}

func recurse(
	digits string,
	digitIndex int,
	set *[]byte,
	result *[]string,
) {
	if digitIndex == len(digits) {
		*result = append(*result, string(*set))
		return
	}

	digit := digits[digitIndex]
	chars := dict[digit]

	for _, ch := range chars {
		*set = append(*set, ch)
		recurse(digits, digitIndex+1, set, result)
		*set = (*set)[:len(*set)-1]
	}
}

var dict map[byte][]byte = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

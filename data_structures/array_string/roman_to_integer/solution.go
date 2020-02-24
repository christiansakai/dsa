package solution

func Solve(s string) int {
	if len(s) == 0 {
		return 0
	}

	total := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			ch := s[i]
			nch := s[i+1]
			str := string([]byte{ch, nch})

			if _, ok := dict2[str]; ok {
				total += dict2[str]

				i += 1
			} else {
				ch := s[i]
				total += dict1[ch]
			}
		} else {
			ch := s[i]
			total += dict1[ch]
		}
	}

	return total
}

var dict1 = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var dict2 = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

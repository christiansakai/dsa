package solution

func TopDown(s string) int {
	if len(s) == 0 {
		return 0
	}

	cache := map[int]int{}
	return recurse(s, len(s)-1, cache)
}

func recurse(s string, index int, cache map[int]int) int {
	if index == -1 {
		return 1
	}

	if result, ok := cache[index]; ok {
		return result
	}

	total := 0
	str := string([]byte{s[index]})
	if _, ok := dict[str]; ok {
		total += recurse(s, index-1, cache)
	}

	if index-1 >= 0 {
		str := string([]byte{s[index-1], s[index]})
		if _, ok := dict[str]; ok {
			total += recurse(s, index-2, cache)
		}
	}

	cache[index] = total

	return total
}

var dict = map[string]bool{
	"1":  true,
	"2":  true,
	"3":  true,
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
	"11": true,
	"12": true,
	"13": true,
	"14": true,
	"15": true,
	"16": true,
	"17": true,
	"18": true,
	"19": true,
	"20": true,
	"21": true,
	"22": true,
	"23": true,
	"24": true,
	"25": true,
	"26": true,
}

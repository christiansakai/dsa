package solution

import "fmt"

func TopDown(s string, wordDict []string) bool {
	bs := []byte(s)
	bc := []byte{}

	return recurse(wordDict, bs, bc)
}

func recurse(wordDict []string, rem, curr []byte) bool {
	if len(rem) == 0 {
		return false
	}

	// Check subprob
	lastCh := rem[len(rem)-1]

	rem = rem[:len(rem)-1]
	curr = append([]byte{lastCh}, curr...)

	fmt.Println(string(rem), string(curr))

	result := false
	for _, word := range wordDict {
		if isSimilar(word, curr) {
			subProb := recurse(wordDict, rem, []byte{})
			fmt.Println("similar ", string(rem), string(curr), word)

			result = result || (true && subProb)
		} else {
			subProb := recurse(wordDict, rem, curr)
			result = result || subProb
		}
	}

	return result
}

func isSimilar(word string, data []byte) bool {
	if len(word) != len(data) {
		return false
	}

	for i := 0; i < len(word); i++ {
		if word[i] != data[i] {
			return false
		}
	}

	return true
}

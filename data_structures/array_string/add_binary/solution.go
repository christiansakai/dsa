package solution

import "fmt"

func Solve(s1 string, s2 string) string {
	b := []byte{}
	s1, s2 = padToSameLength(s1, s2)

	var carry byte = '0'

	for i := len(s1) - 1; i >= 0; i-- {
		currCarry, remainder := add(s1[i], s2[i], carry)
		b = append(b, remainder)

		carry = currCarry
	}

	fmt.Println(string(b))

	if carry == '1' {
		b = append(b, '1')
	}

	reverse(b)
	return string(b)
}

func padToSameLength(s1, s2 string) (string, string) {
	if len(s1) > len(s2) {
		paddedS2 := pad(s2, len(s1))
		return s1, paddedS2
	}

	paddedS1 := pad(s1, len(s2))
	return paddedS1, s2
}

func pad(s string, length int) string {
	b := []byte{}
	padCount := length - len(s)

	for i := 0; i < padCount; i++ {
		b = append(b, '0')
	}

	for _, c := range s {
		b = append(b, byte(c))
	}

	return string(b)
}

func add(c1, c2, c3 byte) (byte, byte) {
	arr := []byte{c1, c2, c3}
	oneCount := 0
	for _, e := range arr {
		if e == '1' {
			oneCount++
		}
	}

	switch oneCount {
	case 3:
		return '1', '1'
	case 2:
		return '1', '0'
	case 1:
		return '0', '1'
	default:
		return '0', '0'
	}
}

func reverse(b []byte) {
	mid := len(b) / 2
	for i := 0; i < mid; i++ {
		temp := b[i]
		b[i] = b[len(b)-1-i]
		b[len(b)-1-i] = temp
	}
}

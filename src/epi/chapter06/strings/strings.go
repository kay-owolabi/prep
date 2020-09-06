package strings

import "strings"

//6.1 Interconvert Strings And Integers
func IntToString(x int) string {
	res := strings.Builder{}
	var isNegative bool
	if x < 0 {
		isNegative = true
		x = -x
	}

	for ok := true; ok; x, ok = x/10, x > 0 {
		d := x % 10
		c := rune('0' + d)
		res.WriteRune(c)
	}

	if isNegative {
		res.WriteRune('-')
	}

	runes := []rune(res.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

//6.1 Interconvert Strings And Integers
func StringToInt(s string) int {
	var res int
	runes := []rune(s)
	var negative bool
	for i, r := range runes {
		if i == 0 && (r == '-' || r == '+') {
			negative = r == '-'
		} else {
			res = res*10 + int(r-'0')
		}
	}

	if negative {
		res = -res
	}

	return res
}

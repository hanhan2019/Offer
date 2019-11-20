package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(strToInt("-10"))
}
func strToInt(str string) int {
	sLen := len(str)
	if sLen < 1 {
		return 0
	}
	index := 0
	for i, v := range []byte(str) {
		if v == ' ' {
			continue
		}
		index = i
		break
	}
	str = str[index:]
	maybeType := str[0]
	if maybeType == '-' || maybeType == '+' {
		str = str[1:]
		if len(str) < 1 {
			return 0
		}
	}
	n := 0
	for _, v := range []byte(str) {
		if v == ' ' {
			continue
		}
		value := v - '0'
		if value > 9 {
			break
		}
		if n > math.MaxInt32/10 || (n == math.MaxInt32/10 && value > 7) {
			if maybeType == '-' {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		n = n*10 + int(value)
	}
	if maybeType == '-' {
		n = -n
	}
	return n
}

package main

import "fmt"

func main() {
	fmt.Println(reverseWords("I am a student."))
}
func reverseWords(s string) string {
	list := [][]byte{}
	a := []byte{}
	for i := 0; i < len(s); i++ {
		fmt.Println(list, a, string(s[i]))
		if s[i] != ' ' {
			a = append(a, s[i])
			fmt.Println(list, a)
		} else {
			list = append(list, a)
			a = a[0:0]
		}
	}
	fmt.Println(list)
	result := make([]byte, 0)
	for j := len(list) - 1; j >= 0; j-- {
		result = append(result, list[j]...)
		result = append(result, []byte(" ")...)
	}
	return string(result)
}

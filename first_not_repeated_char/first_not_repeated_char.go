package main

import "fmt"

func firstNotRepeatingChar(s string) byte {
	indexList := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if _, ok := indexList[s[i]]; ok {
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				indexList[s[j]] = true
			}
		}
		if _, ok := indexList[s[i]]; !ok {
			return s[i]
		}
	}
	return []byte("#")[0]
}

func main() {
	s := "abaccdeff"
	fmt.Println(string(firstNotRepeatingChar(s)))
}

//输入一个字符串,按字典序打印出该字符串中字符的所有排列。例如输入字符串abc,则打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
package main

import "fmt"

func main() {
	s := "abc"
	b := "ABC"
	d := "aab"
	fmt.Println(s, " Permutation: ", stringPermutations(s))
	fmt.Println(b, " Permutation: ", permutation(b))
	fmt.Println(d, " Permutation: ", permutation(d))
}
func stringPermutations(str string) []string {
	result := []string{}
	s := []byte(str)
	len := len(s)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len {
			str := string(s)
			result = append(result, str)
			return
		}
		m := make(map[byte]bool)
		for i := idx; i < len; i++ {
			if _, ok := m[s[i]]; ok {
				continue
			}
			m[s[i]] = true
			s[idx], s[i] = s[i], s[idx]
			dfs(idx + 1)
			s[i], s[idx] = s[idx], s[i]
		}
	}
	dfs(0)
	return result
}

//book
func permutation(str string) []string {
	if len(str) == 0 {
		return []string{}
	}
	result := []string{}
	return permutationOperation(str, 0, 0, result)
}

func permutationOperation(str string, head, begin int, result []string) []string {
	s := []byte(str)
	if head == len(s) {
		result = append(result, str)
	} else {
		for i := begin; i != len(s); i++ {
			s[i], s[begin] = s[begin], s[i]
			permutationOperation(str, 0, begin+1, result)
			s[i], s[begin] = s[begin], s[i]
		}
	}
	return result
}

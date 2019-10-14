//替换空格
//golang 只有值传递
//int、string、struct是非引用类型，无法改变原内容数据
//指针、map、slice、chan是引用类型，可以改变原内容数据
package main

import "fmt"

func main() {
	str := []byte{'W', 'e', ' ', 'A', 'r', 'e', ' ', 'H', 'a', 'p', 'p', 'y', 'x', 'x', 'x', 'x'}
	// s := string(str)
	// for _, v := range s {
	// 	fmt.Printf("%v", v)
	// }
	// fmt.Println("\n")
	// fmt.Printf("%v", []byte(s))
	ReplaceSpace(str, 11)
	fmt.Println(string(str))
}

func ReplaceSpace(s []byte, n int) {
	// s := []byte(str)
	count := 0
	for _, v := range s {
		if v == ' ' {
			count++
		}
	}
	finalLen := n + count*2
	for i := n; i >= 0; i-- {
		if s[i] == ' ' {
			s[finalLen] = '0'
			finalLen--
			s[finalLen] = '2'
			finalLen--
			s[finalLen] = '%'
			finalLen--
		} else {
			s[finalLen] = s[i]
			finalLen--
		}
	}
	// str = string(s)
}

func replaceSpace(str []byte, length int) {
	count := 0
	// 遍历一遍字符串， 统计字符出现的数目, 计算替换后的字符串长度
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			count++
		}
	}
	newlength := length + count*2

	// 两个index，一个指向length-1, 另一个指向newlength-1，遍历一遍字符串，完成替换
	for l, nl := length-1, newlength-1; l >= 0 && nl >= 0; {
		if str[l] == ' ' {
			str[nl] = '0'
			nl--
			str[nl] = '2'
			nl--
			str[nl] = '%'
			nl--
			l--
		} else {
			str[nl] = str[l]
			nl--
			l--
		}
	}
}

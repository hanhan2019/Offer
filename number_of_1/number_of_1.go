//输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。
package main

func main() {
}

func number_of_1(n int) int {
	count := 0
	for i := 1; i < 32; i++ {
		if n>>uint(i)&1 != 0 {
			count++
		}
	}
	return count
}

package main

import "fmt"

func main() {
	fmt.Println(findContinuousSequence(15))
}

func findContinuousSequence(sum int) [][]int {
	a := [][]int{}
	for i, j, s := 1, 1, 1; i <= sum/2+1; i++ {
		fmt.Println(i, j, s)
		for s < sum {
			j++
			s += j
		}
		fmt.Println(s)
		if s == sum && j > i {
			b := []int{}
			for k := i; k <= j; k++ {
				b = append(b, k)
			}
			a = append(a, b)
		}
		s -= i
	}
	return a
}

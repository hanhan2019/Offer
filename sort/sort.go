package main

import "fmt"

func main() {
	a := []int{6, 3, 2, 7, 4, 8}
	fmt.Println(a)
	quickSort(&a)
	fmt.Println(a)
}

//冒泡
func bubbleSort(s []int) []int {
	n := len(s)
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func quickSortPart(s *[]int, low, high int) int {
	key := (*s)[low]
	for low != high {
		for low < high && (*s)[high] >= key {
			high--
		}
		(*s)[low] = (*s)[high]
		for low < high && (*s)[low] <= key {
			low++
		}
		(*s)[high] = (*s)[low]
	}
	(*s)[low] = key
	return low
}

func quickSort(s *[]int) {
	if len(*s) != 0 {
		mid := quickSortPart(s, 0, len(*s)-1)
		fmt.Println((*s))
		fmt.Println((*s)[0:mid])
		quickSortPart(s, 0, mid)
		fmt.Println((*s))
		fmt.Println((*s)[mid+1 : len(*s)])
		quickSortPart(s, mid+1, len(*s)-1)
	}
}

package main

func main() {

}
func multiply(A []int) []int {
	if len(A) == 0 {
		return nil
	}
	B := []int{1}
	for i := 1; i < len(A); i++ {
		B = append(B, B[i-1]*A[i-1])
	}
	temp := 1
	for j := len(A) - 2; j >= 0; j-- {
		temp *= A[j+1]
		B[j] *= temp
	}
	return B
}

//输入：[1, 2, 3, 4, 5]

//输出：[120, 60, 40, 30, 24]

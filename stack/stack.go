package main

func main() {

}

func isPopOrder(pushV []int, popV []int) bool {
	n := len(pushV)
	if n == 0 && len(popV) == 0 {
		return true
	}
	if n != len(popV) {
		return false
	}
	temp := []int{}
	temp = append(temp, pushV[0])
	p, q := 0, 0
	for p < n {
		for len(temp) != 0 && temp[len(temp)-1] == popV[q] {
			temp = temp[:len(temp)-1]
			q++
		}
		p++
		if p < n {
			temp = append(temp, pushV[p])
		}
	}
	if len(temp) == 0 {
		return true
	}
	return false
}

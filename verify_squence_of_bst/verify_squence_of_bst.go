package main

func main() {

}

func verifySequenceOfBST(sequence []int) bool {
	if len(sequence) <= 1 {
		return true
	}
	c := sequence[len(sequence)-1]
	k := 0
	for i, v := range sequence {
		if v >= c {
			k = i
			break
		}
	}
	a := sequence[:k]
	b := sequence[k : len(sequence)-1]
	for _, v := range b {
		if v < c {
			return false
		}
	}
	return verifySequenceOfBST(a) && verifySequenceOfBST(b)
}

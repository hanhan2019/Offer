package main

func main() {

}

func reorder_array(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		if array[i]%2 != 0 {
			continue
		}
		for j := i + 1; j < len(array); j++ {
			if array[j]%2 != 0 {
				array[i], array[j] = array[j], array[i]
				break
			}
		}
	}
	return array
}

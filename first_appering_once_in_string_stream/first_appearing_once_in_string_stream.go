package main

func main() {

}

var mapList map[byte]int = make(map[byte]int)
var queue []byte = []byte{}

func insert(ch byte) {
	if _, ok := mapList[ch]; ok {
		mapList[ch] += 1
	} else {
		mapList[ch] = 1
	}
	if mapList[ch] > 1 {
		for len(queue) != 0 && mapList[queue[0]] > 1 {
			queue = queue[1:len(queue)]
		}
	} else {
		queue = append(queue, ch)
	}
}
func firstAppearingOnce() byte {
	if len(queue) == 0 {
		return []byte("#")[0]
	}
	return queue[0]
}

package main

func main() {

}

type MinStack struct {
	List []int
	Min  []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	this.List = append(this.List, x)
	if len(this.Min) == 0 {
		this.Min = append(this.Min, x)
	} else {
		if this.Min[len(this.Min)-1] >= x {
			this.Min = append(this.Min, x)
		}
	}
}

func (this *MinStack) Pop() {
	value := this.List[len(this.List)-1]
	this.List = this.List[:len(this.List)-1]
	if value == this.Min[len(this.Min)-1] {
		this.Min = this.Min[:len(this.Min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.List[len(this.List)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

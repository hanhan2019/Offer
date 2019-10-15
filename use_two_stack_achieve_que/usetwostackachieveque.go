//用两个栈实现队列
package main

func main() {

}

type Stack struct {
	V []int
}

func (this *Stack) Push(x int) {
	this.V = append(this.V, x)
}

func (this *Stack) Pop() int {
	value := this.Peek()
	this.V = this.V[:len(this.V)-1]
	return value
}

func (this *Stack) Peek() int {
	return this.V[len(this.V)-1]
}

func (this *Stack) Empty() bool {
	return len(this.V) == 0
}

type MyQueue struct {
	S1 Stack
	S2 Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		S1: Stack{},
		S2: Stack{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.S1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.Move()
	return this.S2.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.Move()
	return this.S2.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.S1.Empty() && this.S2.Empty() {
		return true
	}
	return false
}

func (this *MyQueue) Move() {
	if this.S2.Empty() && !this.S1.Empty() {
		for !this.S1.Empty() {
			this.S2.Push(this.S1.Pop())
		}
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

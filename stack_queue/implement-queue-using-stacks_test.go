package stack_queue

type MyQueue struct {
	input  []int
	output []int
}

/** Initialize your data structure here. */
//func Constructor() MyQueue {
//	return MyQueue{}
//}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	//if len(this.output) == 0 {
	//	for _, item := range this.input {
	//		this.output = append(this.output, item)
	//		this.input = this.input[:len(this.input) - 1]
	//	}
	//}
	value := this.Peek()
	this.output = this.output[:len(this.output)-1]
	return value
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			this.output = append(this.output, this.input[len(this.input)-1])
			this.input = this.input[:len(this.input)-1]
		}
	}
	return this.output[len(this.output)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.input) == 0 && len(this.output) == 0
}

//func TestQueue(t *testing.T) {
//	obj := Constructor()
//	obj.Push(1)
//	obj.Push(2)
//	fmt.Println(obj.Pop())
//	fmt.Println(obj.Pop())
//	fmt.Println(obj.Empty())
//}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

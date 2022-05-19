package stack

type MyQueue struct {
	m []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.m = append(this.m, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	p := this.m[0]
	this.m = this.m[1:]
	return p
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.m[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.m) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

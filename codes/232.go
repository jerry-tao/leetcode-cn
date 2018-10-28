type MyQueue struct {
    D []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.D = append(this.D,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    x := this.D[0]
    this.D = this.D[1:]
    return x
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.D[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.D)==0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

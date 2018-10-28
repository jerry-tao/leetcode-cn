type MyStack struct {
    Data []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.Data = append(this.Data, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    x := this.Data[len(this.Data)-1]
    this.Data =  this.Data[:len(this.Data)-1]
    return x
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.Data[len(this.Data)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.Data) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

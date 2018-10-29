type MinStack struct {
    D []int
    min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        D: []int{},
        min: []int{},
    }
}


func (this *MinStack) Push(x int)  {
    if len(this.min)==0 || x<=this.min[len(this.min)-1]{
        this.min  = append(this.min,x)
    }
    this.D = append(this.D,x)
}


func (this *MinStack) Pop()  {
    if this.D[len(this.D)-1] == this.min[len(this.min)-1]{
        this.min = this.min[:len(this.min)-1]
    }
    this.D = this.D[:len(this.D)-1]
}


func (this *MinStack) Top() int {
    return this.D[len(this.D)-1]
}


func (this *MinStack) GetMin() int {
    return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

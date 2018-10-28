type MyCalendar struct {
    D [][]int
}


func Constructor() MyCalendar {
   return MyCalendar{}
}


func (this *MyCalendar) Book(start int, end int) bool {
    for _,v:=range this.D{
        if (end > v[0] && end<v[1]) || (start>=v[0] && start<v[1] || (start<=v[0] && end>=v[1])) {
            return false
        }
    }
    this.D = append(this.D, []int{start,end})
    return true
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

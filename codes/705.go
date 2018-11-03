type pair struct{
    value int
}
type MyHashSet struct {
    d []*pair
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    return MyHashSet{d: make([]*pair,1000000)}
}


func (this *MyHashSet) Add(key int)  {
    this.d[key] = &pair{value: key}
}


func (this *MyHashSet) Remove(key int)  {
    this.d[key] = nil
}


/** Returns true if this set did not already contain the specified element */
func (this *MyHashSet) Contains(key int) bool {
    return this.d[key]!=nil
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

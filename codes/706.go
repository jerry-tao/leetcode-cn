// cheat a little
type pair struct{
    Key int
    Value int
}
type MyHashMap struct {
    D []*pair
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    // use slice as hash
    return MyHashMap{
        D: make([]*pair,1000000),
    }
}


/** value will always be positive. */
func (this *MyHashMap) Put(key int, value int)  {
    this.D[key] = &pair{Key:key,Value:value}
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    if this.D[key] == nil{
        return -1
    }
    return this.D[key].Value
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    this.D[key] = nil
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

type LRUCache struct {
    Capacity int
    Size  int
    Pairs map[int]*Node
    Head *Node
    Tail *Node
}

type Node struct{
    Val int
    Key int
    Pre *Node
    Next *Node
}


func Constructor(capacity int) LRUCache {
    head:= &Node{}
    tail:=&Node{Pre:head}
    head.Next = tail
    return LRUCache{
        Capacity: capacity,
        Pairs: make(map[int]*Node),
        Size:0,
        Head: head,
        Tail: tail,
    }
}

func (this *LRUCache) moveToHead(n *Node){

    if this.Head.Next != n{
        next := n.Next
        pre := n.Pre
        pre.Next = next
        next.Pre = pre
        n.Pre = this.Head
        n.Next = this.Head.Next
        this.Head.Next.Pre = n
        this.Head.Next = n
    }

}
func (this *LRUCache) addToHead(n *Node){
        n.Pre = this.Head
        n.Next = this.Head.Next
        this.Head.Next.Pre=n
        this.Head.Next = n

}
func (this *LRUCache) removeTail(){
    delete(this.Pairs, this.Tail.Pre.Key)

    nt:=this.Tail.Pre.Pre
    this.Tail.Pre = nt
    nt.Next=this.Tail
}

func (this *LRUCache) Get(key int) int {

    node:=this.Pairs[key]
    if node!=nil{
        this.moveToHead(node)

        return node.Val
    }

    return -1


}


func (this *LRUCache) Put(key int, value int)  {
    if n,ok:=this.Pairs[key];ok{
        n.Val=value
        this.moveToHead(n)
        return
    }
    node:=&Node{Val:value,Key:key}
    if this.Size==this.Capacity{
        this.removeTail()
    }else{
        this.Size++
    }

    this.addToHead(node)
    this.Pairs[key]=node

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

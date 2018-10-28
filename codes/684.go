func findRedundantConnection(edges [][]int) []int {
    ids:=make([]int,len(edges)+1)
    for i,_:=range ids{
        ids[i] = i
    }
    find:=func(p int)int{
        for p!=ids[p]{p=ids[p]}
        return p
    }
    _connected:=func(p []int)bool{
        if ids[p[0]] == ids[p[1]] {return true}
        return find(p[0])==find(p[1])
    }

    for _,v:=range edges{
        if _connected(v){
            return v
        }
        p := find(v[0])
        q := find(v[1])
        ids[p]=q
    }
    return nil
}

func judgeCircle(moves string) bool {
    if len(moves)%2!=0{
        return false
    }
    if strings.Count(moves,"U") != strings.Count(moves,"D"){
        return false
    }
        if strings.Count(moves,"L") != strings.Count(moves,"R"){
        return false
    }
    return true
}

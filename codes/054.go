func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0{
        return nil
    }
    a,b:=0,0
    minB,minA:=0,0
    maxA := len(matrix[0])
    maxB := len(matrix)
    count := maxA*maxB
    maxA=maxA-1
    maxB=maxB-1
    arraw := 0 // 0 right，1 down，2 left，3 up 四个方向
    res := make([]int,count)

    for i:=0;i<count;i++{
        res[i] = matrix[b][a]

        if arraw==0{
            if a==maxA{
                minB++
                b++
                arraw = 1
            }else{
                a++
            }
            continue
        }
        if arraw==1{
            if b==maxB{
                maxA--
                a--
                arraw = 2
            }else{
                b++
            }
            continue
        }
        if arraw==2{
            if a == minA{
                minA++
                b--
                arraw=3
            }else{
                a--
            }
            continue
        }
        if arraw==3{
            if b==minB{
                maxB--
                a++
                arraw=0
            }else{
                b--
            }
            continue
        }
    }
    return res
}

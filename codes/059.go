func generateMatrix(n int) [][]int {
    res:=make([][]int,n)
    for i,_:=range res{
        res[i] = make([]int,n)
    }
    a,b:=0,0
    minB,minA:=0,0
    maxA := n-1
    maxB := n-1
    count := n*n
    arraw := 0 // 0 right，1 down，2 left，3 up 四个方向
     for i:=1;i<=count;i++{
         res[b][a] = i
           if arraw==0{
            if a==maxA{
                minB++
                b++
                arraw = 1
            }else{
                a++
            }
        }else if arraw==1{
            if b==maxB{
                maxA--
                a--
                arraw = 2
            }else{
                b++
            }
        }else if arraw==2{
            if a == minA{
                minA++
                b--
                arraw=3
            }else{
                a--
            }
        }else if arraw==3{
            if b==minB{
                maxB--
                a++
                arraw=0
            }else{
                b--
            }
        }

     }
    return res
}

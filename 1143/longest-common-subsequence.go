func longestCommonSubsequence(text1 string, text2 string) int {
    if len(text1) > len(text2){
        text1,text2 = text2,text1
    }
    prev := make([]int, len(text1)+1)
    curr := make([]int, len(text1)+1)

    for i:=1;i<len(text2)+1;i++{
        letter1 := text2[i-1]
        for j:=1;j<len(text1)+1;j++{
            letter2 := text1[j-1]
            if letter1 == letter2{
                curr[j] = prev[j-1]+1
            }else{
                curr[j] = max(prev[j],curr[j-1])
            } 
        }
        curr,prev = prev,curr
    }
    return prev[len(text1)]
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

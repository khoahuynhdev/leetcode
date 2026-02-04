package solution

func firstUniqChar(s string) int {
    chars,dct := make([]int, 26), make(map[int]int)
    var ans int = 1e6
    for i, c := range s {
        idx := int(c - 'a')
        chars[idx]++
        dct[idx]=i
    }
    for i:=0;i<len(chars);i++{
        if chars[i]== 1 {
            if  idx:=dct[i];idx < ans {
                ans = idx
            }
        }
    }
    if ans == 1e6 {
        return -1 
    }
    return ans
}

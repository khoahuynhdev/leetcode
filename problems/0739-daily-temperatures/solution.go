package solution


func dailyTemperatures(temperatures []int) []int {
    ans,dct := make([]int, len(temperatures)), make([]int,0)
    for i:=0;i<len(temperatures);i++ {
        for len(dct) > 0 && temperatures[dct[len(dct) -1 ]] < temperatures[i] {
            idx := dct[len(dct) - 1]
            dct = dct[:len(dct) - 1]
            ans[idx] = i - idx
        }
        dct = append(dct, i)
    }
    return ans
}

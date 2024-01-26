const MOD int = 1000 * 1000 * 1000 + 7

func calculate(m, n, maxMove, cr, cl int, mem map[[3]int]int) (res int) {
    if cr < 0 || cl < 0 || cr == m || cl == n {return 1}
    if res, ok := mem[[3]int{cr,cl,maxMove}]; ok {
        return res
    }
    if maxMove == 0 {return 0}
    clock := [][2]int{[2]int{0,1}, [2]int{0,-1}, [2]int{1,0}, [2]int{-1,0}}
    for _, c := range clock {
        res = (res + calculate(m, n, maxMove-1, cr+c[0], cl+c[1], mem)) % MOD
    }
    mem[[3]int{cr,cl,maxMove}] = res
    return res
}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    return calculate(m, n, maxMove, startRow, startColumn, map[[3]int]int{})
}

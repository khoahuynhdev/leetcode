func kInversePairs(n int, k int) int {
    prevDP := make([]int, k+1)
    dp := make([]int, k+1)
    dp[0] = 1
    
    for i := 2; i <= n; i++ {
        prevDP, dp = dp, prevDP
        sum := 0
        
        for j := range dp {
            sum += prevDP[j]

            if j >= i {
               sum -= prevDP[j-i] 
            }
            
            dp[j] = sum % 1_000_000_007
        }
    }
    
    return dp[k]
}

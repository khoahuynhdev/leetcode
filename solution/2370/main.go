package main

func longestIdealString(s string, k int) int {
    const maxChars = 26
    dp := make([]int, maxChars)
    result := 0

    for _, ch := range s {
        maxLen := 0
        // Compute the max dp value from possible predecessors within range defined by k
        for d := 0; d <= k; d++ {
            if int(ch)-d >= 'a' {
                maxLen = max(maxLen, dp[int(ch)-'a'-d])
            }
            if int(ch)+d <= 'z' {
                maxLen = max(maxLen, dp[int(ch)-'a'+d])
            }
        }
        // Update dp value for current character
        dp[int(ch)-'a'] = maxLen + 1
        // Update result to keep track of the maximum length encountered so far
        result = max(result, dp[int(ch)-'a'])
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

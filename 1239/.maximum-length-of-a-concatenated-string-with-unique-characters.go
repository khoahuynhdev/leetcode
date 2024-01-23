package solution

func maxSubStr(currLen int, usedLetters []int, arr []string) int {
    if len(arr) == 0 {
        return currLen
    }

    maxLen := currLen
    for i, s := range arr {
        canJoin := true

        for _, c := range s {
            pos := int(c) - 97

            if usedLetters[pos] != 0 {
                canJoin = false
            }

            usedLetters[pos]++
        }

        if canJoin {
            sub := maxSubStr(currLen + len(s), usedLetters, arr[i+1:])
            maxLen = max(maxLen, sub)
        }

        for _, c := range s {
            pos := int(c) - 97
            usedLetters[pos]--
        }
    }

    return maxLen
}

func maxLength(arr []string) int {
    usedLetters := make([]int, 26)
    return maxSubStr(0, usedLetters, arr)
}

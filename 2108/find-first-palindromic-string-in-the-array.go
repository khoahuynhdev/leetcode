package solution

func firstPalindrome(words []string) string {
    for _,str := range words {
        if IsPalindrome(str) { return str}
    }
   return ""
}

func IsPalindrome(str string) bool {
    for i:=0;i<len(str) / 2;i++ {
            if str[i] != str[len(str) - i - 1] { return false }
        }
    return true
}

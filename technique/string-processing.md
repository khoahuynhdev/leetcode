# String Processing in Go

## Fundamentals

Go strings are immutable byte slices (`[]byte`) under the hood. When you iterate with `range`, you get runes (Unicode code points), not bytes. For ASCII-only LeetCode problems, `s[i]` (byte access) is fine and faster than rune iteration.

```go
s := "hello"
s[0]          // byte 'h'
len(s)        // byte length (same as char length for ASCII)
string(s[1])  // "e"
```

## Building Strings

Since strings are immutable, concatenation in a loop is O(n^2). Use `strings.Builder` or `[]byte` instead.

```go
// strings.Builder - best for building strings incrementally
var sb strings.Builder
for _, ch := range s {
    sb.WriteRune(ch)     // append a rune
    sb.WriteByte('a')    // append a byte (ASCII)
    sb.WriteString("ab") // append a string
}
result := sb.String()

// []byte - useful when you need index-based mutation
buf := []byte(s)
buf[0] = 'H'
result := string(buf)

// []rune - when you need to modify characters by index (Unicode-safe)
runes := []rune(s)
runes[0] = 'H'
result := string(runes)
```

## Common `strings` Package Functions

```go
import "strings"

strings.Contains(s, "sub")         // substring check
strings.HasPrefix(s, "pre")        // starts with
strings.HasSuffix(s, "suf")        // ends with
strings.Index(s, "sub")            // first index of substring, -1 if not found
strings.LastIndex(s, "sub")        // last index of substring
strings.Count(s, "a")              // count non-overlapping occurrences
strings.Repeat("ab", 3)            // "ababab"
strings.Replace(s, "old", "new", n) // replace first n occurrences (-1 for all)
strings.ReplaceAll(s, "old", "new") // replace all
strings.ToLower(s)
strings.ToUpper(s)
strings.TrimSpace(s)               // trim leading/trailing whitespace
strings.Trim(s, "chars")           // trim specific chars from both ends
strings.Split(s, ",")              // split into []string
strings.Join([]string{"a","b"}, ",") // "a,b"
strings.Map(func(r rune) rune { ... }, s) // transform each rune
```

## Character / Byte Operations

```go
// Check character type (ASCII)
ch := s[i]
ch >= 'a' && ch <= 'z'   // lowercase letter
ch >= 'A' && ch <= 'Z'   // uppercase letter
ch >= '0' && ch <= '9'   // digit

// Convert
ch - 'a'                 // letter to 0-25 index (for frequency arrays)
ch - '0'                 // digit char to int
byte('a' + 3)            // index to letter: 'd'
ch ^ 32                  // toggle case (ASCII trick: 'a'^32='A', 'A'^32='a')
ch | 32                  // to lowercase
ch & ^byte(32)           // to uppercase

// strconv for number conversion
import "strconv"
strconv.Itoa(42)         // int to string: "42"
strconv.Atoi("42")       // string to int (returns int, error)
```

## Frequency Counting

This comes up constantly in string problems (anagrams, permutations, character counts).

```go
// Fixed array for lowercase letters - faster than map
var freq [26]int
for _, ch := range s {
    freq[ch-'a']++
}

// Map for arbitrary characters
freq := make(map[byte]int)
for i := 0; i < len(s); i++ {
    freq[s[i]]++
}
```

## Substrings and Slicing

```go
s[i:j]    // substring from index i to j-1 (like Python s[i:j])
s[i:]     // from i to end
s[:j]     // from start to j-1

// Comparing substrings
s[i:j] == t[k:l]  // direct comparison works
```

## Sliding Window on Strings

A very common pattern for substring problems.

```go
// Fixed-size window
for i := 0; i+k <= len(s); i++ {
    window := s[i : i+k]
    // process window
}

// Variable-size window (two pointers)
left := 0
freq := make(map[byte]int)
for right := 0; right < len(s); right++ {
    freq[s[right]]++
    for /* window invalid */ {
        freq[s[left]]--
        left++
    }
    // process valid window: s[left:right+1]
}
```

## Two Pointers on Strings

Used for palindrome checks, reversal, and partitioning.

```go
// Palindrome check
func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}

// Reverse a string
func reverse(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}
```

## String DP Patterns

Common setups for DP problems involving strings.

```go
// 1D DP on a single string (e.g., decode ways, word break)
dp := make([]int, len(s)+1)
dp[0] = 1 // base case
for i := 1; i <= len(s); i++ {
    // transition using s[i-1] or substrings ending at i
}

// 2D DP on two strings (e.g., edit distance, LCS)
dp := make([][]int, len(s)+1)
for i := range dp {
    dp[i] = make([]int, len(t)+1)
}
// fill base cases, then dp[i][j] uses dp[i-1][j], dp[i][j-1], dp[i-1][j-1]

// 2D DP on substrings of one string (e.g., palindrome partitioning)
// dp[i][j] = something about s[i..j]
n := len(s)
dp := make([][]bool, n)
for i := range dp {
    dp[i] = make([]bool, n)
    dp[i][i] = true
}
// expand by length
```

## Stack-Based String Problems

Used for matching brackets, decoding strings, and removing characters.

```go
// Balanced parentheses / bracket matching
stack := []byte{}
for i := 0; i < len(s); i++ {
    if s[i] == '(' {
        stack = append(stack, s[i])
    } else if s[i] == ')' {
        if len(stack) == 0 {
            // unmatched
        }
        stack = stack[:len(stack)-1]
    }
}

// Build result string with deletions (e.g., remove adjacent duplicates)
stack := []byte{}
for i := 0; i < len(s); i++ {
    if len(stack) > 0 && stack[len(stack)-1] == s[i] {
        stack = stack[:len(stack)-1]
    } else {
        stack = append(stack, s[i])
    }
}
result := string(stack)
```

## Quick Reference: LeetCode String Patterns

| Pattern | When to Use | Key Technique |
|---|---|---|
| Frequency array `[26]int` | Anagrams, char counts | `ch - 'a'` indexing |
| Sliding window | Substring with constraint | Two pointers + map/array |
| Two pointers | Palindromes, reversal | `l, r := 0, len(s)-1` |
| Stack | Brackets, decoding, deletions | `[]byte` as stack |
| DP 1D | Single string, sequential decisions | `dp[i]` = answer for `s[:i]` |
| DP 2D | Two strings or substring ranges | `dp[i][j]` |
| Trie | Prefix matching, word search | Struct with `[26]*Node` |
| String hashing | Pattern matching, dedup substrings | Rolling hash |

# 2911. Minimum Changes to Make K Semi-palindromes

[LeetCode Link](https://leetcode.com/problems/minimum-changes-to-make-k-semi-palindromes/)

Difficulty: Hard
Topics: Two Pointers, String, Dynamic Programming
Acceptance Rate: 61.4%

## Hints

### Hint 1

This problem has two separate parts. One is "split the string into exactly `k` pieces and minimize the total". The other is "how much does one piece cost?". Try solving them separately. The first part is a well-known DP pattern.

### Hint 2

If you knew `cost[i][j]`, the minimum changes needed to turn `s[i..j]` into a semi-palindrome, the rest is a standard partition DP: `dp[p][j] = min over i (dp[p-1][i] + cost[i][j-1])`. With `n <= 200`, you can afford to precompute `cost` for all O(n²) substrings, as long as each one is cheap.

### Hint 3

For a fixed divisor `d`, the groups are just the positions `g, g+d, g+2d, ...` inside the substring. The cheapest way to make one group a palindrome is to count its mismatched mirror pairs (two pointers from both ends). Each mismatch takes exactly one change. Add these counts over all `d` groups, and take the minimum over all proper divisors `d` of the length. One divisor check touches each character once, so it costs O(L). Also note that a substring of length 1 has no valid divisor, so every piece must have length ≥ 2.

## Approach

**Step 1: Precompute proper divisors.** For each length `L` from 2 to `n`, list every `d` with `1 <= d < L` and `L % d == 0`. A sieve-style loop does this: for each `d`, add it to `2d, 3d, ...`.

**Step 2: Cost of every substring.** For every substring `s[i..j]` with length `L = j - i + 1 >= 2`:
- For each proper divisor `d` of `L`, there are `d` groups, each with `m = L / d` characters. Group `g` consists of `s[i+g], s[i+g+d], ..., s[i+g+(m-1)d]`.
- Walk two pointers `lo = 0`, `hi = m-1` toward each other. Each pair with `s[i+g+lo*d] != s[i+g+hi*d]` needs one change.
- The cost for `d` is the total number of mismatches across all groups. `cost[i][j]` is the minimum over all `d`.

Why do mismatched pairs count exactly? In a palindrome, each mirror pair must match, and the pairs don't overlap. Changing one character of a mismatched pair fixes it, and nothing cheaper can fix it.

**Step 3: Partition DP.** Let `dp[p][j]` be the minimum changes needed to split the prefix `s[0..j)` into `p` valid pieces. Then:
- `dp[0][0] = 0`, and every other starting state is infinity.
- `dp[p][j] = min over i in [2(p-1), j-2] of dp[p-1][i] + cost[i][j-1]`.

The answer is `dp[k][n]`. Since `k <= n/2`, a valid split always exists. Only the previous row is needed, so the DP uses a rolling array.

**Example:** `s = "abcac"`, `k = 2`.
- `cost["ab"] = 1` (d=1, mismatch a/b).
- `cost["cac"] = 0` (d=1, already a palindrome).
- `dp[2][5] = dp[1][2] + cost[2][4] = 1 + 0 = 1`, and no other split does better.

## Complexity Analysis

Time Complexity: O(n³ · D + k · n²), where D is the maximum number of divisors of any length ≤ n (at most 12 for n ≤ 200). There are O(n²) substrings, each checked against D divisors at O(L) per check. In practice this is much faster than the bound.
Space Complexity: O(n²) for the cost table. The DP itself uses O(n) with rolling rows.

## Edge Cases

- **Length-1 pieces are invalid:** A single character has no proper divisor, so it can never be a semi-palindrome. The DP must force every piece to have length ≥ 2. If you treat a length-1 piece as cost 0, you will get wrong answers.
- **`k = n/2` with even `n`:** Every piece must be exactly length 2, so the answer is just the number of unequal adjacent pairs at positions (0,1), (2,3), and so on.
- **`k = 1`:** The whole string is one piece. This tests that all divisors are tried, not only `d = 1`. For example, `"abab"` costs 0 with `d = 2`, but would cost 2 with `d = 1`.
- **Prime lengths:** The only proper divisor is 1, so the substring must be an ordinary palindrome.
- **Strings that are already valid:** Strings like `"aabbaa"` or a string of all the same letter should return 0.

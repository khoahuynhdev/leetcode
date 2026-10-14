# 3723. Maximize Sum of Squares of Digits

[LeetCode Link](https://leetcode.com/problems/maximize-sum-of-squares-of-digits/)

Difficulty: Medium
Topics: Math, Greedy
Acceptance Rate: 65.7%

## Hints

### Hint 1

Don't try to enumerate good integers, because there are far too many. Ask yourself: if you have a fixed "budget" (the digit sum) to split across `num` slots, what kind of split makes the sum of squares as large as possible?

### Hint 2

Compare `5² + 5² = 50` with `9² + 1² = 82`, or `3² + 3² = 18` with `6² + 0² = 36`. Squaring rewards concentration. Moving one unit from a smaller digit to a larger digit always increases the score. Think greedily about which digits to make big.

### Hint 3

Make every digit as large as possible, from left to right: put `9` in each position while the remaining sum is at least 9, then one digit with whatever is left, then zeros. This maximizes the score *and* gives the largest number among the optimal answers. The only time no answer exists is when `sum > 9 * num`.

## Approach

**Feasibility.** Each digit is at most 9, so the largest possible digit sum is `9 * num`. If `sum > 9 * num`, return `""`. Since `sum >= 1`, any feasible answer can start with a non-zero digit, so leading zeros are never an issue.

**Why "as many 9s as possible" maximizes the score.** Take two digits `a >= b` with `b > 0` and `a < 9`. Replace them with `a + 1` and `b - 1`. The digit sum stays the same, and the score changes by

```
(a+1)² + (b-1)² - a² - b² = 2(a - b) + 2 > 0
```

So any arrangement with two digits that are both strictly between 0 and 9 can be improved. In an optimal answer, at most one digit lies strictly between 0 and 9. Every other digit is 9 or 0. That pins the multiset of digits down exactly: `q = sum / 9` nines, one digit `r = sum % 9` (if non-zero), and the rest zeros.

**Why this is also the largest number.** All optimal answers use that same multiset of digits. Among numbers with the same digits, the largest one has them in non-increasing order: 9s first, then `r`, then 0s.

**Algorithm.** Walk the `num` positions left to right. At each position, write `min(remaining, 9)` and subtract it from `remaining`.

**Example:** `num = 5, sum = 20`
- position 0: write 9, remaining 11
- position 1: write 9, remaining 2
- position 2: write 2, remaining 0
- positions 3–4: write 0

Result: `"99200"`, score `81 + 81 + 4 = 166`.

## Complexity Analysis

Time Complexity: O(num), since we write each digit exactly once.
Space Complexity: O(num) for the output string.

## Edge Cases

- **`sum > 9 * num`**: there is no valid integer, so return the empty string (Example 3: `num = 1, sum = 10`).
- **`sum == 9 * num`**: every digit is 9. Check that the boundary is inclusive.
- **`sum` divisible by 9 with digits left over**: there is no partial digit, so the 9s are followed straight by zeros (e.g. `num = 3, sum = 9` → `"900"`).
- **`sum = 1` with many digits**: the answer is `"1000…0"`. This confirms the leading digit is never zero.
- **Large `num` (up to 2·10⁵)**: build the result in a byte slice instead of concatenating strings in a loop, which would be O(num²).

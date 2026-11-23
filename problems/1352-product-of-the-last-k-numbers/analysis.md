# 1352. Product of the Last K Numbers

[LeetCode Link](https://leetcode.com/problems/product-of-the-last-k-numbers/)

Difficulty: Medium
Topics: Array, Math, Design, Data Stream, Prefix Sum
Acceptance Rate: 64.4%

## Hints

### Hint 1

If the question were "sum of the last k numbers," how would you answer it in O(1)? Think about what you could precompute as each number arrives.

### Hint 2

Prefix sums have a multiplicative cousin: prefix products. The product of a range is `prefix[n] / prefix[n-k]`. Division works for products as long as nothing in the way is... what?

### Hint 3

Zeros break division. But a zero also makes *every* product that includes it equal to 0. So whenever you see a zero, throw away the whole prefix list and start over. If `k` reaches back past the most recent reset (i.e. `k` is at least the current list length), the window contains a zero and the answer is 0.

## Approach

Maintain a slice `prefix` where `prefix[i]` is the product of all numbers added since the last zero, with a sentinel `prefix[0] = 1`.

- **add(num)**:
  - If `num == 0`, reset `prefix` to `[1]`. Any window that includes this zero has product 0, and no window that ends later can "see through" it without including it.
  - Otherwise, append `prefix[last] * num`.
- **getProduct(k)**:
  - Let `n = len(prefix) - 1` be the number of non-zero values since the last zero.
  - If `k > n`, the window of the last `k` numbers reaches the zero, so return 0.
  - Otherwise, return `prefix[n] / prefix[n-k]`. This is exact integer division because `prefix[n]` is literally `prefix[n-k]` multiplied by the last `k` numbers.

Example: add 3, 0, 2, 5, 4

- after 3: `[1, 3]`
- after 0: `[1]` (reset)
- after 2, 5, 4: `[1, 2, 10, 40]`
- `getProduct(2)` → `40 / 2 = 20`
- `getProduct(3)` → `40 / 1 = 40`
- `getProduct(4)` → `4 > 3`, so 0
- add 8: `[1, 2, 10, 40, 320]`; `getProduct(2)` → `320 / 10 = 32`

Why it doesn't overflow: the problem guarantees any contiguous product fits in 32 bits, and since prefixes reset at every zero, each stored prefix is itself a contiguous product of non-zero values. We use `int` (64-bit in Go) for extra safety.

## Complexity Analysis

Time Complexity: O(1) for both `Add` and `GetProduct` (amortized O(1) for `Add` due to slice growth).
Space Complexity: O(n), where n is the number of values added since the last zero.

## Edge Cases

- **Zero in the stream**: the reason for the reset; without it, division by a zero prefix (or zeroed-out prefixes) breaks.
- **k exactly equal to the count since the last zero**: should return the full prefix product, not 0. Watch for an off-by-one in the `k > n` comparison.
- **k reaching exactly the zero**: e.g. stream `[0, 2]`, `k = 2` must return 0.
- **Zero as the most recent element**: any `k >= 1` returns 0.
- **Ones in the stream**: prefix stays the same; division still works.
- **No zeros at all**: `k` equal to the whole stream length uses the sentinel `prefix[0] = 1`.

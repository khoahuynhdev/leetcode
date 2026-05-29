# 3300. Minimum Element After Replacement With Digit Sum

[LeetCode Link](https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/)

Difficulty: Easy
Topics: Array, Math
Acceptance Rate: 87.1%

## Hints

### Hint 1

Read the problem carefully. You only need two things: a way to transform each element, and a way to track the smallest transformed value. No fancy data structure is needed here — think about what a single linear pass through the array can accomplish.

### Hint 2

The transformation is "replace a number with the sum of its decimal digits." How do you extract digits from an integer in any language? Repeatedly take the value modulo 10 (the last digit) and divide by 10 (drop the last digit) until the value reaches 0.

### Hint 3

You do not need to actually rewrite the array. Compute the digit sum for each element on the fly, compare it against the running minimum, and update the minimum if smaller. Initialize the minimum to a safely large value (or to the digit sum of the first element) so the very first comparison works correctly.

## Approach

Walk the array once. For each element `n`, compute its digit sum by looping:

1. Set `sum = 0`.
2. While `n > 0`, add `n % 10` to `sum`, then set `n = n / 10`.
3. After the inner loop, `sum` is the digit sum of the original number.

Keep a `best` variable that holds the minimum digit sum seen so far. Update it whenever the freshly computed `sum` is smaller. After processing every element, return `best`.

Example walk-through with `nums = [999, 19, 199]`:

- `999` → `9 + 9 + 9 = 27`, `best = 27`.
- `19` → `1 + 9 = 10`, `best = 10`.
- `199` → `1 + 9 + 9 = 19`, `best` stays `10`.

Return `10`, matching the expected output.

This works because the minimum-of-transformed-values is independent of order — we never need to look back at earlier elements, so a single streaming pass is sufficient.

## Complexity Analysis

Time Complexity: O(n · d), where `n = len(nums)` and `d` is the number of decimal digits per element. Since `nums[i] <= 10^4`, `d <= 5`, so this is effectively O(n).
Space Complexity: O(1) — only a few integer accumulators are used; the input array is not modified.

## Edge Cases

- **Single-element array**: With `len(nums) == 1`, the answer is simply the digit sum of that one element. Initializing `best` correctly (e.g., to the first element's digit sum or to a sentinel) avoids returning a stale default.
- **All single-digit numbers**: When every `nums[i] < 10`, each digit sum equals the original number, so the answer is just `min(nums)`. The same code path handles this without any special case.
- **Numbers with trailing zeros (e.g., 10, 100, 1000)**: The `n % 10` / `n /= 10` loop correctly skips zero contributions; `10` becomes `1`, `100` becomes `1`, etc.
- **Largest allowed value (10000)**: Digit sum is `1`. The loop must continue until `n` is exactly `0`, not `<= 0`, to handle this cleanly.

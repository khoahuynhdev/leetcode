# 1979. Find Greatest Common Divisor of Array

[LeetCode Link](https://leetcode.com/problems/find-greatest-common-divisor-of-array/)

Difficulty: Easy
Topics: Array, Math, Number Theory
Acceptance Rate: 80.8%

## Hints

### Hint 1

Read the problem carefully: you don't need the GCD of *every* pair of numbers, and you don't need the GCD of the whole array. You only care about two specific values in `nums`. Which two are they, and how expensive is it to find them?

### Hint 2

The two values you need are the smallest and the largest elements. You can find both in a single linear pass over the array — no sorting required. Once you have `min` and `max`, the rest of the problem reduces to a classic, well-known subroutine.

### Hint 3

Computing the greatest common divisor of two integers is done efficiently with the **Euclidean algorithm**: `gcd(a, b) == gcd(b, a % b)`, repeating until the second argument becomes `0`, at which point the first argument is the answer. Combine "one pass to find min and max" with "one Euclidean GCD" and you're done.

## Approach

The problem asks only for the GCD of the smallest and largest numbers in the array, so we never need to consider the other elements beyond scanning them to locate those two extremes.

1. **Find min and max in one pass.** Initialize both `min` and `max` to the first element, then iterate through the array once, updating `min` when a smaller value appears and `max` when a larger value appears. This is O(n) and avoids the O(n log n) cost of sorting.

2. **Compute the GCD with the Euclidean algorithm.** The Euclidean algorithm relies on the identity `gcd(a, b) = gcd(b, a mod b)`. Each step replaces the pair `(a, b)` with `(b, a mod b)`, shrinking the numbers quickly. When `b` reaches `0`, `a` holds the greatest common divisor.

For example, with `nums = [2, 5, 6, 9, 10]`: the single pass gives `min = 2` and `max = 10`. Then `gcd(10, 2)`: `10 mod 2 = 0`, so we move to `gcd(2, 0)`, and since the second argument is `0`, the answer is `2`. That matches the expected output.

Because every `nums[i]` is at least `1`, both `min` and `max` are positive, so the GCD is always well-defined and at least `1`. When all elements are equal (e.g. `[3, 3]`), `min == max` and the GCD is simply that shared value.

## Complexity Analysis

Time Complexity: O(n + log(max)) — O(n) to scan the array for min and max, plus O(log(max)) for the Euclidean algorithm, which is dominated by the scan for typical inputs.
Space Complexity: O(1) — only a constant number of scalar variables are used.

## Edge Cases

- **All elements equal** (e.g. `[3, 3]`): `min` and `max` are identical, and `gcd(x, x) = x`. The algorithm returns that value correctly.
- **Coprime min and max** (e.g. `[7, 5, 6, 8, 3]` → `min = 3`, `max = 8`): the GCD is `1`, the smallest possible answer. Make sure the loop terminates and doesn't accidentally return `0`.
- **min or max equal to 1**: since `1 <= nums[i]`, a `1` in the array forces `min = 1`, and the GCD becomes `1`. The Euclidean algorithm handles this naturally.
- **Minimum-length array** (`nums.length == 2`): the constraint guarantees at least two elements, so initializing min/max from the first element is always safe — there's no empty-array case to guard against.

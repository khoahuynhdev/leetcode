# 3411. Maximum Subarray With Equal Products

[LeetCode Link](https://leetcode.com/problems/maximum-subarray-with-equal-products/)

Difficulty: Easy
Topics: Array, Math, Sliding Window, Enumeration, Number Theory
Acceptance Rate: 55.6%

## Hints

### Hint 1

Look at the constraints: `n <= 100` and every value is between 1 and 10. Checking every subarray is affordable. The real question is how to check one cheaply and without overflowing.

### Hint 2

Fix a start index and extend the end one step at a time, keeping a running product, GCD, and LCM. Each one updates in O(1) (plus a GCD call) when you add an element. Watch out: the product can blow past 64-bit range on long subarrays.

### Hint 3

The LCM of numbers in `[1, 10]` can never be more than `lcm(1..10) = 2520`, and the GCD can never be more than 10. So `lcm * gcd <= 25200`. The product never goes down, because every element is at least 1. Once the product passes 25200, no longer subarray from this start can work, so you can stop extending. That also takes care of overflow.

(Number theory view: for a subarray of length 3 or more, the condition holds exactly when the elements are pairwise coprime. A subarray of length 2 always works, because `a*b = lcm(a,b)*gcd(a,b)`.)

## Approach

1. For each start index `i`, set `prod = 1`, `g = 0` (since `gcd(0, x) = x`), and `l = 1`.
2. For each end index `j` from `i` onward:
   - `prod *= nums[j]`. If `prod > 25200`, break. Extending further can't help, because the product only grows and `lcm*gcd` is capped.
   - `g = gcd(g, nums[j])`
   - `l = l / gcd(l, nums[j]) * nums[j]` (divide first so the intermediate value stays small)
   - If `prod == l * g`, update the answer with `j - i + 1`.
3. Return the best length found.

**Why the early break is correct:** `lcm * gcd` is at most 2520 × 10 for any subarray of these values. Multiplying by more elements (each at least 1) can't make `prod` smaller, so once `prod` is over the bound it stays over it.

**Example:** `[2, 3, 4, 5, 6]` starting at index 1. `[3]` gives 3 = 3·1, length 1. `[3,4]` gives 12 = 12·1, length 2. `[3,4,5]` gives 60 = 60·1, length 3. `[3,4,5,6]` gives prod 360 but lcm 60 and gcd 1, so 360 ≠ 60. The answer is 3.

Long runs of 1s don't change the product, so a subarray like `[1,2,1,1,1]` keeps extending, which is why example 1 returns 5.

## Complexity Analysis

Time Complexity: O(n² · log V), where V ≤ 10. In practice the early break keeps the inner loop short unless the array has many 1s.
Space Complexity: O(1)

## Edge Cases

- **All ones:** the product stays 1 forever, so the whole array is valid. The answer is `n`.
- **Repeated values > 1 (e.g. `[2,2,2]`):** only length-2 windows work. Since `n >= 2`, the answer is always at least 2.
- **Pairwise coprime values spread among 1s (e.g. `[1,7,1,9,1,10,1]`):** the whole array is valid, even though the product (630) is fairly large.
- **Overflow:** without the early break, a run of 100 tens would overflow `int64`. The cap avoids this.
- **LCM update order:** compute `l / gcd(l, x) * x` rather than `l * x / gcd(l, x)` to keep intermediate values small.

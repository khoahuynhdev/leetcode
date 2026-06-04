# 3751. Total Waviness of Numbers in Range I

[LeetCode Link](https://leetcode.com/problems/total-waviness-of-numbers-in-range-i/)

Difficulty: Medium
Topics: Math, Dynamic Programming, Enumeration
Acceptance Rate: 85.1%

## Hints

### Hint 1

Start by reading the constraint carefully: `1 <= num1 <= num2 <= 10^5`. That upper bound is tiny — at most about 100,000 numbers to consider. Before reaching for anything clever, ask whether you can simply *visit every number* in the range and measure its waviness directly. The topic tags (Math, DP, Enumeration) hint at a more scalable technique, but for this "I" version the range is small enough that direct enumeration is comfortably fast.

### Hint 2

The real sub-problem is: given a single number, how do you count its peaks and valleys? Turn the number into its sequence of digits. A digit only qualifies if it has *two* neighbors, so you only ever inspect interior positions `i` from `1` to `len-2`. For each interior digit compare it against its left and right neighbor — that's a constant-size local check, no extra data structure needed.

### Hint 3

The definitions are strict and symmetric: a position contributes `+1` if `digit[i] > digit[i-1] && digit[i] > digit[i+1]` (a peak) **or** `digit[i] < digit[i-1] && digit[i] < digit[i+1]` (a valley). Equality on either side disqualifies the position entirely. Sum that contribution across all interior positions to get one number's waviness, then sum that over the whole range. Any number with fewer than 3 digits has no interior position and therefore contributes `0` automatically — the loop just never executes.

## Approach

We solve the "I" variant by **direct enumeration** because the range size is bounded by `10^5`.

1. **Per-number waviness.** Write a helper that takes an integer `n` and returns its waviness:
   - Extract the digits. The cleanest way is to repeatedly take `n % 10` and `n /= 10`, which produces digits least-significant first; push them into a slice. (The peak/valley test is symmetric under reversal, so the digit order doesn't actually affect the count, but keeping them in natural order makes the logic easy to reason about.)
   - If there are fewer than 3 digits, return `0`.
   - For each interior index `i` (`1 <= i <= len-2`), check whether `digit[i]` is strictly greater than both neighbors (peak) or strictly less than both neighbors (valley). Add `1` for each such position.

2. **Sum over the range.** Loop `x` from `num1` to `num2` inclusive, accumulate `waviness(x)`.

**Why it works:** waviness is defined purely locally — each interior digit independently is or isn't a peak/valley — so summing local contributions is exactly the definition. Enumerating every number guarantees we never miss or double-count any contribution.

**Worked example (`num1 = 198, num2 = 202`):**
- `198` → digits `[1,9,8]`, interior `i=1`: `9 > 1 && 9 > 8` → peak → `1`.
- `199`, `200` → no qualifying interior digit → `0`.
- `201` → digits `[2,0,1]`, `i=1`: `0 < 2 && 0 < 1` → valley → `1`.
- `202` → digits `[2,0,2]`, `i=1`: `0 < 2 && 0 < 2` → valley → `1`.

Total `= 1 + 1 + 1 = 3`. ✅

**Scaling note (honest about difficulty):** the "II" style of this problem would push the bound far beyond `10^5`, where enumeration is hopeless. That is where **digit DP** comes in: you count, over all numbers `<= N`, the total waviness by walking digit positions while tracking the previous two digits and a "tight"/"started" state, then answer the range query as `f(num2) - f(num1 - 1)`. It's a satisfying but considerably trickier technique. For *this* problem you don't need it — but it's the natural next step if you want to stretch.

## Complexity Analysis

Let `R = num2 - num1 + 1` be the count of numbers in the range and `D` the number of digits (`D <= 6` here).

Time Complexity: O(R · D) — for each of `R` numbers we do an `O(D)` digit scan. With `R <= 10^5` and `D <= 6` this is well under a million operations.
Space Complexity: O(D) — a single reusable digit buffer per number; everything else is scalar accumulation.

## Edge Cases

- **Numbers with fewer than 3 digits** (e.g. `1`–`99`): no interior position exists, waviness is `0`. The interior loop simply doesn't run — no special-casing required, but worth confirming the bounds are `1..len-2`.
- **`num1 == num2`** (single-element range, e.g. Example 3 `4848`): the outer loop runs exactly once; make sure the loop is inclusive on both ends.
- **Plateaus / equal neighbors** (e.g. `121` vs `122`, or `2002`): equality must *not* count. Using strict `>` and strict `<` handles this; a digit equal to a neighbor is neither a peak nor a valley.
- **Both a peak and a valley in one number** (e.g. `4848` → waviness `2`): a single number can contribute more than `1`; accumulate per interior position rather than stopping at the first hit.
- **Leading-digit / trailing-digit positions**: the first and last digits can never qualify by definition; the `1..len-2` interior range enforces this for free.

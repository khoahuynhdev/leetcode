# 2243. Calculate Digit Sum of a String

[LeetCode Link](https://leetcode.com/problems/calculate-digit-sum-of-a-string/)

Difficulty: Easy
Topics: String, Simulation
Acceptance Rate: 72.7%

## Hints

### Hint 1

There is no clever trick hiding here: the problem tells you exactly what to do. This is a **simulation** problem, so think about how to faithfully repeat the described process until the stopping condition is met.

### Hint 2

Use an outer loop that runs while `len(s) > k`. Inside it, walk through `s` in chunks of `k` characters (the last chunk may be shorter), sum the digits of each chunk, and append the decimal string of that sum to a new string. Then replace `s` with the new string.

### Hint 3

Converting a digit character to its value is just `c - '0'`, and converting the group sum back to a string is `strconv.Itoa(sum)`. Use a `strings.Builder` (or byte slice) to build each new round efficiently.

## Approach

1. While `len(s) > k`:
   - Create an empty builder.
   - For `i := 0; i < len(s); i += k`, take the group `s[i : min(i+k, len(s))]`.
   - Sum the digits of the group.
   - Append `strconv.Itoa(sum)` to the builder.
   - Set `s` to the builder's string.
2. Return `s`.

**Example:** `s = "11111222223"`, `k = 3`

- Round 1: groups `"111" "112" "222" "23"` → sums `3 4 6 5` → `"3465"` (length 4 > 3, continue)
- Round 2: groups `"346" "5"` → sums `13 5` → `"135"` (length 3 ≤ 3, stop)

Result: `"135"`.

**Why it terminates:** A group of `g` digits sums to at most `9g`, which needs at most `g` characters to write, so a round never makes the string longer. In practice sums are usually much shorter than their groups (a group of 3 digits sums to at most 27 — two characters), so the string shrinks quickly. The rare round that keeps the length (e.g. `"999"`, `k = 2` → `"189"`) is followed by one that shrinks it (`"189"` → `"99"`).

## Complexity Analysis

Time Complexity: O(n · R), where R is the number of rounds. Each round is a linear pass, and because the string shrinks quickly R is small (with n ≤ 100, the work is tiny).
Space Complexity: O(n) for the new string built each round.

## Edge Cases

- **`len(s) <= k` initially:** No rounds happen; return `s` unchanged (e.g., `"1234"`, `k = 4`).
- **All zeros:** Sums are `0`, which must be written as `"0"` (not an empty string) — `strconv.Itoa(0)` handles this.
- **Last group shorter than `k`:** Be careful with slicing bounds; clamp the end index to `len(s)`.
- **Group sum with multiple digits:** A sum like `13` contributes two characters, so the new string can be longer than the number of groups — this is why multiple rounds may be needed.
- **Leading zeros in result:** The result is a string, so leading zeros like `"000"` are preserved and correct.

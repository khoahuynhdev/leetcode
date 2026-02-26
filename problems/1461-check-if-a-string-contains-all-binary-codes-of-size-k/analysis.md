# 1461. Check If a String Contains All Binary Codes of Size K

[LeetCode Link](https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/)

Difficulty: Medium
Topics: Hash Table, String, Bit Manipulation, Rolling Hash, Hash Function
Acceptance Rate: 57.7%

## Hints

### Hint 1

Think about how many distinct binary strings of length `k` exist. What data structure can you use to track which ones you've seen as you scan through `s`?

### Hint 2

A sliding window of size `k` over the string gives you every candidate substring. If you store each substring in a set, you just need to check whether the set reaches the expected size. But creating string keys for every window position is expensive — can you represent each k-length binary substring as an integer instead?

### Hint 3

Use a rolling hash based on bit manipulation. Maintain a `k`-bit integer as you slide across `s`: shift left by 1, OR in the new bit, and mask off the top bit with `(1 << k) - 1`. This gives you the integer value of the current window in O(1) per step. Track seen values in a boolean array of size `2^k` and count down — once the count reaches zero, every code has appeared and you can return early.

## Approach

There are exactly `2^k` distinct binary strings of length `k` (from `000...0` to `111...1`). We need to check whether every one of them appears as a substring of `s`.

**Algorithm (Rolling Bit Hash):**

1. **Early exit:** If `len(s) < k`, there are no substrings of length `k` at all, so return `false`. Also, if the number of length-`k` windows (`len(s) - k + 1`) is less than `2^k`, it's impossible for all codes to appear — return `false`.

2. **Setup:** Create a boolean array `seen` of size `2^k`. Maintain a counter `need` initialized to `2^k` that tracks how many distinct codes we still need to find.

3. **Rolling hash:** Iterate through `s` one character at a time, maintaining an integer `hash` that represents the current window:
   - Shift `hash` left by 1 and OR in the current bit (`s[i] - '0'`).
   - Mask with `(1 << k) - 1` to keep only the lowest `k` bits.
   - Once we've processed at least `k` characters (i.e., `i >= k - 1`), check if `seen[hash]` is false. If so, mark it true and decrement `need`.
   - If `need` reaches 0, return `true` immediately.

4. If we finish scanning without finding all codes, return `false`.

**Example walkthrough** with `s = "00110110"`, `k = 2`:
- Window "00" → hash=0, seen[0]=true, need=3
- Window "01" → hash=1, seen[1]=true, need=2
- Window "11" → hash=3, seen[3]=true, need=1
- Window "10" → hash=2, seen[2]=true, need=0 → return true

## Complexity Analysis

Time Complexity: O(n), where n is the length of `s`. Each character is processed exactly once with O(1) bit operations.

Space Complexity: O(2^k) for the boolean array tracking seen codes. Since `k <= 20`, this is at most ~1 million entries.

## Edge Cases

- **String shorter than k:** If `len(s) < k`, no substring of length `k` exists. Must return `false`.
- **Not enough windows:** If `len(s) - k + 1 < 2^k`, there physically aren't enough substrings to cover all codes. This early check avoids unnecessary work.
- **k = 1:** Only two codes ("0" and "1"). The string just needs to contain both characters.
- **All same characters:** A string like "0000" can only produce the code "00" for `k=2`, so it should return `false` for `k >= 2`.
- **Exact fit:** When the string has exactly `2^k + k - 1` characters and every code appears exactly once (a de Bruijn sequence), the algorithm should still work correctly.

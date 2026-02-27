# 0401. Binary Watch

[LeetCode Link](https://leetcode.com/problems/binary-watch/)

Difficulty: Easy
Topics: Backtracking, Bit Manipulation
Acceptance Rate: 58.9%

## Hints

### Hint 1

Think about what the 10 LEDs on the watch really represent. There are only 4 bits for hours and 6 bits for minutes — the total search space is quite small. Can you enumerate all possible hour and minute values and check which ones are valid?

### Hint 2

Instead of thinking about choosing which LEDs to turn on (a combinatorial problem), consider iterating over every possible time (0:00 to 11:59) and counting how many bits are set in the hour and minute values combined. If that count equals `turnedOn`, the time is valid.

### Hint 3

Use bit counting: for each hour `h` in [0, 11] and each minute `m` in [0, 59], compute `bits.OnesCount(h) + bits.OnesCount(m)`. If the sum equals `turnedOn`, format the time as `"h:mm"` and add it to your result. This avoids any complex backtracking and runs in constant time since the search space is fixed.

## Approach

The key insight is that the problem has a **fixed, tiny search space**. A binary watch can only display hours 0–11 and minutes 0–59. That's at most 12 × 60 = 720 possible times.

Rather than using backtracking to choose which of the 10 LEDs to turn on (which is a valid but more complex approach), we can simply:

1. Iterate over all valid hours (0 to 11).
2. For each hour, iterate over all valid minutes (0 to 59).
3. Count the number of 1-bits in the binary representation of both the hour and the minute.
4. If the total number of set bits equals `turnedOn`, format the time string and add it to the result.

For formatting, the hour has no leading zero (just `%d`), while the minute is always two digits (`%02d`).

**Example walkthrough with turnedOn = 1:**
- h=0, m=1 → bits(0)+bits(1) = 0+1 = 1 ✓ → "0:01"
- h=0, m=2 → bits(0)+bits(2) = 0+1 = 1 ✓ → "0:02"
- h=1, m=0 → bits(1)+bits(0) = 1+0 = 1 ✓ → "1:00"
- h=3, m=0 → bits(3)+bits(0) = 2+0 = 2 ✗

This naturally produces all valid times in sorted order (hours ascending, then minutes ascending).

## Complexity Analysis

Time Complexity: O(1) — We always check exactly 12 × 60 = 720 possible times, regardless of input. Bit counting is O(1) for small fixed-size integers.

Space Complexity: O(1) — Excluding the output list, we only use a constant amount of extra space.

## Edge Cases

- **turnedOn = 0**: Only one valid time: `"0:00"`. Both hour and minute have zero bits set.
- **turnedOn >= 9**: No valid time exists. The maximum bits possible is 4 (hour=11, binary 1011) + 6 (minute=59, binary 111011) = 8 set bits total when hour=11 and minute=59, but that's only 8. Actually max is when h=11(3 bits) and m=59(5 bits) = 8 bits. So turnedOn=9 or 10 returns an empty list.
- **Large turnedOn values (e.g., 10)**: Return empty — impossible to have that many LEDs on with valid hour/minute ranges.

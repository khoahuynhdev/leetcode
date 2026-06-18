# 1344. Angle Between Hands of a Clock

[LeetCode Link](https://leetcode.com/problems/angle-between-hands-of-a-clock/)

Difficulty: Medium
Topics: Math

Acceptance Rate: 66.1%

## Hints

### Hint 1

This is a pure math problem — there is no data structure or search involved. The whole challenge is translating "where the hands point" into angles. Think about each hand independently: if you can compute the absolute angle (measured clockwise from the 12 o'clock position) of the minute hand and of the hour hand, the answer is just the gap between them.

### Hint 2

Convert a full revolution (360 degrees) into the units each hand uses. The minute hand sweeps the whole clock in 60 minutes, so each minute is worth `360 / 60 = 6` degrees. The hour hand sweeps the whole clock in 12 hours, so each hour is worth `360 / 12 = 30` degrees. The tricky part is the hour hand: it does not jump from hour to hour — it drifts continuously as the minutes pass.

### Hint 3

The hour hand moves a little every minute. In 60 minutes it covers 30 degrees, so it gains `0.5` degrees per minute. That means:

- minute angle = `minutes * 6`
- hour angle = `(hour % 12) * 30 + minutes * 0.5`

Take the absolute difference of the two angles. Because the difference might be the "long way around" (greater than 180), the smaller angle is `min(diff, 360 - diff)`. Don't forget `hour % 12` so that 12 maps to 0.

## Approach

The clock face is a circle of 360 degrees. We measure each hand's position as an angle clockwise from the 12 o'clock mark, then take the smaller of the two arcs between them.

Step by step:

1. **Minute hand.** It completes one full circle (360 degrees) every 60 minutes, so it advances `6` degrees per minute. Its angle is `minuteAngle = minutes * 6`.

2. **Hour hand.** It completes one full circle every 12 hours, so it advances `30` degrees per hour. Critically, it also keeps moving during the hour: across 60 minutes it covers those 30 degrees, i.e. `0.5` degrees per minute. We also reduce the hour modulo 12 so that `12` o'clock is treated as `0`. Its angle is `hourAngle = (hour % 12) * 30 + minutes * 0.5`.

3. **Difference.** Compute `diff = |hourAngle - minuteAngle|`. This is one of the two arcs separating the hands; the other arc is `360 - diff`. The answer is the smaller one: `min(diff, 360 - diff)`.

**Worked example** (`hour = 12, minutes = 30`):

- `minuteAngle = 30 * 6 = 180`
- `hourAngle = (12 % 12) * 30 + 30 * 0.5 = 0 + 15 = 15`
- `diff = |15 - 180| = 165`, and `360 - 165 = 195`, so the answer is `min(165, 195) = 165`. ✅

This matches the expected output, and the same formula cleanly produces `75` for `3:30` and `7.5` for `3:15`.

## Complexity Analysis

Time Complexity: O(1) — a fixed number of arithmetic operations regardless of input.
Space Complexity: O(1) — only a couple of scalar values are stored.

## Edge Cases

- **`hour = 12`:** Without `hour % 12`, the hour hand would be computed as `12 * 30 = 360` degrees instead of `0`. The modulo keeps 12 o'clock at the top of the clock.
- **Continuous hour-hand drift:** Forgetting the `minutes * 0.5` term is the most common bug. At `3:30` the hour hand is halfway between 3 and 4, not exactly on 3 — leaving out the drift would give the wrong answer.
- **Reflex angle (long way around):** When the raw difference exceeds 180 degrees, the true "smaller angle" is `360 - diff`. Always taking the minimum of the two arcs handles this.
- **`minutes = 0`:** Both hands sit on exact hour marks; the formula still works (e.g. `3:00` gives `90`).
- **Floating-point tolerance:** Answers within `10^-5` are accepted, so results like `7.5` should be compared with a small epsilon rather than for exact equality.

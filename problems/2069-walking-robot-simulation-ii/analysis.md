# 2069. Walking Robot Simulation II

[LeetCode Link](https://leetcode.com/problems/walking-robot-simulation-ii/)

Difficulty: Medium
Topics: Design, Simulation
Acceptance Rate: 30.4%

## Hints

### Hint 1

The robot only ever walks along the perimeter of the grid. Think about what shape the robot's path traces and how you can represent its position more compactly than (x, y, direction).

### Hint 2

The perimeter of a W x H grid has exactly `2*(W-1) + 2*(H-1)` cells. You can map the robot's position to a single integer index along this perimeter and use modular arithmetic to avoid simulating each individual step.

### Hint 3

The tricky edge case is at position (0, 0). Initially the robot faces East, but if it returns to (0, 0) after walking, it arrived via the left edge traveling South. You need to track whether the robot has ever moved to distinguish these two cases.

## Approach

Model the robot's position as a single index along the perimeter of the grid. The perimeter has `P = 2*(W-1) + 2*(H-1)` positions, and the robot walks in a cycle:

1. **Bottom edge (East):** index 0 to W-1, cells (i, 0)
2. **Right edge (North):** index W to W+H-2, cells (W-1, i-W+1)
3. **Top edge (West):** index W+H-1 to 2W+H-3, cells (2W+H-3-i+W+H-1... simplified: W-2-offset, H-1)
4. **Left edge (South):** index 2W+H-2 to P-1, cells (0, H-2-offset)

When `step(num)` is called, simply do `pos = (pos + num) % P`. This is O(1) per call instead of O(num).

For `getPos()`, convert the perimeter index back to (x, y) coordinates by determining which edge the index falls on.

For `getDir()`, the direction is determined by which edge the index belongs to:
- Bottom edge -> "East"
- Right edge -> "North"
- Top edge -> "West"
- Left edge -> "South"

**Special case:** Position 0 is (0, 0). Initially the robot faces "East", but after any movement that lands back on index 0, the robot arrived from the South direction. Track a `moved` boolean to handle this.

## Complexity Analysis

Time Complexity: O(1) per `step`, `getPos`, and `getDir` call.
Space Complexity: O(1) - only stores width, height, position index, and a boolean flag.

## Edge Cases

- **Full perimeter walk:** After exactly P steps, the robot returns to (0, 0) but faces South, not East.
- **Large step counts:** `num` can be up to 10^5 with up to 10^4 calls, so total steps can reach 10^9. Modular arithmetic is essential; naive simulation would TLE.
- **Minimum grid (2x2):** Perimeter is only 4 cells. Each edge has exactly 1 step. Ensure the formulas work at these boundaries.
- **Multiple full loops:** `num` much larger than P should reduce correctly via modulo.
- **Corner positions:** Corners belong to the edge the robot was traveling along when it arrived (e.g., (W-1, 0) is East, not North).

# 3661. Maximum Walls Destroyed by Robots

[LeetCode Link](https://leetcode.com/problems/maximum-walls-destroyed-by-robots/)

Difficulty: Hard
Topics: Array, Binary Search, Dynamic Programming, Sorting
Acceptance Rate: 33.2%

## Hints

### Hint 1

Since bullets are blocked by neighboring robots, think about what happens when you sort the robots by position. Which robots can possibly destroy a wall located between two consecutive robots?

### Hint 2

After sorting, walls between consecutive robots `i` and `i+1` can only be reached by robot `i` firing right or robot `i+1` firing left (any other robot's bullet would be blocked). This means each "gap" between consecutive robots is independent and affected by at most two robots. Try formulating a DP over sorted robots where the state tracks the firing direction.

### Hint 3

Use DP where `dp[i][d]` represents the maximum walls destroyed considering robots `0..i`, with robot `i` firing in direction `d` (left or right). For each gap between robots `i-1` and `i`, precompute how many walls are reachable by robot `i-1` firing right, robot `i` firing left, or both. The transition combines the previous robot's direction choice with the current one to determine how many gap walls are destroyed. Use binary search on sorted walls to efficiently count walls within any range.

## Approach

1. **Sort robots by position** (keeping their distances paired) and sort walls.

2. **Classify walls into independent zones:**
   - **Walls at robot positions**: These are always destroyed (distance 0 from the robot at that position), so count them directly.
   - **Left outer walls**: Walls to the left of the leftmost robot, reachable only if robot 0 fires left.
   - **Right outer walls**: Walls to the right of the rightmost robot, reachable only if the last robot fires right.
   - **Gap walls**: Walls strictly between consecutive robots `i-1` and `i`. These can only be reached by robot `i-1` firing right or robot `i` firing left.

3. **For each gap**, precompute using binary search on sorted walls:
   - `wallsR`: walls reachable by robot `i-1` firing right (capped at `pos[i-1] + dist[i-1]` and blocked by robot `i`).
   - `wallsL`: walls reachable by robot `i` firing left (capped at `pos[i] - dist[i]` and blocked by robot `i-1`).
   - `wallsBoth`: walls in the overlap region (reachable by either).
   - `wallsUnion = wallsR + wallsL - wallsBoth`.

4. **DP transition**: For each robot `i` (from 1 to n-1), compute:
   - `dp[i][left] = max(dp[i-1][left] + wallsL, dp[i-1][right] + wallsUnion)`
   - `dp[i][right] = max(dp[i-1][left] + 0, dp[i-1][right] + wallsR)`

5. **Final answer**: `max(dp[n-1][left], dp[n-1][right] + rightOuter) + wallsAtRobotPositions`

## Complexity Analysis

Time Complexity: O((n + m) log(n + m)) where n = number of robots, m = number of walls. Sorting takes O(n log n + m log m). The DP is O(n) with O(log m) binary searches per gap.

Space Complexity: O(n + m) for the sorted arrays and robot set.

## Edge Cases

- **Single robot**: No blocking possible; pick the direction that destroys more walls.
- **Wall at a robot's position**: Always destroyed regardless of firing direction (distance 0).
- **Bullet blocked before reaching any wall**: Robot's contribution in that direction is 0 for the gap.
- **Adjacent robots very close together**: The gap may contain no walls, contributing 0.
- **All walls unreachable**: If every wall is beyond all robots' effective ranges, answer is 0 (plus any walls at robot positions).
- **Overlapping ranges in a gap**: Both adjacent robots can reach the same walls; the union avoids double-counting.

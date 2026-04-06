# 0874. Walking Robot Simulation

[LeetCode Link](https://leetcode.com/problems/walking-robot-simulation/)

Difficulty: Medium
Topics: Array, Hash Table, Simulation
Acceptance Rate: 59.3%

## Hints

### Hint 1

Think about how to represent the robot's direction and how turning left or right changes it. A common trick for grid movement is to use a direction array. Also, consider what data structure would let you quickly check whether a position is blocked.

### Hint 2

Use four direction vectors (North, East, South, West) stored in an array, and cycle through them with modular arithmetic when the robot turns. For obstacle lookup, storing obstacles in a hash set will give you O(1) checks. The key detail is that the robot moves **one unit at a time**, so you must check for obstacles at each step.

### Hint 3

The critical insight is that you don't need to track the full path — you just need to track the maximum squared distance seen so far. Move the robot one cell at a time for each forward command, checking the obstacle set before each step. After each step (not just after each command), update the maximum distance. Encode obstacle coordinates into a single key (e.g., a struct or a combined integer) for efficient set lookup.

## Approach

1. **Direction modeling**: Define four direction vectors in order: North `(0,1)`, East `(1,0)`, South `(0,-1)`, West `(-1,0)`. Start facing North (index 0).
   - Turn right (`-1`): increment direction index by 1 (mod 4).
   - Turn left (`-2`): increment direction index by 3 (mod 4), which is equivalent to decrementing by 1.

2. **Obstacle set**: Convert the obstacles array into a Go `map` keyed on a struct `{x, y}` (or encode as a single int64). This gives O(1) lookup per step.

3. **Simulation**: For each command:
   - If it's a turn command, update the direction index.
   - If it's a move command `k`, loop `k` times. Each iteration, compute the next position `(x + dx, y + dy)`. If that position is in the obstacle set, stop moving for this command. Otherwise, update position and check if the new squared distance `x*x + y*y` exceeds the current maximum.

4. **Return** the maximum squared distance observed.

## Complexity Analysis

Time Complexity: O(N * 9 + M) where N is the number of commands and M is the number of obstacles. Each move command moves at most 9 steps, so the simulation is O(N). Building the obstacle set is O(M).

Space Complexity: O(M) for the obstacle hash set.

## Edge Cases

- **No obstacles**: The robot moves freely; the answer is simply the max distance reached along the path.
- **Obstacle at the origin (0, 0)**: The robot starts there and ignores it initially, but cannot return to it later. The simulation naturally handles this since we only check obstacles *before* stepping into them, and the robot starts at (0,0) without "stepping into" it.
- **All turn commands, no movement**: The robot stays at the origin; the answer is 0.
- **Robot blocked immediately**: A move command where the first step is blocked results in no movement for that command.
- **Maximum distance mid-path**: The answer might occur in the middle of the path, not at the end, so we must track the max after every single step.

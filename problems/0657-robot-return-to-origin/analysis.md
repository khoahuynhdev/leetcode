# 0657. Robot Return to Origin

[LeetCode Link](https://leetcode.com/problems/robot-return-to-origin/)

Difficulty: Easy
Topics: String, Simulation
Acceptance Rate: 76.8%

## Hints

### Hint 1

Think about what it means for the robot to return to the origin. Each move changes the robot's position along one of two axes. What must be true about the total displacement for the robot to end up back where it started?

### Hint 2

You don't need to track the full path — just the net effect. Consider counting how each type of move contributes to the robot's final position. Opposite moves cancel each other out.

### Hint 3

The robot returns to the origin if and only if the number of 'L' moves equals the number of 'R' moves AND the number of 'U' moves equals the number of 'D' moves. You can verify this with simple counters — no simulation of coordinates is even necessary.

## Approach

Iterate through the move string and track the robot's position using two variables: one for the horizontal axis (x) and one for the vertical axis (y).

- 'R' increments x by 1
- 'L' decrements x by 1
- 'U' increments y by 1
- 'D' decrements y by 1

After processing all moves, the robot is back at the origin if and only if both x and y are zero.

An equivalent approach is to simply count the occurrences of each character and check that the count of 'L' equals the count of 'R' and the count of 'U' equals the count of 'D'. Both approaches are linear in time and constant in space.

## Complexity Analysis

Time Complexity: O(n), where n is the length of the moves string. We process each character exactly once.

Space Complexity: O(1). We only use a fixed number of integer variables regardless of input size.

## Edge Cases

- **Single move**: A string of length 1 (e.g., "R") can never return to the origin since there is no counterbalancing move.
- **All same direction**: A string like "RRRR" results in a non-zero displacement on one axis.
- **Balanced on one axis only**: A string like "RLUD" must be balanced on both axes independently — being balanced on just one axis is not sufficient (though this specific example is balanced on both).
- **Long alternating moves**: Strings like "RLRLRL..." should correctly return true since every move is immediately cancelled.

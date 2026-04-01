# 2751. Robot Collisions

[LeetCode Link](https://leetcode.com/problems/robot-collisions/)

Difficulty: Hard
Topics: Array, Stack, Sorting, Simulation
Acceptance Rate: 57.0%

## Hints

### Hint 1

Think about which robots can actually collide. Two robots only collide if one is moving right and another (to its right) is moving left. This is similar to classic "asteroid collision" style problems. What data structure is commonly used for those?

### Hint 2

Sort the robots by position so you can process them left to right. Use a **stack** to track right-moving robots that haven't collided yet. When you encounter a left-moving robot, it may collide with the right-moving robots on the stack. Process these collisions one at a time.

### Hint 3

The key insight is handling the collision resolution loop: when a left-moving robot meets the top of the stack (a right-moving robot), compare their healths. The one with lower health is destroyed, and the survivor loses 1 health. If they're equal, both are destroyed. A single left-moving robot might destroy multiple right-moving robots from the stack before it is itself destroyed or survives. After all collisions are resolved, collect survivors and return them in their **original input order**.

## Approach

1. **Index the robots**: Create an array of indices `[0, 1, ..., n-1]` and sort them by their position. This lets us process robots left-to-right spatially while remembering their original order.

2. **Use a stack for right-movers**: Iterate through the sorted indices. If a robot moves right (`R`), push its index onto the stack. If it moves left (`L`), it could collide with right-moving robots already on the stack.

3. **Resolve collisions**: For each left-moving robot, repeatedly compare it with the top of the stack:
   - If the stack's top robot has **less** health: pop it (destroyed), decrease the left-mover's health by 1, and continue checking the stack.
   - If the stack's top robot has **equal** health: pop it (both destroyed), mark the left-mover as destroyed, and stop.
   - If the stack's top robot has **more** health: decrease its health by 1, mark the left-mover as destroyed, and stop.
   - If the stack is empty, the left-mover survives (no more right-movers to collide with).

4. **Collect survivors**: After processing all robots, the survivors are: any left-movers that weren't destroyed, plus any right-movers still on the stack. Gather their healths and return them in the original input order.

**Example walkthrough** (Example 2: positions=[3,5,2,6], healths=[10,10,15,12], directions="RLRL"):
- Sorted by position: index 2 (pos=2,L), index 0 (pos=3,R), index 1 (pos=5,L), index 3 (pos=6,L)
- Process index 2 (L): stack empty, survives.
- Process index 0 (R): push to stack. Stack: [0]
- Process index 1 (L): collide with stack top (index 0). Health 10 vs 10 => both destroyed. Stack: []
- Process index 3 (L): stack empty, survives.
- Wait -- but index 2 is at position 2 moving left, and index 3 is at position 6 moving left. They never collide with each other since they move in the same direction. The survivors are index 2 (health 15) and... let me re-check.

Actually, re-sorting: positions are [3,5,2,6] so sorted order by position is: index 2 (pos=2, dir=R), index 0 (pos=3, dir=L), index 1 (pos=5, dir=R), index 3 (pos=6, dir=L).

- Process index 2 (pos=2, R): push to stack. Stack: [2]
- Process index 0 (pos=3, L): collide with stack top (index 2, health=15). Left-mover health=10 < 15, so left-mover destroyed, stack top health becomes 14. Stack: [2]
- Process index 1 (pos=5, R): push to stack. Stack: [2, 1]
- Process index 3 (pos=6, L): collide with stack top (index 1, health=10). Left-mover health=12 > 10, so stack top destroyed, left-mover health becomes 11. Now collide with new stack top (index 2, health=14). Left-mover health=11 < 14, so left-mover destroyed, stack top health becomes 13. Stack: [2]
- Survivors: index 2 with health 13... Hmm, that doesn't match the expected output of [14].

The expected output [14] means robot 3 (0-indexed index 2, position 2, health 15) survives with health 14. Let me re-read: directions = "RLRL", so index 0='R', index 1='L', index 2='R', index 3='L'.

Sorted by position: index 2 (pos=2, dir='R'), index 0 (pos=3, dir='R'), index 1 (pos=5, dir='L'), index 3 (pos=6, dir='L').

- Index 2 (pos=2, R): push. Stack: [2]
- Index 0 (pos=3, R): push. Stack: [2, 0]
- Index 1 (pos=5, L, health=10): collide with top (index 0, health=10). Equal => both destroyed. Stack: [2]
- Index 3 (pos=6, L, health=12): collide with top (index 2, health=15). 12 < 15 => index 3 destroyed, index 2 health becomes 14. Stack: [2]
- Survivor: index 2 with health 14. Output in original order: [14]. Correct!

## Complexity Analysis

Time Complexity: O(n log n) -- dominated by the sort. Each robot is pushed and popped from the stack at most once, so the collision processing is O(n).

Space Complexity: O(n) -- for the sorted indices array and the stack.

## Edge Cases

- **All robots moving the same direction**: No collisions occur. Return all healths unchanged.
- **All robots destroyed**: Every collision results in mutual destruction. Return an empty array.
- **Single robot**: No collisions possible. Return its health.
- **Chain of collisions**: A single left-moving robot with high health destroys multiple right-moving robots in sequence, losing 1 health each time.
- **Unsorted positions**: The positions are not necessarily in order, so sorting by position is essential before processing.

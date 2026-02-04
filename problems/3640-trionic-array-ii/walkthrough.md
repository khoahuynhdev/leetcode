# Trionic Array II - Solution Walkthrough

## Step 1: Understanding What We're Looking For

A trionic subarray has three parts that form an "up-down-up" pattern:

```
        peak (p)
         /\
        /  \
       /    \
      /      \
start (l)    valley (q)
                    \
                     \
                      \
                       end (r)
```

The constraints are strict: each segment must be STRICTLY monotonic (no equal consecutive elements), and we need indices `l < p < q < r`, meaning each segment has at least 2 elements.

## Step 2: Why Brute Force Fails

The naive approach would enumerate all possible (l, p, q, r) combinations:

```go
// O(n^4) - way too slow for n = 10^5
for l := 0; l < n; l++ {
    for p := l+1; p < n; p++ {
        for q := p+1; q < n; q++ {
            for r := q+1; r < n; r++ {
                if isTrionic(nums, l, p, q, r) {
                    ans = max(ans, sum(nums[l:r+1]))
                }
            }
        }
    }
}
```

Even O(n^2) might be tight. We need O(n).

## Step 3: The Key Insight - Think in States

Instead of finding specific indices, think about what "phase" we're in as we scan left to right:

```
Phase 1: Building the first increasing segment
Phase 2: Building the decreasing segment (after at least one increase)
Phase 3: Building the second increasing segment (valid trionic!)
```

At each position, we don't need to remember WHERE we started - we only need to know the BEST SUM we can achieve if we're in that phase.

## Step 4: Defining the DP States

```
dp1 = maximum sum of a strictly increasing segment ending at current position
dp2 = maximum sum of (increasing then decreasing) ending at current position
dp3 = maximum sum of a complete trionic pattern ending at current position
```

Why "ending at current position"? Because we're scanning left to right, and at each step we extend or transition based on how `nums[i]` compares to `nums[i-1]`.

## Step 5: Working Out the Transitions

At each position i, we compare `nums[i]` with `nums[i-1]`:

### Case A: nums[i] > nums[i-1] (strictly increasing)

We can:

1. **Extend or start dp1**: Add current element to an increasing segment
2. **Extend dp3**: Continue the final increasing phase of a trionic
3. **Transition dp2 → dp3**: The decreasing phase just ended, start final increase

### Case B: nums[i] < nums[i-1] (strictly decreasing)

We can:

1. **Extend dp2**: Continue the decreasing phase
2. **Transition dp1 → dp2**: The first increasing phase just ended, start decreasing

### Case C: nums[i] == nums[i-1] (equal)

Strict monotonicity is broken. All states must reset because we can't continue any segment through equal elements.

## Step 6: The Subtle Part - Starting Fresh vs Extending

For dp1, we have a choice when `nums[i] > nums[i-1]`:

```go
// Option 1: Extend the previous increasing segment
extend := dp1 + nums[i]

// Option 2: Start a fresh segment with just [nums[i-1], nums[i]]
fresh := nums[i-1] + nums[i]
```

Why would we ever start fresh? Consider:

```
nums = [-100, -99, 5, 6]
```

At index 3 (value 6):

- If we extended from index 1: dp1 was -100 + -99 = -199, so extend = -199 + 6 = -193
- If we start fresh: 5 + 6 = 11

Starting fresh is much better! This happens when earlier elements drag down the sum.

## Step 7: Tracing Through an Example

Let's trace `nums = [1, 4, 2, 7]`:

```
Initial: dp1 = -∞, dp2 = -∞, dp3 = -∞

i=1: nums[1]=4 > nums[0]=1 (increasing)
  - dp1: fresh = 1+4 = 5, no previous dp1 to extend
  - dp1 = 5
  - dp2, dp3 unchanged (not decreasing, no dp2 to transition)

  State: dp1=5, dp2=-∞, dp3=-∞

i=2: nums[2]=2 < nums[1]=4 (decreasing)
  - dp2: transition from dp1 = 5 + 2 = 7
  - dp2 = 7
  - dp1 resets (not increasing)

  State: dp1=-∞, dp2=7, dp3=-∞

i=3: nums[3]=7 > nums[2]=2 (increasing)
  - dp1: fresh = 2+7 = 9
  - dp1 = 9
  - dp3: transition from dp2 = 7 + 7 = 14
  - dp3 = 14

  State: dp1=9, dp2=-∞, dp3=14

Answer: max(dp3) = 14 ✓
```

## Step 8: A More Complex Example

Let's trace `nums = [0, -2, -1, -3, 0, 2, -1]`:

```
i=1: -2 < 0 (decreasing)
  - Can't start dp2 (no valid dp1 yet)
  State: dp1=-∞, dp2=-∞, dp3=-∞

i=2: -1 > -2 (increasing)
  - dp1 = -2 + -1 = -3
  State: dp1=-3, dp2=-∞, dp3=-∞

i=3: -3 < -1 (decreasing)
  - dp2 = dp1 + -3 = -3 + -3 = -6
  State: dp1=-∞, dp2=-6, dp3=-∞

i=4: 0 > -3 (increasing)
  - dp1 = -3 + 0 = -3 (fresh start)
  - dp3 = dp2 + 0 = -6 + 0 = -6 (first valid trionic!)
  State: dp1=-3, dp2=-∞, dp3=-6

i=5: 2 > 0 (increasing)
  - dp1 = max(-3+2, 0+2) = max(-1, 2) = 2 (fresh is better!)
  - dp3 = -6 + 2 = -4 (extend the trionic)
  State: dp1=2, dp2=-∞, dp3=-4

i=6: -1 < 2 (decreasing)
  - dp2 = dp1 + -1 = 2 + -1 = 1
  State: dp1=-∞, dp2=1, dp3=-∞ (dp3 resets, but we already recorded -4)

Answer: max(dp3) across all steps = -4 ✓
```

## Step 9: The Final Code with Annotations

```go
package main

const negInf = int64(-1e18)

func maximumTrionicSubarraySum(nums []int) int64 {
    // Three DP states tracking the best sum for each phase
    dp1, dp2, dp3 := negInf, negInf, negInf
    ans := negInf

    for i := 1; i < len(nums); i++ {
        curr := int64(nums[i])
        prev := int64(nums[i-1])

        // New states for this position (start invalid)
        newDp1, newDp2, newDp3 := negInf, negInf, negInf

        if curr > prev { // INCREASING
            // dp1: Can extend previous increasing OR start fresh
            fresh := prev + curr
            if dp1 != negInf {
                newDp1 = max(dp1+curr, fresh)
            } else {
                newDp1 = fresh
            }

            // dp3: Can extend previous trionic OR transition from dp2
            if dp3 != negInf {
                newDp3 = dp3 + curr
            }
            if dp2 != negInf {
                newDp3 = max(newDp3, dp2+curr)
            }

        } else if curr < prev { // DECREASING
            // dp2: Can extend previous decreasing OR transition from dp1
            if dp2 != negInf {
                newDp2 = dp2 + curr
            }
            if dp1 != negInf {
                newDp2 = max(newDp2, dp1+curr)
            }
        }
        // If curr == prev: all newDp stay as negInf (reset)

        // Update states for next iteration
        dp1, dp2, dp3 = newDp1, newDp2, newDp3

        // Track the best complete trionic seen so far
        ans = max(ans, dp3)
    }

    return ans
}
```

## Step 10: Complexity Analysis

**Time: O(n)** - Single pass through the array, O(1) work per element.

**Space: O(1)** - Only tracking 6 variables (3 current states, 3 new states, plus ans).

## Key Takeaways for Similar Problems

1. **State Machine DP** works great for problems with sequential phases or patterns. Other examples: Best Time to Buy/Sell Stock series, Longest Mountain Array.

2. **Track "best so far" not "current"** - When values can be negative, you might want to abandon earlier work and start fresh.

3. **States should be independent** - Each state only depends on the previous iteration's states, enabling O(1) space.

4. **Reset on pattern break** - When strict conditions are violated (like equal elements breaking strict monotonicity), all states must reset.

## Related Problems to Practice

- LeetCode 53: Maximum Subarray (Kadane's algorithm foundation)
- LeetCode 121-123: Best Time to Buy and Sell Stock (state machine DP)
- LeetCode 845: Longest Mountain in Array (similar up-down-up pattern)
- LeetCode 152: Maximum Product Subarray (tracking multiple states)

---

## Why Your Greedy Approach Failed

Your solution has a fundamental conceptual flaw: you're trying to find a single trionic pattern using greedy forward scanning, when the problem requires
considering overlapping possibilities that greedy cannot handle.

### The Core Issue

Your code commits to decisions about where each phase starts/ends. Once you include elements in phase 1, you can't reconsider whether those elements might
work better in a different trionic subarray starting later.

## Specific Bugs in Your Implementation

- Index Management Chaos: You manually increment i inside nested loops while i is also the outer loop variable. The i = i - 1 backtracking attempts patch
  over the fact that your algorithm has lost information about previous positions.

- The Reset Problem: When patterns break, your code resets everything (l, p, q = nums[i]...), but there's no principled way to know where to restart. Should
  you go back one position? Keep some accumulated sum? Greedy can't answer this.

- Missing "Best Ending Here" Paradigm: The correct DP solution asks at each position: "what is the best sum of a partial/complete trionic that ENDS here?"
  Your approach asks "what is THE trionic subarray?" - a question greedy can't answer optimally.

## Why DP Works Here

The DP solution maintains three states at every position:

- dp1: best increasing segment ending here
- dp2: best (increasing→decreasing) ending here
- dp3: best complete trionic ending here

At each step, it considers: extend previous pattern, start fresh, or transition states - and takes the maximum. Your code commits to one path without this
"max" decision.

## The Key Lesson

When you feel the urge to "reset and try again" in your algorithm, that's a signal you might need DP instead of greedy. Greedy makes irreversible
decisions; DP maintains the best option for each possible state.

Practice state machine DP on problems like Best Time to Buy/Sell Stock III and Longest Mountain Array to internalize this pattern.

# 3942. Minimum Operations to Sort a Permutation

[LeetCode Link](https://leetcode.com/problems/minimum-operations-to-sort-a-permutation/)

Difficulty: Medium
Topics: Array
Acceptance Rate: 37.1%

## Hints

### Hint 1

Don't simulate the operations on the array. Ask a different question first: what arrays can you *ever* reach from `nums` using only reverse and rotate-left? The set is much smaller than it looks.

### Hint 2

Think of the array as beads on a necklace. Rotating spins the necklace; reversing flips it over. Neither one changes which beads sit next to each other in the cycle. So `nums` can be sorted only if, read cyclically, it is `0,1,...,n-1` (forwards) or `n-1,...,1,0` (backwards). That leaves at most `2n` reachable configurations.

### Hint 3

Describe every reachable configuration by two numbers: `flip` (has it been reversed an odd number of times?) and `s` (an offset), so that `a[i] = nums[(s + i) mod n]` when `flip = 0`, and `a[i] = nums[(s - i) mod n]` when `flip = 1`. Work out how each operation changes `(flip, s)`. Then the sorted target is fixed by `s = index of 0`, and the answer comes from counting a few short paths. Remember that rotating "backwards" can be done with reverse, rotate, ..., reverse.

## Approach

**1. Model the state.** Every array you can reach is described by `(flip, s)`:

- `flip = 0`: `a[i] = nums[(s + i) mod n]`. Start state is `(0, 0)`.
- `flip = 1`: `a[i] = nums[(s - i) mod n]`.

**2. How the operations change the state.** Substituting into the definitions:

| Operation   | from `flip = 0`      | from `flip = 1`      |
|-------------|----------------------|----------------------|
| Rotate left | `(0, s + 1)`         | `(1, s - 1)`         |
| Reverse     | `(1, s - 1)`         | `(0, s + 1)`         |

Two reverses in a row cancel out, so a shortest path never contains two adjacent reverses. A cheapest path to any state therefore looks like "rotate k times", "rotate k times then reverse", "reverse then rotate k times", or "reverse, rotate k times, reverse". Paths with more reverses never help.

**3. Find the target.** Sorted means `a[i] = i`. In particular `a[0] = 0`, which forces `s = p`, where `p` is the index of `0` in `nums`.

- **Target `(0, p)`** works only if `nums` is *increasing cyclically* from `p`: `nums[(p + i) mod n] == i` for all `i`.
  - Rotate left `p` times: cost `p`.
  - Reverse → `(1, n-1)`, rotate `k` times → `(1, n-1-k)`, reverse → `(0, n-k)`. We need `n-k ≡ p`, so the cost is `2 + (n - p) mod n`. This is the "rotate right" trick, and it wins when `0` is near the end.
- **Target `(1, p)`** works only if `nums` is *decreasing cyclically*: `nums[(p - i) mod n] == i` for all `i`.
  - Rotate `k` times, then reverse: ends at `(1, k - 1)`, so `k = (p + 1) mod n` and the cost is `(p + 1) mod n + 1`.
  - Reverse, then rotate `k` times: ends at `(1, n - 1 - k)`, so the cost is `1 + (n - 1 - p) mod n`.

Take the minimum over every target that is valid. If neither is valid, return `-1`. For `n ≤ 2` both patterns can be valid at once, which is why we check both and don't stop at the first match.

**Example:** `nums = [1, 0, 2]`, `p = 1`.
- Increasing from `p`? `nums[1]=0, nums[2]=2`, so no.
- Decreasing from `p`? `nums[1]=0, nums[0]=1, nums[2]=2`, so yes.
- Costs: `(1+1) mod 3 + 1 = 3` and `1 + (3-1-1) mod 3 = 2`. Answer `2` (reverse, then rotate).

The formulas were checked against a brute-force BFS over every permutation with `n ≤ 7`.

## Complexity Analysis

Time Complexity: O(n) to find `0` and do two linear checks.
Space Complexity: O(1).

## Edge Cases

- **`n = 1`**: already sorted, so the answer is `0`. Modulo arithmetic must not break here.
- **Already sorted**: `p = 0` with an increasing pattern gives `0`. Don't add operations by accident.
- **`n = 2`**: `[1, 0]` is both an increasing and a decreasing rotation. One rotate or one reverse works, so the answer is `1`. Check both patterns.
- **Fully reversed array**: a single reverse, cost `1`.
- **`0` near the end of an increasing rotation**: e.g. `[1,2,3,4,0]`. Reverse + 1 rotate + reverse (cost 3) beats 4 rotations. A solution that only counts left rotations gets this wrong.
- **Not a cyclic run at all**: e.g. `[2,0,1,3]`. Return `-1`. Neither operation can change which elements are neighbors in the cycle.
- **Negative modulo in Go**: `(p - i) % n` can be negative, so normalize with `((x % n) + n) % n`.

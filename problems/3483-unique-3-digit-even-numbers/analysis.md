# 3483. Unique 3-Digit Even Numbers

[LeetCode Link](https://leetcode.com/problems/unique-3-digit-even-numbers/)

Difficulty: Easy
Topics: Array, Hash Table, Recursion, Enumeration
Acceptance Rate: 73.6%

## Hints

### Hint 1

Look hard at the constraints before reaching for clever combinatorics: `digits.length` is at most 10. That is a tiny input, and a tiny input is a strong signal that you are allowed to just *try every possibility* rather than derive a counting formula. Ask yourself: how many three-digit numbers could you possibly build from 10 digits, and is that number small enough to enumerate directly?

### Hint 2

The word **distinct** in the problem statement is doing real work. Two different arrangements of the input can produce the same number (input `[0,2,2]` can form `202` in two ways by picking either copy of `2` last). So enumeration alone is not enough — you need a way to deduplicate the *results*. Every valid answer is a three-digit number, so it lives in a small, fixed range. What data structure does that suggest for the "have I already counted this?" check?

### Hint 3

The trap in this problem is the double meaning of "duplicate." Duplicates in the **input** are legal and must be usable (`[0,2,2]` really can form `220`), while duplicates in the **output** must be counted once. These pull in opposite directions, so handle them with two different mechanisms:

- Enumerate over **positions** (indices `i`, `j`, `k` that are pairwise distinct), never over digit *values*. Picking distinct indices enforces "each copy of a digit can only be used once per number" for free, and naturally lets a repeated digit be reused.
- Deduplicate over **values**, by recording the assembled number `100*d[i] + 10*d[j] + d[k]` in a `seen` set.

Add the two filters the problem states — `digits[i] != 0` (no leading zeros) and `digits[k] % 2 == 0` (even) — and you are done.

## Approach

The search space is tiny. With `n <= 10` digits, the number of ordered triples of distinct indices is at most `10 * 9 * 8 = 720`. There is no need for backtracking, recursion, or a permutation-counting formula: three nested loops cover the entire space.

The algorithm:

1. Allocate `seen`, a `[1000]bool` array. Every three-digit number is in `[100, 999]`, so a fixed-size array is a perfect hash set here — no map, no hashing, no allocation.
2. Loop `i` over all indices for the **hundreds** place. Skip immediately if `digits[i] == 0`, since a leading zero would make the value a two-digit number.
3. Loop `j` over all indices for the **tens** place, skipping `j == i`.
4. Loop `k` over all indices for the **ones** place, skipping `k == i` and `k == j`. Skip if `digits[k]` is odd, since the number must be even.
5. Assemble `value := digits[i]*100 + digits[j]*10 + digits[k]`. If `seen[value]` is false, mark it true and increment the counter.

The key design decision is iterating over **indices, not values**. This is what makes the multiset semantics correct without any extra bookkeeping. The constraint `i != j != k` says "three different physical copies from the array," which is exactly the problem's rule that each copy is used once per number. If instead you looped over the values `0..9` you would need an explicit frequency counter and careful decrement/restore logic — more code and more places to get it wrong.

Walking through `digits = [0,2,2]` (indices `0,1,2` holding values `0,2,2`):

- `i = 0` is skipped outright — `digits[0] == 0` would be a leading zero.
- `i = 1` (value `2`): with `j = 0` (value `0`) and `k = 2` (value `2`) we build `202`, which is even and unseen → count it. With `j = 2` (value `2`) and `k = 0` (value `0`) we build `220`, even and unseen → count it.
- `i = 2` (the other `2`) produces `202` and `220` all over again, but both are already in `seen`, so neither is counted.

Final answer: `2`, matching the expected output. Notice how the second copy of `2` was genuinely usable (that is how `220` gets built at all) yet contributed no double-counting.

If you want to practice the recursion tag attached to this problem, the same idea expresses nicely as a depth-3 backtracking routine over a `used[]` array that inserts into the set at depth 3. It is the same complexity and the same correctness argument; the iterative version just keeps the control flow flat and obvious.

## Complexity Analysis

Time Complexity: O(n^3), where `n = len(digits) <= 10`. That is at most 720 iterations of constant work, so it is effectively O(1) for this problem's constraints.
Space Complexity: O(1) — a fixed `[1000]bool` array regardless of input size.

An O(n) alternative exists: count digit frequencies once, then iterate over the 900 possible values `100..999`, checking whether each value's digit multiset fits inside the input's. That is asymptotically nicer in `n` but strictly slower in practice at `n <= 10`, and it is more code. The brute force is the right call here.

## Edge Cases

- **All zeros, e.g. `[0,0,0]`** — every candidate starts with a leading zero, so the answer is `0`. The `digits[i] == 0` guard must be applied to the hundreds position only; forgetting it produces the bogus `000`.
- **Zero in the tens or ones place, e.g. `[2,0,0]` → `200`** — the mirror of the above. Zero is perfectly legal in the last two positions, so do not filter zeros globally.
- **No even digit at all, e.g. `[1,3,5]`** — the answer is `0`. The evenness check belongs on the *ones* digit specifically, not on any digit.
- **Repeated input digits, e.g. `[0,2,2]` and `[6,6,6]`** — a digit appearing `m` times may be used up to `m` times in one number. Index-based iteration handles this; value-based iteration without a frequency counter would wrongly reject `220` and `666`.
- **Duplicate results from duplicate inputs** — `[6,6,6]` yields `666` through six different index orderings but the answer is `1`. Without the `seen` set you would return `6`.
- **Maximum input size (10 digits)** — worth a sanity check that the triple loop stays comfortably fast; it is only 720 iterations.

This one is genuinely an Easy, and the enumeration itself is straightforward. The part that trips people up is the interplay between duplicate inputs (allowed, must be usable) and duplicate outputs (must be collapsed). Get those two straight and the rest falls out.

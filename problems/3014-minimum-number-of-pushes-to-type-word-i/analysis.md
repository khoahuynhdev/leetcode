# 3014. Minimum Number of Pushes to Type Word I

[LeetCode Link](https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-i/)

Difficulty: Easy
Topics: Math, String, Greedy
Acceptance Rate: 69.9%

## Hints

### Hint 1

Read the constraints again, especially the one that is easy to skim past: **all letters in `word` are distinct**. That means every letter appears exactly once, so no letter is "more important" than another. Ask yourself what the answer can possibly depend on if all letters are interchangeable.

### Hint 2

You get to design the keypad. There are 8 usable keys (2 through 9), and on any single key the letters cost 1, 2, 3, ... pushes depending on their position in that key's list. You want as many letters as possible sitting in the cheap slots. Think greedily: fill every "first slot" across all keys before you are forced to use any "second slot".

### Hint 3

Since the letters are distinct and each is typed once, the total cost is just the sum of the per-letter costs, and the multiset of costs you can achieve is fixed by how you fill the keys. Filling greedily gives: the first 8 letters cost 1 push each, the next 8 cost 2 each, the next 8 cost 3 each, and the rest cost 4. So if you index the letters `0..n-1`, the letter at index `i` costs `i/8 + 1` — and the identity of the letters never matters, only `n = len(word)`.

## Approach

The key realization is that this problem is not really about strings at all. It reduces to a counting argument.

**Why only the length matters.** Each letter of `word` is distinct, so each letter is typed exactly once. The total number of pushes is therefore `sum over letters of (position of that letter on its key)`. Since letters are interchangeable, we only need to choose the *shape* of the assignment: how many letters go on each key.

**Why greedy filling is optimal.** There are 8 keys. Across all keys there are exactly 8 slots that cost 1 push (the first slot on each key), 8 slots that cost 2, 8 that cost 3, and so on. To minimize the sum of `n` chosen slot costs, you simply take the `n` cheapest slots available — there is no constraint preventing this, because slot costs on a key must be used in order (you cannot put a letter in slot 2 without slot 1), and greedy filling naturally respects that: distributing letters round-robin across the 8 keys uses slot `k+1` on a key only after slot `k` is used on every key.

So sorting the achievable slot costs ascending gives:

```
cost:  1 1 1 1 1 1 1 1   2 2 2 2 2 2 2 2   3 3 3 3 3 3 3 3   4 4
index: 0 1 2 3 4 5 6 7   8 9 ...      15   16 ...       23   24 25
```

The letter at index `i` (0-based) costs `i/8 + 1` pushes, using integer division.

**The algorithm.**

1. Let `n = len(word)`.
2. Sum `i/8 + 1` for `i` from `0` to `n-1`.

That is a single loop, and since `n <= 26` the loop runs at most 26 times.

**Worked example — `word = "xycdefghij"`, `n = 10`.** Indices 0..7 each cost 1 (8 pushes), indices 8 and 9 each cost 2 (4 pushes). Total = 8 + 4 = 12, matching the expected output. Note that the answer would be 12 for *any* 10 distinct letters — `"abcdefghij"`, `"qrstuvwxyz"`, anything.

**Closed form (optional).** You can also compute it without a loop. With `q = n / 8` full rounds and `r = n % 8` leftovers, the answer is `8 * (1 + 2 + ... + q) + r * (q + 1)`, i.e. `4*q*(q+1) + r*(q+1)`. The loop is clearer and just as fast here, but the closed form is a nice thing to be able to derive if an interviewer pushes on it. The included solution uses the loop for readability.

A common wrong turn is to build an actual frequency map and sort letters by count. That is the correct approach for the *follow-up* problem (3016, "Minimum Number of Pushes to Type Word II", where letters repeat), but here the distinctness constraint makes all that machinery unnecessary. It is worth solving this one first precisely because it isolates the greedy slot-cost idea before repetition complicates it.

## Complexity Analysis

Time Complexity: O(n), where `n = len(word)` and `n <= 26`. (O(1) with the closed form.)
Space Complexity: O(1) — only a running total is stored.

## Edge Cases

- **`n = 1`** (shortest allowed input): answer is 1. The loop must run at least once; an off-by-one that starts at `i = 1` would break here.
- **`n = 8`** (exactly fills the cheap slots): answer is 8. This is the boundary where every letter still costs 1 push — a good check that you are not prematurely rolling over to cost 2.
- **`n = 9`** (first rollover): answer is 10, not 9. If your division boundary is off by one you will get 9 here.
- **`n = 26`** (all letters, the maximum): answer is `8*1 + 8*2 + 8*3 + 2*4 = 8 + 16 + 24 + 8 = 56`. This exercises the fourth tier, which only a full alphabet reaches.
- **Letter identity is irrelevant**: `"abc"` and `"xyz"` must both return 3. Worth a test to lock in the insight that the content of the string never enters the computation.
- **Empty string**: not reachable under the stated constraints (`1 <= word.length`), but the loop-based solution naturally returns 0, which is the sensible answer anyway.

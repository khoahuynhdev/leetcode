# 3043. Find the Length of the Longest Common Prefix

[LeetCode Link](https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/)

Difficulty: Medium
Topics: Array, Hash Table, String, Trie
Acceptance Rate: 58.4%

## Hints

### Hint 1

The naive idea is to compare every number in `arr1` against every number in `arr2`. With both arrays up to `5 * 10^4` elements, that is `2.5 * 10^9` pairs — far too many. You need a way to avoid the pairwise comparison entirely. Think about what structure lets you ask "does this prefix appear in the other array?" quickly.

### Hint 2

A prefix of an integer is just the number you get by chopping digits off the right end. For example, the prefixes of `12345` are `12345`, `1234`, `123`, `12`, and `1`. Notice that you can generate every prefix of a number cheaply with repeated integer division by 10. If you collect *all* prefixes of one array into a fast-lookup structure, you can check membership for prefixes of the other array.

### Hint 3

Put every prefix of every number in `arr1` into a hash set. Then, for each number `y` in `arr2`, walk its prefixes from longest to shortest (`y`, `y/10`, `y/100`, ...). The first prefix that exists in the set is the longest common prefix involving that `y`. Its digit count is a candidate answer — take the maximum over all of `arr2`. A hash set works because a prefix is uniquely represented by its integer value; a Trie of digits is the equivalent idea, but the integer-set version is simpler in Go.

## Approach

The key realization is that a "common prefix" of two integers is itself an integer, and every prefix can be produced by integer division by 10. So instead of comparing pairs, we precompute one side and query the other.

**Step 1 — Build the prefix set.**
For each `x` in `arr1`, repeatedly insert `x` into a hash set, then do `x /= 10`, until `x` becomes 0. This records every prefix of every number in `arr1`. For `arr1 = [1, 10, 100]` the set becomes `{1, 10, 100}` (1 contributes `1`; 10 contributes `10, 1`; 100 contributes `100, 10, 1`).

**Step 2 — Query with arr2.**
For each `y` in `arr2`, we want the longest prefix of `y` that is also a prefix of some number in `arr1`. Start from `y` itself (its full digit count) and strip one digit at a time. The moment a value is found in the set, it is the longest common prefix for that `y`, so we record its length and stop. Track the global maximum across all of `arr2`.

For `arr2 = [1000]`: check `1000` (not in set), then `100` (in set!) — length 3. Answer is 3.

**Why it works.** A common prefix `c` of `x` and `y` must be a prefix of both. By storing all prefixes of `arr1`, the set contains exactly the integers that are prefixes of *something* in `arr1`. Stripping digits off `y` from the right enumerates `y`'s prefixes from longest to shortest, so the first hit is guaranteed to be the longest shared prefix for that `y`. Taking the max over `arr2` gives the answer over all pairs.

This is the hash-set incarnation of the Trie approach: inserting digit-by-digit into a Trie and then matching is logically the same as inserting integer prefixes and looking them up, but avoids the bookkeeping of node objects.

## Complexity Analysis

Let `N = len(arr1)`, `M = len(arr2)`, and `D` = maximum number of digits (here at most 9, since values are `<= 10^8`).

Time Complexity: O((N + M) * D) — each number contributes at most `D` insertions or lookups, and each hash operation is O(1) average. With `D` bounded by a small constant, this is effectively linear in the input size.

Space Complexity: O(N * D) — the hash set holds at most `D` prefixes per element of `arr1`.

## Edge Cases

- **No common prefix at all** (e.g. `arr1 = [1,2,3]`, `arr2 = [4,4,4]`): no prefix of any `arr2` element is in the set, so the answer stays `0`. The problem explicitly asks to return `0` here.
- **Same-array prefixes must not count**: only cross-array pairs `(x in arr1, y in arr2)` matter. The algorithm naturally respects this because the set is built only from `arr1` and queried only from `arr2`.
- **One number is a full prefix of the other** (e.g. `arr1 = [100]`, `arr2 = [1]`): the common prefix is `1`, length 1 — not 3. Walking `y = 1`'s prefixes (just `1`) finds it immediately. The shorter number caps the prefix length.
- **Identical numbers** (e.g. `arr1 = [12345]`, `arr2 = [12345]`): the whole number is a common prefix, giving its full digit count (5).
- **Maximum-size values** (e.g. `100000000`): nine digits; integer division stays well within `int` range, so no overflow concerns in Go.

This problem is approachable once the "prefix = repeated /10" insight lands, but the trap is the O(N*M) brute force — recognizing you must precompute one side is the real challenge. Stick with it; the pattern of "hash all prefixes, then probe" generalizes to many string/Trie problems.

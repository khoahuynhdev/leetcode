# 3517. Smallest Palindromic Rearrangement I

[LeetCode Link](https://leetcode.com/problems/smallest-palindromic-rearrangement-i/)

Difficulty: Medium
Topics: String, Sorting, Counting Sort
Acceptance Rate: 67.6%

## Hints

### Hint 1

A palindrome is fully determined by its first half. If you know the left half and the (optional) middle character, the right half is forced — it is just the mirror image. So stop thinking about "rearranging a string" and start thinking about "what multiset of characters goes in the left half?"

Also, the alphabet is tiny (26 lowercase letters). Whenever the alphabet is bounded and you need order, a frequency array plus a sweep from `'a'` to `'z'` beats a comparison sort.

### Hint 2

Count how many times each letter appears. The input is guaranteed to be a palindrome, which tells you something strong about those counts: every letter appears an even number of times, except possibly one letter when the length is odd. That single odd letter has exactly one "unpairable" copy.

Each letter contributes `count/2` copies to the left half and `count/2` copies to the right half. The leftover copy (if any) must sit dead center — you have no choice about which character goes in the middle.

### Hint 3

The critical insight: comparing two palindromes of the same length is decided by their left halves, because the right half is a mirror of the left. So minimizing the palindrome reduces to minimizing the left half as a string.

And since the left half's multiset of characters is *fixed* (it is exactly `count[c]/2` copies of each `c`), the smallest arrangement of a fixed multiset is simply its sorted order. There is no greedy choice to agonize over: emit letters `'a'` through `'z'` in order, `count[c]/2` of each. Then append the forced middle character, then the reverse of the half. Done — no sorting of the original string required, just a counting sweep.

## Approach

**Step 1 — Count.** Build `cnt[26]`, where `cnt[c]` is the number of occurrences of letter `c` in `s`. This is a single O(n) pass.

**Step 2 — Build the left half.** For `c` from `'a'` to `'z'`, append `cnt[c]/2` copies of `c` to a buffer `half`. Because we iterate the alphabet in ascending order, `half` comes out already sorted ascending, which is the lexicographically smallest permutation of that multiset. Its length is `n/2`.

**Step 3 — Find the middle.** If `n` is odd, exactly one letter has an odd count (guaranteed, since `s` is a palindrome). Scan the counts for the letter with `cnt[c]%2 == 1` and place it in the center. If `n` is even, there is no middle character.

**Step 4 — Mirror.** The answer is `half + middle + reverse(half)`.

**Why the sorted half is optimal.** Take any palindromic permutation `p` of `s`. Its first `n/2` characters must be a permutation of the same multiset that `half` uses — the counts of `p` equal the counts of `s`, and `p` being a palindrome forces its first half to contain exactly `cnt[c]/2` copies of each `c`. Among all permutations of a fixed multiset, the ascending-sorted one is lexicographically smallest. Since `p` and our answer have equal length, the comparison is resolved at the first index where they differ, and that index falls inside the first half unless the halves are identical (in which case the middles are also forced equal, so the strings are identical). Hence our answer is `<=` every palindromic permutation.

**Worked example — `s = "babab"`:**

- Counts: `a: 2`, `b: 3`. Length `n = 5` (odd).
- Left half: `a` contributes `2/2 = 1` copy, `b` contributes `3/2 = 1` copy → `half = "ab"`.
- Middle: `b` has an odd count → middle is `'b'`.
- Answer: `"ab"` + `"b"` + reverse(`"ab"`) = `"ab" + "b" + "ba"` = `"abbba"`. ✓

**Worked example — `s = "daccad"`:**

- Counts: `a: 2`, `c: 2`, `d: 2`. Length `n = 6` (even).
- Left half: `"a" + "c" + "d"` = `"acd"`.
- No middle.
- Answer: `"acd" + "dca"` = `"acddca"`. ✓

Honest note on difficulty: this is a Medium that feels like an Easy *once* you see that the right half is not a free variable. The trap is trying to greedily construct the whole string left-to-right while maintaining palindromicity — that framing is much harder than it needs to be. If you got stuck there, that is a useful lesson to keep: for palindrome-construction problems, always ask "what is the minimal set of independent choices here?"

## Complexity Analysis

Time Complexity: O(n + Σ) where `n = len(s)` and `Σ = 26` is the alphabet size — one pass to count, one sweep over the alphabet emitting a total of `n/2` characters, and one reversal pass. This simplifies to O(n).

Space Complexity: O(n) for the output string and the half buffer, plus O(Σ) = O(1) for the frequency array. Excluding the required output, the auxiliary space is O(n) for the half buffer (or O(1) auxiliary if you write directly into a preallocated result of size `n`, which is what the implementation does — it fills position `i` and its mirror `n-1-i` together).

## Edge Cases

- **Single character (`"z"`).** `n = 1`, the half is empty, and the whole answer is the middle character. Make sure your code does not assume a non-empty half or index into it.
- **Odd length with the odd-count letter not being the smallest.** In `"babab"` the middle is `'b'`, not `'a'`. You do not get to choose the middle character — it is dictated by which count is odd. A common bug is placing the smallest letter in the center.
- **Even length — no middle at all.** Guard the middle-character logic behind an `n%2 == 1` check, otherwise you will emit a spurious character and produce a string of the wrong length.
- **All characters identical (`"aaaa"`, `"aaaaa"`).** The answer equals the input. Exercises both the even and odd branches with a single distinct letter.
- **Already-minimal input (`"acddca"`).** The function must be idempotent — feeding it its own output returns the same string. Good sanity check.
- **Input is already sorted-looking but not minimal (`"bbaabb"` → `"abbbba"`).** Confirms you are rebuilding from counts rather than lightly permuting the input.
- **Maximum length (10^5).** Avoid an O(n log n) sort of the full string or repeated string concatenation in a loop (which is O(n^2) in Go). Use a counting sweep and a preallocated `[]byte`.

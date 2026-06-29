# 1967. Number of Strings That Appear as Substrings in Word

[LeetCode Link](https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/)

Difficulty: Easy
Topics: Array, String
Acceptance Rate: 84.1%

## Hints

### Hint 1

This problem is essentially a counting problem layered on top of a substring-membership test. Ask yourself: for a single pattern, how would you decide whether it appears somewhere inside `word`? Once you can answer that for one pattern, the rest is just repetition.

### Hint 2

The core operation you need is "does string `B` contain string `A` as a contiguous block of characters?" This is the classic *substring search* operation. Most languages give it to you directly (`strings.Contains` in Go), so you don't need to hand-roll any matching algorithm here. Think about wrapping that check in a loop.

### Hint 3

Iterate over every pattern, run the substring check against `word`, and increment a counter each time the check succeeds. There is no trick beyond this — the constraints are tiny (at most 100 patterns, each at most 100 characters, `word` at most 100 characters), so a straightforward scan is more than fast enough. Don't over-engineer it.

## Approach

The task asks for the number of strings in `patterns` that occur as a substring of `word`.

The algorithm is direct:

1. Initialize a counter `count = 0`.
2. For each `p` in `patterns`, test whether `p` is a substring of `word`.
3. If it is, increment `count`.
4. Return `count`.

In Go, step 2 is exactly `strings.Contains(word, p)`, which returns `true` when `p` appears as a contiguous sequence inside `word`. Under the hood this is an efficient substring search (Go uses a Rabin–Karp / Boyer–Moore-style implementation), but conceptually you can think of it as sliding `p` across `word` and comparing.

**Walking through Example 1:** `patterns = ["a","abc","bc","d"]`, `word = "abc"`.

- `"a"` → contained in `"abc"` → count = 1
- `"abc"` → contained in `"abc"` → count = 2
- `"bc"` → contained in `"abc"` → count = 3
- `"d"` → not contained → count stays 3

Result: `3`, which matches the expected output.

Why it works: "appears as a substring" is precisely the membership relation that `strings.Contains` tests, and counting independent successes is just a running tally. Because each pattern is tested independently, duplicates in `patterns` are each counted separately (see Example 3, where `["a","a","a"]` yields `3`).

If you wanted to implement the substring check yourself, you would, for each starting index `i` in `word`, compare `word[i:i+len(p)]` against `p` and stop at the first match. That is the naive O(n·m) check, and it is perfectly acceptable at these constraints.

## Complexity Analysis

Let `k` be the number of patterns, `m` the maximum pattern length, and `n` the length of `word`.

Time Complexity: O(k · n · m) in the worst case — for each of the `k` patterns, the substring search may scan up to `n` starting positions, each comparing up to `m` characters. With the given limits (k, m, n ≤ 100) this is at most on the order of 10^6 character comparisons, effectively instant. (Go's optimized `strings.Contains` is typically closer to O(k · (n + m)).)

Space Complexity: O(1) — only a single integer counter is used beyond the input.

## Edge Cases

- **Pattern longer than `word`:** A pattern with more characters than `word` can never be a substring, so the check correctly returns `false`. The constraints allow patterns up to length 100 and `word` as short as length 1, so this genuinely happens.
- **Duplicate patterns:** `["a","a","a"]` must count each occurrence, giving `3`, not `1`. Because we test each list element independently, duplicates are handled naturally — do not deduplicate.
- **No matches at all:** Every pattern failing the check should return `0`, not an error. The counter simply stays at its initial value.
- **All patterns match:** When every pattern is found, the answer equals `len(patterns)`.
- **Single-character patterns and word:** The smallest inputs (`word` of length 1, patterns of length 1) still work because `strings.Contains` treats single characters as ordinary substrings.

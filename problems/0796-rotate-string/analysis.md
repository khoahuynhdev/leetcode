# 0796. Rotate String

[LeetCode Link](https://leetcode.com/problems/rotate-string/)

Difficulty: Easy
Topics: String, String Matching
Acceptance Rate: 65.6%

## Hints

### Hint 1

A "shift" produces a *rotation* of the original string. Before reaching for a loop that simulates each shift, ask yourself: is there a structural property of all rotations of a string that lets you check membership in one shot? Think of this as a string-matching problem rather than a simulation problem.

### Hint 2

If you concatenate `s` with itself, what does the resulting string contain? Try writing out `s + s` for `s = "abcde"` and circle every length-`5` window. Notice anything?

### Hint 3

Every rotation of `s` is a contiguous substring of `s + s`. So `goal` is reachable from `s` via shifts **iff** the two strings have the same length **and** `goal` appears as a substring of `s + s`. Reduce the whole problem to one substring check.

## Approach

The brute-force idea is to try every possible shift of `s` (there are `n` of them) and compare against `goal`, giving `O(n^2)` work. That's perfectly fine given the constraint `n <= 100`, but there's a much cleaner observation that turns this into a single substring search.

**Key observation.** Any rotation of `s` is obtained by picking some split point `i` and concatenating `s[i:] + s[:i]`. If we look at `s + s`, every such rotation appears as the length-`n` window starting at index `i`:

```
s     = "abcde"
s + s = "abcdeabcde"
         ^ window 0: "abcde"
          ^ window 1: "bcdea"
           ^ window 2: "cdeab"  <-- matches goal
            ^ window 3: "deabc"
             ^ window 4: "eabcd"
```

So the question "can `s` become `goal` via shifts?" is equivalent to:

1. `len(s) == len(goal)` (otherwise impossible), **and**
2. `goal` is a substring of `s + s`.

Both checks together solve the problem. The length check is essential: without it, `s = "ab"` and `goal = "a"` would incorrectly succeed because `"a"` is a substring of `"abab"`.

In Go, `strings.Contains(s+s, goal)` performs the substring search. The standard library's implementation runs in linear time on average (and worst-case linear with the underlying Rabin–Karp / two-way variant Go uses), so the overall solution is `O(n)` time with `O(n)` extra space for the doubled string.

## Complexity Analysis

Time Complexity: O(n), where `n = len(s)`. The doubled string has length `2n`, and the substring search runs in linear time.

Space Complexity: O(n) for the concatenated string `s + s`. The check itself uses constant additional space.

## Edge Cases

- **Different lengths.** If `len(s) != len(goal)`, no number of shifts can make them equal — return `false` immediately. This also guards against the "substring of a doubled string" trick falsely matching shorter `goal`s.
- **Identical strings (`s == goal`).** Zero shifts is allowed, so this should return `true`. The substring approach handles it naturally because `s` is a prefix of `s + s`.
- **Single-character strings.** With `n = 1`, the only rotation is `s` itself; the answer is `s == goal`. The general logic still works.
- **Same multiset of characters but not a rotation.** E.g. `s = "aabbcc"`, `goal = "abcabc"`. They share characters but `goal` is not a rotation, so it must not appear in `s + s = "aabbccaabbcc"`. The substring check correctly returns `false`.
- **Repeated patterns.** E.g. `s = "aaaa"`, `goal = "aaaa"`. Every rotation equals `s`, and `goal` trivially appears in `s + s`.

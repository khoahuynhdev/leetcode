# 3499. Maximize Active Section with Trade I

[LeetCode Link](https://leetcode.com/problems/maximize-active-section-with-trade-i/)

Difficulty: Medium
Topics: String, Enumeration

Acceptance Rate: 41.1%

## Hints

### Hint 1

The described trade sounds complicated, but start by asking: what does a single trade *actually change* in the string? Try writing out a small example like `"1000100"` and physically applying the two steps. Notice how the region you touch changes shape — this is a classic case where compressing the string into **runs of equal characters** (run-length encoding) makes the structure obvious.

### Hint 2

A trade always operates on a `1`-block that is **surrounded by `0`-blocks on both sides**, i.e. a pattern `[0s][1s][0s]`. Step one turns the middle `1`s into `0`s (merging the three runs into one big `0`-block), and step two turns that merged `0`-block back into `1`s. So really you are enumerating every `1`-run that sits between two `0`-runs. What is the *change* in the number of active sections for each such choice?

### Hint 3

For a region `[a zeros][m ones][b zeros]`, after the trade it becomes `a + m + b` ones. It used to contribute `m` ones, so the **net gain is exactly `a + b`** — the middle `m` cancels out. That means the answer is simply `totalOnes + max(a + b)` over all valid sandwiches, and if no `1`-run is surrounded by two `0`-runs, no trade helps and the answer is just `totalOnes`.

## Approach

The key realization is that the number of active sections can never drop below the original count of `'1'`s (you can always choose to do nothing). So the base of the answer is `totalOnes`.

A valid trade requires a block of `'1'`s that is surrounded by `'0'`s on both sides. The augmentation note (`t = '1' + s + '1'`) only affects whether `0`-blocks at the ends count as "surrounded by `1`s" — it never creates a new `1`-block sandwich, so we can reason directly about the original string's runs.

Consider any region shaped like `[a zeros][m ones][b zeros]`:

1. Convert the `m` ones (surrounded by zeros) to zeros → the three runs merge into one block of `a + m + b` zeros.
2. Convert that merged zero block (now surrounded by ones) to ones → `a + m + b` active sections.

Before the trade this region held `m` ones; after, it holds `a + m + b`. The **net gain** is `(a + m + b) - m = a + b`. The middle `1`-block size is irrelevant to the gain — only the two neighboring `0`-block sizes matter.

So the algorithm is:

1. Count `totalOnes`.
2. Scan the string as consecutive runs. Whenever a `0`-run is preceded (two runs back) by another `0`-run — meaning there is a `1`-run sandwiched between them — record the sum of the two `0`-run lengths as a candidate gain.
3. Answer = `totalOnes + max candidate gain` (or `totalOnes` if no sandwich exists).

Because runs strictly alternate between `0` and `1`, tracking just the length of the *previous* `0`-run is enough: when we reach a new `0`-run and a previous `0`-run exists, there is guaranteed to be exactly one `1`-run between them.

**Worked example** — `"1000100"`:
- Runs: `1`(len 1), `000`(len 3), `1`(len 1), `00`(len 2).
- `totalOnes = 2`.
- The `00` run is preceded by the `000` run (with a `1` in between): candidate gain `3 + 2 = 5`.
- Answer `= 2 + 5 = 7`. ✔

## Complexity Analysis

Time Complexity: O(n) — a single pass to count ones plus a single pass over the runs.
Space Complexity: O(1) — only a few counters are kept; no extra structure proportional to the input.

## Edge Cases

- **All `'1'`s (e.g. `"111"`)**: no `0`-runs, so no trade is possible; answer is `totalOnes`.
- **All `'0'`s (e.g. `"0000"`)**: only one `0`-run and no `1`-run to sandwich; gain is 0; answer is 0.
- **Single character (`"0"` or `"1"`)**: too short for any sandwich; answer is `totalOnes`.
- **No valid sandwich with ones at the ends (e.g. `"1001"`)**: the `0`-block is surrounded by `1`s but there is no `1`-block surrounded by `0`s, so no trade applies; answer is `totalOnes`.
- **Multiple candidate sandwiches (e.g. `"01010"`)**: must take the maximum `a + b`, not just the first one found.

Don't be discouraged if the phrasing of the trade felt intimidating at first — the whole puzzle collapses into "add the two 0-runs around a 1-run" once you compress the string into runs. That reframing is the real skill this problem is training.

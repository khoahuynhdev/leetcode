# 1291. Sequential Digits

[LeetCode Link](https://leetcode.com/problems/sequential-digits/)

Difficulty: Medium
Topics: Enumeration
Acceptance Rate: 66.0%

## Hints

### Hint 1

Before reaching for a data structure, ask a smaller question: *how many* integers with sequential digits even exist? A number that starts at some digit and increases by one each step is completely determined by two things — where it starts and how long it is. That severely limits the possibilities. Think about counting the whole universe of valid numbers rather than scanning every number in `[low, high]`.

### Hint 2

The range can be as large as `10^9`, so iterating from `low` to `high` and checking each number is far too slow. Instead, *generate* only the numbers that qualify. Every sequential-digit number is a contiguous substring of a single fixed string. What is that string, and how can you slide a window across it?

### Hint 3

The key insight: every sequential-digit number is a substring of `"123456789"`. There are only 36 such substrings of length ≥ 2. Slide a window of each possible length (2 through 9) across `"123456789"`, convert each window into an integer, and keep the ones that fall inside `[low, high]`. If you enumerate by increasing length, and within each length from left to right, the results come out already sorted.

## Approach

The whole problem hinges on one observation: a number has sequential digits exactly when its digit string is a contiguous slice of `"123456789"`. For example `345` is the slice `"345"`, and `6789` is the slice `"6789"`. No sequential-digit number can contain a `0` (there is nothing before `1`), and none can wrap past `9`, so this single string captures every possibility.

That means the entire universe of sequential-digit numbers is tiny and fixed, independent of the input:

- Length 2: `12, 23, 34, 45, 56, 67, 78, 89`
- Length 3: `123, 234, ..., 789`
- ...
- Length 9: `123456789`

The algorithm:

1. Let `digits = "123456789"`.
2. For each window `length` from `2` to `9`:
   - For each starting index `start` such that the window `[start, start+length)` fits inside the string:
     - Build the integer by folding the window's characters: `num = num*10 + (digits[i] - '0')`.
     - If `low <= num <= high`, append it to the result.
3. Return the result.

Because we iterate length ascending (shorter numbers are always numerically smaller) and, within a length, left-to-right (larger starting digit means a larger number), the output is naturally in sorted order. A defensive `sort.Ints` is included but is not strictly required.

**Walkthrough for `low = 100, high = 300`:** Length-2 windows produce `12..89`, all below `100`, so none qualify. Length-3 windows produce `123, 234, 345, ...`; `123` and `234` are in `[100, 300]`, but `345` and beyond exceed `300`. Length-4+ numbers are all `≥ 1234 > 300`. Result: `[123, 234]`. ✅

## Complexity Analysis

Time Complexity: O(1) — the number of candidate substrings of `"123456789"` is a fixed constant (36 numbers of length ≥ 2), independent of `low` and `high`. Building each number costs at most 9 steps. The optional sort is over ≤ 36 elements, also constant.

Space Complexity: O(1) — aside from the output list, only a few scalar variables are used. The output itself holds at most 36 integers.

## Edge Cases

- **No numbers in range** (e.g. `low = 140, high = 150`): No sequential-digit number lands in this window, so the answer is an empty list. The generate-and-filter approach handles this naturally by simply appending nothing.
- **Single valid number** (e.g. `low = 123, high = 123`): The range collapses to one value; the inclusive bounds `low <= num <= high` must both be non-strict so the endpoint is captured.
- **Boundary around a length change** (e.g. `low = 58, high = 155`): The valid numbers straddle two-digit and three-digit results (`67, 78, 89, 123`), confirming that mixing lengths still yields correctly sorted output.
- **Upper limit `10^9`**: The largest sequential-digit number is `123456789` (9 digits), which is less than `10^9` (10 digits). So a range like `low = high = 10^9` yields an empty list — make sure the loop never tries a length beyond 9.

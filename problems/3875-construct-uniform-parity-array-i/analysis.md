# 3875. Construct Uniform Parity Array I

[LeetCode Link](https://leetcode.com/problems/construct-uniform-parity-array-i/)

Difficulty: Easy
Topics: Array, Math
Acceptance Rate: 80.6%

## Hints

### Hint 1

The actual values of `nums1` never matter — only whether each one is odd or even. Rewrite the problem in your head with `nums1` replaced by a string of `E`s and `O`s and ask what each allowed move does to a single letter. Anything about the numbers being distinct, or bounded by 100, is scenery.

### Hint 2

Parity arithmetic has only four rules: `E - E = E`, `O - O = E`, `E - O = O`, `O - E = O`. So the move `nums2[i] = nums1[i] - nums1[j]` is really "flip index `i`'s parity if `nums1[j]` is odd, leave it alone if `nums1[j]` is even." Keeping `nums2[i] = nums1[i]` is "leave it alone." Now ask: what do you need available in the array in order to flip a chosen index?

### Hint 3

You need exactly one thing: a single odd element somewhere. Nothing in the problem says an index `j` may be used only once, so one odd value at index `k` can be subtracted from *every* other index. That means whenever the array is mixed (at least one odd and at least one even), you can keep every odd as-is and subtract `nums1[k]` from every even, turning them all odd. And if the array is *not* mixed, it is already uniform, so you keep everything. There is no arrangement of parities left over — the answer is always `true`.

## Approach

Work entirely in parity space. Let `O` be the set of indices holding odd values and `E` the set holding even values.

**Case 1: the array is already uniform** (`O` is empty, or `E` is empty). Choose the first option, `nums2[i] = nums1[i]`, for every index. `nums2` is a copy of `nums1`, so it is all even or all odd respectively. Done.

**Case 2: the array is mixed** — both `O` and `E` are non-empty. Pick any single index `k ∈ O`. Build `nums2` like this:

- for `i ∈ O`: take `nums2[i] = nums1[i]`, which is odd;
- for `i ∈ E`: take `nums2[i] = nums1[i] - nums1[k]`, which is `even - odd = odd`.

The second choice is legal for every `i ∈ E`, because `i` holds an even value and `k` holds an odd one, so `i != k` is automatic — we never need the forbidden `j == i`. And critically, the problem places no budget on `k`: the same index may serve as the subtrahend for all of `E`. Every entry of `nums2` is odd, so the array is uniform.

Those two cases cover every possible input, so the answer is unconditionally `true`.

A worked example, `nums1 = [4, 7, 10, 12]`: parities are `E O E E`, which is mixed, so take `k = 1` (value `7`). Then `nums2 = [4-7, 7, 10-7, 12-7] = [-3, 7, 3, 5]` — all odd. Note how index 1 was reused three times.

It is worth being explicit about why the *other* target is not needed. Making everything **even** is the harder goal: an odd index `i` can only be fixed by subtracting some *other* odd value, so a lone odd element would be stuck. That is exactly the case that tempts people into answering `false`. But the all-odd target rescues it: with one odd element present, every even element has something to borrow, and the lone odd simply keeps itself. Having two escape routes and only needing one is what collapses this problem.

Both of the problem's examples return `true`, and no `false` example is given — a fair signal that you are meant to discover there is no failing input.

The honest note on difficulty: the *code* here is a one-liner, but the proof is the whole exercise, and "prove no counterexample exists" is a genuinely different skill from "find the algorithm." Being suspicious of a trivial answer is healthy; the fix is to write out the case analysis until you are convinced, rather than trusting the hunch either way.

## Complexity Analysis

Time Complexity: O(1) — the answer does not depend on reading the input at all. Any implementation that scans the array to classify parities is O(n), which is also perfectly acceptable.
Space Complexity: O(1) — `nums2` never has to be materialized, only shown to exist.

## Edge Cases

- **`n == 1`**: the second option is unusable, since no index `j != i` exists. It does not matter — a one-element array is vacuously "all odd" or "all even" via the keep-it option. This is the case most likely to trip up an implementation that assumes a partner index is always available.
- **All elements even** (`[4, 6]`, example 2): no odd value exists to flip anything with, so the all-odd target is unreachable — but the array is already all even, so keeping everything works.
- **All elements odd**: symmetric; keep everything.
- **Exactly one odd, many evens** (`[1, 2, 4, 6]`): the case that breaks the all-*even* target, because the single odd has no other odd to pair with. The all-odd target handles it, and the single odd is reused as the subtrahend for every even. Good sanity check for anyone who derived a "need at least two odds" rule.
- **Exactly one even, many odds** (`[2, 3, 5, 7]`): the mirror case; the single even borrows any odd.
- **Reuse of `j`**: the construction leans on subtracting the same `nums1[k]` from many indices. If you mistakenly assume each `j` may be consumed only once, you will derive a much stricter condition and start returning `false` on inputs like `[1, 2, 4]`. Re-read the statement: the only restriction on `j` is `j != i`.
- **Distinct values and the `1 <= nums1[i] <= 100` bound**: neither constrains the answer. They exist to keep the input tidy, not to gate the logic — do not build them into your reasoning.

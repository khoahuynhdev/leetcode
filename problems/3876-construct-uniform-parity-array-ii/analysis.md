# 3876. Construct Uniform Parity Array II

[LeetCode Link](https://leetcode.com/problems/construct-uniform-parity-array-ii/)

Difficulty: Medium
Topics: Array, Math
Acceptance Rate: 57.8%

## Hints

### Hint 1

The problem talks about "all odd or all even", so nothing about the actual magnitude of the numbers matters except through parity. Try splitting the decision into two completely independent questions: *can I make everything even?* and *can I make everything odd?* The answer is `true` if either sub-question is `true`. Also remember the only arithmetic fact you need: `even - even = even`, `odd - odd = even`, `even - odd = odd`, `odd - even = odd`.

### Hint 2

Fix a target parity and look at one index at a time. An element that already has the target parity can just be kept (`nums2[i] = nums1[i]`), so it is never a problem. An element with the *wrong* parity is forced to use the subtraction option, and the parity rules above pin down exactly the parity of the partner `nums1[j]` it needs. On top of parity, the constraint `nums1[i] - nums1[j] >= 1` means the partner has to be **strictly smaller** than it. So each "wrong parity" element needs a smaller element of one specific parity — and nothing stops many indices from reusing the same partner `j`.

### Hint 3

Apply Hint 2 to each target separately.

- **All even:** every odd element needs a strictly smaller *odd* partner. The smallest odd element in the array has no smaller odd element, so it can never be fixed. Therefore all-even is possible **only when there are no odd elements at all**.
- **All odd:** every even element needs a strictly smaller *odd* partner. A single odd element can serve every one of them, and the best candidate is the smallest odd value. So all-odd is possible exactly when at least one odd exists and `min(odd) < min(even)`.

That reduces the whole problem to two minimums in one pass — no sorting, no searching over pairs.

## Approach

Scan the array once and track two values:

- `minOdd` — the smallest odd element (absent if the array has no odd numbers)
- `minEven` — the smallest even element (absent if the array has no even numbers)

Then decide:

1. **No odd elements.** Every element is already even, so keep everything as-is: `nums2 = nums1` is all even. Return `true`.
2. **No even elements.** Every element is already odd, so keep everything as-is. Return `true`.
3. **Both parities present.** All-even is impossible (case 1's argument: the smallest odd has no smaller odd to subtract, and subtracting an even keeps it odd). So the only hope is all-odd. Each even element `x` must become `x - y` for some odd `y < x`; using `y = minOdd` is optimal, because if `minOdd` fails for some `x` then no odd value works for it. So the answer is `minOdd < minEven`: if the smallest odd beats the smallest even, it beats every even value, and every even element can be repaired. Otherwise the smallest even element is stuck.

Note that reuse is free — the problem lets any number of indices `i` pick the same partner `j` — which is why one well-chosen `minOdd` is enough for all evens at once. Also, the odd elements themselves need no work in the all-odd case; they simply keep their original value.

Worked example, `nums1 = [1, 4, 7]`: `minOdd = 1`, `minEven = 4`. Both parities appear, and `1 < 4`, so the answer is `true`. Concretely, `1` and `7` stay put, and `4` becomes `4 - 1 = 3`, giving `[1, 3, 7]`.

Counter-example, `nums1 = [2, 3]`: `minOdd = 3`, `minEven = 2`, and `3 < 2` is false. The element `2` has no smaller odd partner (`3` is larger, so `2 - 3 = -1` violates the `>= 1` rule), and `3` has no smaller odd partner either, so neither target parity is reachable — `false`.

The distinctness guarantee is a convenience rather than a load-bearing fact here; the comparison `minOdd < minEven` is between values of different parity, so they can never be equal anyway.

This is a Medium that is genuinely easy *once* you make the "solve each target parity separately" split. The trap is trying to reason about both targets at once, or trying to match each element to a partner greedily — that leads to needless sorting or pairing logic.

## Complexity Analysis

Time Complexity: O(n) — a single pass over `nums1`.
Space Complexity: O(1) — only the two running minimums.

## Edge Cases

- **`n == 1`.** The subtraction option requires an index `j != i`, so it is unavailable; but keeping the single element always yields a uniform array. The code returns `true` because one of the two "only one parity present" branches fires.
- **All elements even.** Return `true` immediately by keeping everything. Do not fall into the `minOdd < minEven` comparison with a missing `minOdd`.
- **All elements odd.** Same idea in the other direction — `true`, with no `minEven` to compare against. Both of these are why the absence of a parity must be tracked with a sentinel rather than assuming both minimums exist.
- **Smallest even below the smallest odd** (e.g. `[2, 3]`, or `[2, 5, 7, 9]`). The smallest even can never be repaired, so the answer is `false` even though every other even element might have been fixable.
- **Smallest odd immediately below the smallest even** (e.g. `[1, 2]`). The strict inequality is what matters; since parities differ the values can never tie, so `<` and `<=` behave identically, but `<` states the intent correctly.
- **Large values (up to 1e9) and large `n` (up to 1e5).** Values fit comfortably in Go's `int`, and the single pass avoids any O(n log n) or O(n^2) work.

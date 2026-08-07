# 3348. Smallest Divisible Digit Product II

[LeetCode Link](https://leetcode.com/problems/smallest-divisible-digit-product-ii/)

Difficulty: Hard
Topics: Math, String, Backtracking, Greedy, Number Theory
Acceptance Rate: 23.9%

This one earns its Hard rating: the greedy is short, but every piece of it needs
a small proof before you can trust it. Take the hints one at a time.

## Hints

### Hint 1

`num` can have 200,000 digits, so the answer is a string you *construct*, never a
number you search for. Stop thinking about "the next number after `num`" and
start thinking about what a digit actually contributes.

The answer is zero-free, so it uses only the digits 1 through 9. Write out the
prime factorization of each of those nine digits. What primes can ever appear in
a digit product? What does that immediately tell you about most values of `t`?

### Hint 2

Every digit factors into 2s, 3s, 5s and 7s only, so `t` is relevant only through
its exponent vector `(e2, e3, e5, e7)` — and if `t` has any other prime factor,
the answer is `"-1"` with no work at all.

That reframes the whole problem. As you consume digits left to right, each digit
knocks down some of those four exponents (independently — a factor of 5 never
helps with 2s). So the state after any prefix is just "what exponents are still
owed". Now the question becomes: given `k` free positions at the end and a
leftover requirement, can `k` digits possibly cover it?

If you can answer that question in O(1), what classic "next larger string" shape
does the rest of the problem collapse into?

### Hint 3

Build the oracle `minDigits(e2, e3, e5, e7)` = the fewest digits in 1..9 whose
product covers that requirement.

- Only the digit 5 supplies a factor 5, and only 7 supplies a factor 7. So those
  exponents cost exactly `e5 + e7` digits, no cleverness available.
- For 2s and 3s the densest carriers are 8 (`2^3`) and 9 (`3^2`), giving
  `ceil(e2/3) + ceil(e3/2)`. The one improvement is the digit 6, which carries
  one 2 *and* one 3: it saves a digit exactly when the leftovers after that
  packing are one 2 and one 3, i.e. when `e2 % 3 == 1 && e3 % 2 == 1`. Using a
  second 6 never helps — after the first swap both remainders are back to zero,
  so a further 6 costs a full digit and saves nothing.

With that O(1) feasibility test, the answer is the standard "smallest string ≥
`num`" greedy: keep the longest possible prefix of `num`, bump one position up,
and fill the tail with the smallest digits that still satisfy the oracle.

## Approach

**Step 0 — factor `t`.** Strip out 2s, 3s, 5s and 7s into `(e2, e3, e5, e7)`. If
anything is left over, no digit product can ever contain that prime: return
`"-1"`. This is the only `"-1"` case — if `t` is 7-smooth, some (possibly very
long) number always works.

**Step 1 — prefix requirements.** Precompute `rem[i]`, the exponent vector still
owed after consuming `num[:i]`. Each prime is tracked separately and clamped at
zero, so `rem[i+1] = rem[i]` minus the factorization of `num[i]`, floored at 0 per
prime. Clamping per prime is exact precisely because divisibility by
`2^a·3^b·5^c·7^d` decomposes into four independent conditions.

**Step 2 — is `num` itself the answer?** If `num` has no `'0'` and `rem[n]` is
all zeros, return `num`.

**Step 3 — an answer of the same length.** Otherwise the answer of length `n`
looks like `num[:i] + d + tail`, where `d > num[i]` and `tail` has `n-1-i` free
digits. Two constraints:

- The kept prefix must itself be zero-free, so `i <= firstZero`, the index of the
  first `'0'` in `num` (or `n` if there is none).
- Feasibility: `minDigits(rem[i] / d) <= n-1-i`.

Scan `i` downward from `min(firstZero, n-1)` and, for each `i`, try `d` upward
from `num[i]+1`. The first hit is the answer.

*Why largest `i` wins:* two same-length candidates that break at `i1 > i2` agree
with `num` (and each other) below `i2`; at position `i2` the first still equals
`num[i2]` while the second is strictly greater. So the later break is smaller.

*Why smallest `d` wins at a fixed `i`:* the prefix is identical, so position `i`
decides the comparison, and the tail is minimized independently afterwards. Note
you must genuinely test each `d` in order — feasibility is **not** monotone in
`d` (for a requirement of 3s, `d = 6` works where `d = 5` does not).

**Step 4 — no same-length answer.** Then grow. The shortest usable length is
`max(n+1, minDigits(t))`, and any zero-free string of that length already beats
every length-`n` number, so just fill it greedily from scratch.

**Filling a tail of `k` digits.** Walk left to right; at each position pick the
smallest digit `d` in 1..9 such that `minDigits(remaining / d) <= ` the number of
positions still to its right. This front-loads 1s and pushes the heavy digits to
the end, which is exactly what "lexicographically smallest" wants. It can never
get stuck: if `m` digits suffice for the current requirement, then removing any
one member of an optimal size-`m` multiset leaves something coverable by `m-1`.

**Worked example — `num = "1234"`, `t = 256 = 2^8`.**

`rem = [(8), (8), (7), (7), (5)]` (2-exponents; the `2` and `4` in `num` pay off
one and two 2s). `rem[4]` is nonzero, so `num` itself fails.

- `i = 3`: 0 free slots, best digit `8` leaves `2^4` → needs 2 more. Fail.
- `i = 2`: 1 free slot, `d ∈ 4..9`; best is `8`, leaving `2^4` → needs 2. Fail.
- `i = 1`: 2 free slots, `rem[1] = 2^8`. `d = 3` leaves `2^8` (3 digits, too
  many); `d = 4` leaves `2^6`, which needs exactly 2 digits. Hit.

Tail of 2 covering `2^6`: first position wants the smallest `d` leaving something
one digit can cover — only `8` works (`2^3` remains), then `8` again. Answer:
`"1" + "4" + "88" = "1488"`. ✓

## Complexity Analysis

Time Complexity: O(n + log t), where n = len(num). Factoring `t` is O(log t); the
prefix scan, the outer break-position scan (≤ 9 O(1) feasibility checks each),
and the tail fill (≤ 9 checks per position) are each O(n).

Space Complexity: O(n) for the `rem` array and the output buffer. The exponent
vector itself is O(1) — with `t ≤ 10^14` the exponents are bounded by 46, and
`minDigits(t)` never exceeds ~20 digits (worst case `5^20`).

## Edge Cases

- **`t` not 7-smooth** (e.g. `t = 26`): the only `"-1"` case, and it must be
  checked first. Any 7-smooth `t` always has an answer, since you can always make
  the number longer.
- **`t = 1`**: the requirement is empty from the start. The problem degenerates to
  "smallest zero-free number ≥ `num`", and the same code path handles it —
  `"10"` → `"11"`, while `"99"` → `"99"`.
- **`num` already qualifies**: needs `num` zero-free *and* the product divisible,
  and the check must come before any bumping. Missing it returns something
  strictly larger than the true answer.
- **Zeros inside `num`**: the kept prefix must stop at or before the first `'0'`.
  Forgetting this is the classic bug — you cannot keep `"10"` from `"1023"`, so
  the break position is capped at index 1.
- **All-nines positions**: `num[i] == '9'` offers no digit to bump to, so those
  positions are silently skipped; `"99"` with `t = 5` must grow to `"115"`.
- **Growth driven by `t`, not by `n`**: with `num = "12"` and `t = 10^14`, the
  answer is 19 digits long. The new length is `max(n+1, minDigits(t))`, not just
  `n+1`.
- **Feasibility is not monotone in the bumped digit**: iterate `d` upward and
  test each one; do not break early assuming a larger digit is always at least as
  useful.
- **Length**: `n` up to 2·10^5 rules out anything that materializes candidate
  numbers or does per-candidate work; `t` up to 10^14 requires `int64`.

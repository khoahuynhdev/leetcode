# 3622. Check Divisibility by Digit Sum and Product

[LeetCode Link](https://leetcode.com/problems/check-divisibility-by-digit-sum-and-product/)

Difficulty: Easy
Topics: Math
Acceptance Rate: 70.3%

## Hints

### Hint 1

This is a pure digit-manipulation problem — no clever data structure is needed. The only real question is: how do you get at the individual digits of an integer? There are two standard ways, and one of them avoids string conversion entirely.

### Hint 2

The classic technique is repeated division: `n % 10` peels off the last digit, and `n / 10` (integer division) removes it. Loop until the number becomes zero, and you have visited every digit exactly once. You need two accumulators — one starting at `0` for a running sum, one starting at `1` for a running product.

### Hint 3

Two details make this click. First, **do not mutate the original `n`** while extracting digits — you still need it for the final modulo check, so iterate over a copy. Second, worry about division by zero: the divisor is `digitSum + digitProduct`. The product *can* be zero (any `n` containing a `0` digit kills it), but since `n >= 1` the digit sum is always at least `1`, so the total is never zero. The check `n % (sum + product) == 0` is always safe.

## Approach

Walk the digits once with the standard extract-and-shrink loop, accumulating both quantities in the same pass.

1. Initialize `digitSum = 0` and `digitProduct = 1`. The identity element differs: `0` for addition, `1` for multiplication.
2. Copy `n` into a scratch variable `rest`. This matters — the original value is needed for the divisibility test at the end.
3. While `rest > 0`: take `d := rest % 10`, add it to `digitSum`, multiply it into `digitProduct`, then set `rest /= 10`.
4. Return `n % (digitSum + digitProduct) == 0`.

Tracing example 1 with `n = 99`: `rest = 99` gives `d = 9`, so `digitSum = 9`, `digitProduct = 9`, `rest = 9`. Next iteration gives `d = 9` again, so `digitSum = 18`, `digitProduct = 81`, `rest = 0`, and the loop ends. The divisor is `18 + 81 = 99`, and `99 % 99 == 0`, so the answer is `true`.

Tracing example 2 with `n = 23`: the digits are `3` then `2`, yielding `digitSum = 5` and `digitProduct = 6`. The divisor is `11`, and `23 % 11 == 1`, so the answer is `false`.

The digit-extraction loop hands you digits in reverse order (least significant first), but that is irrelevant here: both addition and multiplication are commutative, so order never affects the result.

Overflow is a non-issue at these constraints. With `n <= 10^6` there are at most 7 digits, so the product is bounded by `9^7 = 4,782,969` — comfortably inside a 32-bit `int`, let alone Go's 64-bit `int` on typical platforms.

A string-based alternative (`strconv.Itoa` then iterate the bytes) is equally correct and arguably reads a little more directly, but it allocates and forces a `byte - '0'` conversion per character. The arithmetic version is faster and just as clear once the `% 10` / `/ 10` idiom is familiar — and that idiom shows up constantly in digit problems, so it is worth having in muscle memory.

Honest difficulty note: this one really is Easy, and the ~70% acceptance rate reflects that. The value here is not the puzzle but the reflex — one clean pass, correct identity elements, and a habit of asking "can this divisor be zero?" before writing a `%`.

## Complexity Analysis

Time Complexity: O(log n) — the loop runs once per digit, and a number `n` has `floor(log10(n)) + 1` digits. Under the stated constraint that is at most 7 iterations, so it is effectively constant.
Space Complexity: O(1) — only a fixed number of integer accumulators, regardless of input size.

## Edge Cases

- **Single-digit `n` (1 through 9)**: both the sum and the product equal `n` itself, so the divisor is `2n`. Since `n % 2n == n` and `n >= 1`, every single-digit input returns `false`. A tempting wrong answer is `true` for these, so it is worth testing at least one.
- **`n` containing a zero digit** (10, 100, 1000000): the product collapses to `0`, leaving the digit sum as the sole divisor. For `n = 10` the divisor is just `1`, so the answer is `true`. Initializing the product to `0` instead of `1` breaks every other case, and this family is where that bug hides.
- **Divisor can never be zero**: because `n >= 1`, the digit sum is at least `1`, so `digitSum + digitProduct >= 1`. No guard is required before the modulo — but you should confirm this reasoning rather than assume it.
- **Preserving the original `n`**: if you decrement `n` directly in the digit loop, it is `0` by the time you reach the divisibility check, and `0 % k == 0` makes the function return `true` for every input. This is the most common bug in this problem.
- **Upper bound `n = 1000000`**: the largest legal input and a 7-digit number containing zeros — a good smoke test that the loop terminates correctly and the product path behaves.

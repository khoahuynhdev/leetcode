# 3612. Process String with Special Operations I

[LeetCode Link](https://leetcode.com/problems/process-string-with-special-operations-i/)

Difficulty: Medium
Topics: String, Simulation

Acceptance Rate: 68.3%

## Hints

### Hint 1

This problem is asking you to *replay* a sequence of instructions one by one and watch how a string evolves. When a problem describes a list of rules applied "from left to right," it is usually a **simulation** problem — you don't need a clever formula, you need to faithfully model what each character does.

### Hint 2

Think about which data structure makes each of the four operations cheap and easy to express:

- appending a letter,
- removing the last character,
- duplicating the whole string,
- reversing the whole string.

A growable sequence of bytes (a slice/`[]byte` or string builder) handles all of these naturally. The only operation you must be careful with is removal, since the string might be empty.

### Hint 3

The whole solution is a single pass over `s` with a `switch` on the current character. The key insight is that the constraints are tiny (`s.length <= 20`), so you can afford to literally perform each operation, even the ones that look expensive like duplicate and reverse. There is no hidden trick — correctness and careful edge-case handling (empty `result` on `*`) are what matter.

## Approach

Maintain a mutable buffer (`result`) that starts empty. Walk through `s` left to right and apply each character:

1. **Lowercase letter** — append it to `result`.
2. **`'*'`** — remove the last character of `result`, *but only if `result` is non-empty*. Removing from an empty buffer is a no-op.
3. **`'#'`** — duplicate `result` by appending a copy of it to itself (`result = result + result`).
4. **`'%'`** — reverse `result` in place.

After the loop, `result` is the answer.

Walking through Example 1, `s = "a#b%*"`:

| Step | Char | Operation | `result` |
|------|------|-----------|----------|
| 0 | `a` | append | `"a"` |
| 1 | `#` | duplicate | `"aa"` |
| 2 | `b` | append | `"aab"` |
| 3 | `%` | reverse | `"baa"` |
| 4 | `*` | remove last | `"ba"` |

Final answer: `"ba"`.

Using a `[]byte` buffer keeps every operation simple: append is `append(...)`, remove is a length truncation `result = result[:len(result)-1]`, duplicate is `append(result, result...)`, and reverse is a standard two-pointer swap. Each maps directly to the rule it implements, which keeps the code readable and easy to verify.

## Complexity Analysis

Let `n = len(s)`. The buffer can grow when `#` duplicates it. In the worst case the length doubles on a `#`, so the final length is bounded by `O(2^n)`, but with `n <= 20` this is a small, fixed bound.

Time Complexity: O(n + L) where `L` is the total work done across duplicate/reverse operations on the buffer. Each `#` and `%` touches at most the current buffer length, so the total is proportional to the sum of buffer sizes seen — bounded by the final length. With the given constraints this is effectively constant-bounded and very fast.

Space Complexity: O(L) for the buffer holding the result, where `L` is the final string length.

## Edge Cases

- **`'*'` on an empty `result`**: Must be a no-op. Forgetting the empty check causes an index-out-of-range panic (slicing `result[:len(result)-1]` when `len == 0`). See Example 2 where `z*` empties the buffer before further ops.
- **`'#'` on an empty `result`**: Duplicating `""` yields `""` — naturally handled, but worth confirming your code doesn't special-case it incorrectly.
- **`'%'` on an empty or single-character `result`**: Reversing should leave it unchanged; a correct two-pointer loop handles this automatically.
- **A buffer that becomes empty mid-processing** then receives more letters: ensure appends still work after the buffer was emptied.
- **Strings consisting only of special characters** (e.g. `"*%#"`): the result stays empty throughout.

# 0835. Image Overlap

[LeetCode Link](https://leetcode.com/problems/image-overlap/)

Difficulty: Medium
Topics: Array, Matrix
Acceptance Rate: 65.5%

## Hints

### Hint 1

The problem hands you a search space that is smaller than it first appears. You are told you may slide the image "any number of units" in any direction, but a shift that pushes every `1` off the board is useless. Start by asking: how many *distinct* translations can actually produce a non-zero overlap? Once you can enumerate the candidates, brute force over them becomes a legitimate strategy — the interesting work is in deciding what to enumerate and how to score each candidate cheaply.

### Hint 2

There are two ways to attack this, and they lead to different complexities.

The first is to enumerate translations directly: for each vertical shift `dr` and horizontal shift `dc` in the range `[-(n-1), n-1]`, count the positions where the shifted `img1` and `img2` both hold a `1`. That is `O(n^2)` shifts, each costing `O(n^2)` to score.

The second flips the loop inside out. Instead of asking "for this shift, how many `1`s line up?", ask "for this *pair* of `1`s, which shift would align them?" Think about what a single `1` at `(i, j)` in `img1` and a single `1` at `(k, l)` in `img2` tell you about a translation.

### Hint 3

A `1` at `img1[i][j]` lands on top of a `1` at `img2[k][l]` under exactly **one** translation: the offset `(k - i, l - j)`. So every ordered pair of ones *votes* for precisely one shift.

That means you can collect the coordinates of all `1`s in each image, take every pair across the two lists, compute the offset, and tally the votes in a hash map. The offset with the most votes is the best translation, and its vote count *is* the overlap — because the overlap under a given shift is by definition the number of ones that align under it, which is exactly the number of pairs that voted for it. No shifting of matrices required; you never move a single bit.

## Approach

The key reframing is to stop thinking about sliding matrices and start thinking about **pairs of ones voting for offsets**.

**Why the vote count equals the overlap.** Fix a translation `(dr, dc)` meaning "move `img1` down by `dr` and right by `dc`." Under it, the `1` at `img1[i][j]` moves to `(i + dr, j + dc)`. It contributes to the overlap exactly when `img2[i + dr][j + dc] == 1`. Setting `k = i + dr` and `l = j + dc`, that condition is "there is a one at `img2[k][l]`," and the offset is recovered as `dr = k - i`, `dc = l - j`. So the ones contributing to the overlap at `(dr, dc)` are in bijection with the pairs `((i, j), (k, l))` whose coordinate difference is `(dr, dc)`. Counting pairs by their difference and taking the maximum therefore gives the maximum overlap directly.

Note this also handles the "bits translated outside the borders are erased" rule for free. A one that slides off the grid has no partner in `img2` to pair with, so it simply never casts a vote. There is no clamping or bounds arithmetic to get wrong.

**Algorithm.**

1. Scan `img1` once and collect the coordinates of every `1` into a slice `ones1`. Do the same for `img2` into `ones2`.
2. If either slice is empty, return `0` immediately — no alignment is possible.
3. For each `a` in `ones1` and each `b` in `ones2`, compute the offset `(b.row - a.row, b.col - a.col)` and increment `count[offset]` in a map.
4. Return the largest value in `count`.

Because the offsets live in `[-(n-1), n-1]` on both axes, the key can be packed into a single `int` — for instance `(dr + n) * (2*n) + (dc + n)` — which keeps the map keys compact and avoids hashing a struct. Either works; the packed integer is a little faster and is what the solution uses.

**Walking Example 1.** With

```
img1 = [[1,1,0],    img2 = [[0,0,0],
        [0,1,0],            [0,1,1],
        [0,1,0]]            [0,0,1]]
```

the ones are `ones1 = [(0,0), (0,1), (1,1), (2,1)]` and `ones2 = [(1,1), (1,2), (2,2)]`.

Consider the votes cast for offset `(1, 1)`, i.e. down 1 and right 1:

- `(0,0)` vs `(1,1)` → difference `(1,1)` ✓
- `(0,1)` vs `(1,2)` → difference `(1,1)` ✓
- `(1,1)` vs `(2,2)` → difference `(1,1)` ✓
- `(2,1)` would need a partner at `(3,2)`, which is off the board, so it contributes nothing.

Offset `(1,1)` receives 3 votes, and no other offset beats it, so the answer is `3` — matching the problem's explanation of translating right 1 and down 1.

**On complexity and why this is the right trade.** The naive shift-enumeration approach is `O(n^4)`, which at `n = 30` is 810,000 operations — perfectly acceptable here. The pair-voting approach is `O(m1 * m2)` where `m1` and `m2` are the counts of ones. In the worst case (both images entirely ones) that is also `O(n^4)`, so asymptotically the two tie. The practical win is that voting scales with the *density of ones* rather than the matrix size, so sparse images — the common case — finish far faster. It is worth being honest that this problem's small constraint (`n <= 30`) means brute force passes comfortably; the pair-voting insight is the reason the problem is interesting, not a necessity for getting accepted.

## Complexity Analysis

Time Complexity: O(m1 · m2), where `m1` and `m2` are the numbers of `1`s in `img1` and `img2`. This is bounded by O(n⁴) in the worst case (both matrices entirely ones) and drops to near O(n²) for sparse inputs, since the initial scan of both matrices costs O(n²) regardless.

Space Complexity: O(m1 + m2) to store the coordinate lists, plus O(min(m1 · m2, n²)) for the offset-count map — there are at most `(2n - 1)²` distinct offsets, so the map is O(n²) bounded. Overall O(n²).

## Edge Cases

- **Either image is all zeros.** With no ones in one of the images there are no pairs, so no votes are cast and the map stays empty. Returning the max over an empty map is a classic crash or wrong-answer source in Go (the zero value of the running max works here, but only if you initialize it to `0` rather than to something like the first map entry). The solution guards with an early return of `0`.
- **Both images are all zeros.** Example 3 in the problem. Same reasoning as above; the answer is `0`, not `n²`. Zeros aligning with zeros does not count as overlap.
- **`n == 1`.** The smallest legal input. Only the zero offset `(0, 0)` is possible; the answer is `1` if both cells are `1` (Example 2) and `0` otherwise.
- **Negative offsets.** Translations go up and left as well as down and right, so `dr` and `dc` range over `[-(n-1), n-1]`. If you pack the offset into an integer key, you must bias by `+n` (or similar) before packing, or negative components will collide with positive ones and silently inflate counts.
- **No translation at all is best.** The identity shift `(0, 0)` is a legitimate candidate and must be included in the search — pairs where `a == b` naturally produce offset `(0, 0)`, so the voting approach covers it without a special case.
- **Both images entirely ones.** The dense worst case, `900 × 900 = 810,000` pairs at `n = 30`. Still fast, but it is the input to keep in mind when reasoning about the runtime. The answer is `n²` via the zero offset.
- **Asymmetry of the roles.** The offset is `img2`-coordinate minus `img1`-coordinate, and the problem is symmetric (translating `img1` right is equivalent to translating `img2` left), so consistency matters more than direction — just do not mix the subtraction order between the row and column components.

# 3655. XOR After Range Multiplication Queries II

[LeetCode Link](https://leetcode.com/problems/xor-after-range-multiplication-queries-ii/)

Difficulty: Hard
Topics: Array, Divide and Conquer
Acceptance Rate: 40.5%

## Hints

### Hint 1

Each query multiplies certain indices by a value. Since multiplication mod a prime is commutative and associative, the order of multiplications doesn't matter — only the total product applied to each index matters. Think about how to efficiently accumulate the product of multipliers per index.

### Hint 2

Consider the step size `k` in each query. When `k` is large, very few elements are affected per query. When `k` is small, many elements are affected but they follow a regular stride pattern. Can you treat these two cases differently? Think **sqrt decomposition** on the step size.

### Hint 3

For small `k` (≤ √n), view the array as `k` interleaved strands (indices 0, k, 2k, ... form one strand; 1, 1+k, 1+2k, ... form another). A query with step `k` updates a contiguous range within a single strand. You can use a **multiplicative difference array** on each strand: multiply the start position by `v` and multiply the position after the end by `v⁻¹` (modular inverse). Then a prefix-product sweep applies all updates at once.

## Approach

### Key Insight: Sqrt Decomposition on Step Size

Since multiplication modulo a prime is commutative, we only need the total product of all multipliers applied to each index. We split queries into two groups based on a threshold `B = ⌈√n⌉`:

**Large step (k > B):** Each such query touches at most `⌈n/k⌉ < √n` elements. Apply these directly by iterating through the affected indices. Total work across all large-step queries: `O(q · √n)`.

**Small step (k ≤ B):** For a fixed step `k`, the array splits into `k` strands where strand `j` contains indices `j, j+k, j+2k, ...`. A query `[l, r, k, v]` affects a contiguous segment within strand `l % k`.

We use a **multiplicative difference array**: for each small `k`, maintain an array `diff` of size `n` initialized to `1`. For a query affecting positions from `l` to `last` (where `last = l + ⌊(r−l)/k⌋·k`):
- Set `diff[l] *= v`
- If `last + k < n`, set `diff[last + k] *= v⁻¹ mod p`

After processing all queries for a given `k`, sweep each strand computing prefix products. This applies all accumulated multipliers in one pass.

### Algorithm

1. Compute `B = ⌈√n⌉` and precompute modular inverses for values `1..10⁵`.
2. For large-step queries: multiply `totalMult[idx]` directly for each affected index.
3. For small-step queries grouped by `k`: build multiplicative diff arrays, then sweep per strand.
4. Compute `result = XOR of (nums[i] · totalMult[i] mod p)` for all `i`.

### Modular Inverse

Since `10⁹+7` is prime, use the batch formula: `inv[i] = (p − ⌊p/i⌋ · inv[p mod i] mod p) mod p` to precompute all needed inverses in O(maxV) time.

## Complexity Analysis

Time Complexity: O((n + q) · √n)
- Large-step queries: O(q · √n) total
- Small-step queries: O(q) for diff updates + O(B · n) = O(n · √n) for sweeps
- Inverse precomputation: O(maxV) = O(10⁵)

Space Complexity: O(n + maxV)
- `totalMult` array: O(n)
- `diff` array (reused): O(n)
- Inverse table: O(maxV)

## Edge Cases

- **Step k larger than range (r − l):** Only the starting index `l` is affected. The diff undo position falls outside the array, so no inverse is written.
- **Multiplying by 1:** No change to values — the algorithm handles this naturally since multiplying by 1 is a no-op in the product.
- **Single element array:** Only one index exists. All queries collapse to direct multiplications on index 0.
- **Multiple queries on the same index:** Products accumulate correctly due to commutativity of modular multiplication.
- **Large intermediate products:** All arithmetic uses `int64` to avoid overflow. The maximum intermediate value `(10⁹+6)² ≈ 10¹⁸` fits within `int64` range.

# 3629. Minimum Jumps to Reach End via Prime Teleportation

[LeetCode Link](https://leetcode.com/problems/minimum-jumps-to-reach-end-via-prime-teleportation/)

Difficulty: Medium
Topics: Array, Hash Table, Math, Breadth-First Search, Number Theory
Acceptance Rate: 35.1%

## Hints

### Hint 1

You are asked for the **minimum number of jumps** between two nodes (indices `0` and `n-1`) in an unweighted graph where each move costs exactly one jump. That is the textbook signature for a Breadth-First Search. The hard part is not the BFS itself — it is figuring out what the graph looks like and how to keep it small.

### Hint 2

The graph has obvious adjacent edges (`i ↔ i+1`), but the prime-teleportation rule is dangerous: if many positions share a prime factor with many other positions, naively wiring up "every prime index `i` to every divisible index `j`" can blow up to `O(n²)` edges. A common BFS trick when groups of nodes are all mutually connected through some shared key is the **virtual / super node**: instead of `k²` edges, you connect the `k` nodes through one shared key node and pay only `O(k)` edges. Here the natural shared key is the prime `p` itself.

Equivalently — and a bit easier to code — bucket every index `j` under each prime `p` that divides `nums[j]`, and during BFS, the *first* time you stand on an index whose value is the prime `p`, expand the entire bucket at once and then never expand `p` again.

### Hint 3

Two further observations make this efficient:

1. You only ever teleport *from* an index `i` whose value is itself a prime `p`. So the only primes that matter are those that actually appear as a value in `nums`. Restrict your prime-to-indices buckets to those primes.
2. Each `nums[j] ≤ 10⁶` has at most about 7 distinct prime factors (since `2·3·5·7·11·13·17·19 > 10⁶`). That bounds the total bucket size to `O(n · log(maxVal))`, well within the limits.

Combine the above: precompute smallest prime factors with a linear sieve, factorize each `nums[j]`, build the buckets, then run a level-order BFS with a "prime already expanded" flag set so each prime's bucket is dumped into the queue at most once.

## Approach

We model the problem as an unweighted graph and run BFS from index `0`. Each node is an index `0..n-1`. The edges are:

- **Adjacent edges**: index `i` connects to `i-1` and `i+1` if in range.
- **Prime-teleport edges**: if `nums[i]` is a prime `p`, then `i` connects to every index `j` with `p | nums[j]`.

To avoid a quadratic edge count, we use the standard trick of expanding teleport groups lazily during BFS:

1. **Sieve of smallest prime factors** up to `max(nums)`. With `spf[v] = v` iff `v` is prime, primality and factorization both become `O(log v)` per query.
2. Walk `nums` once to collect the set `primeValues` of primes that appear as actual array values. Only those primes can ever be the source of a teleport.
3. For each index `j`, factorize `nums[j]` and append `j` to `primeToIdx[p]` for every prime factor `p ∈ primeValues`. This builds the "bucket per prime".
4. Run a level-order BFS from `0`. For each popped index `idx`:
   - Push the unvisited neighbors `idx-1` and `idx+1`.
   - If `nums[idx]` is a prime that has not yet been expanded, push every unvisited `j ∈ primeToIdx[nums[idx]]` and mark that prime as expanded so it cannot be expanded again.
5. The first time we pop `n-1`, the current level is the answer.

The "expand each prime at most once" rule is what keeps the algorithm linear-ish: every index can only be added to the queue by the prime expansion of one of its prime factors, so the total work across all teleport expansions is bounded by the total bucket size.

**Walking through Example 1** — `nums = [1, 2, 4, 6]`:

- Primes appearing as values: `{2}`. Bucket `primeToIdx[2] = [1, 2, 3]` (since `2`, `4`, `6` are all divisible by 2).
- Level 0: pop `0`. Push `1` (adjacent). `nums[0] = 1` is not prime.
- Level 1: pop `1`. Push `2` (adjacent). `nums[1] = 2` is prime — expand bucket `[1, 2, 3]`: `1` and `2` are visited, push `3`.
- Level 2: pop `2` (not target), pop `3` (target). Answer: `2`.

## Complexity Analysis

Let `n = len(nums)` and `M = max(nums) ≤ 10⁶`.

Time Complexity: `O(M · log log M + n · log M)`
- Smallest-prime-factor sieve: `O(M · log log M)`.
- Factoring each `nums[j]` and bucketing: `O(log nums[j])` per element, so `O(n · log M)` total.
- BFS: each index is enqueued at most once and each prime bucket is expanded at most once, so the work is bounded by the total bucket size `O(n · log M)` plus `O(n)` for adjacent edges.

Space Complexity: `O(M + n · log M)`
- The SPF sieve uses `O(M)`.
- The buckets in total hold `O(n · log M)` indices.
- Visited arrays and BFS queue add `O(n)`.

## Edge Cases

- **`n == 1`**: We start at the destination, so the answer is `0` jumps. Handle this before any BFS work.
- **`nums[i] == 1`**: `1` is not prime, so it never enables a teleport. The sieve must treat `spf[1] = 0` (or otherwise reject 1) so `isPrime(1)` returns false.
- **Same prime appears at multiple indices**: e.g. `[2, 2, 2, ...]`. The "expand each prime once" guard is what prevents quadratic behavior; without it, every index with value `2` would re-enqueue every even index.
- **A prime `p` appears as a value but has no other multiples in `nums`** (e.g. `[7, 1, 1]`): `primeToIdx[7]` will only contain the index of `7` itself, so the teleport step is a no-op and BFS falls back to adjacent steps. This is correct.
- **No primes in `nums` at all** (e.g. `[1, 4, 6, 8]` — all composite or `1`): `primeValues` is empty, no teleports are ever possible, BFS reduces to a straight walk and returns `n - 1`.
- **Largest prime hits the destination directly from index 0**: e.g. `nums[0]` is prime `p` and `p | nums[n-1]`. The answer is `1`, recovered on the very first BFS expansion.

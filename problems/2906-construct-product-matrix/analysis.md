# 2906. Construct Product Matrix

[LeetCode Link](https://leetcode.com/problems/construct-product-matrix/)

Difficulty: Medium
Topics: Array, Matrix, Prefix Sum
Acceptance Rate: 37.8%

## Hints

### Hint 1

This problem is very similar to the classic "Product of Array Except Self" problem. Think about how you would solve that 1D problem, then consider how to adapt it to a 2D matrix.

### Hint 2

You can flatten the 2D matrix into a 1D sequence and apply a prefix/suffix product technique. For each position, the answer is the product of everything before it multiplied by the product of everything after it — all taken modulo 12345.

### Hint 3

The critical insight is that you only need two passes over the flattened array: one forward pass to build prefix products and one backward pass to build suffix products. At each position, multiply the prefix product (of all elements before it) by the suffix product (of all elements after it), modulo 12345. This avoids computing the total product and dividing (which wouldn't work well with modular arithmetic anyway).

## Approach

1. **Flatten the matrix conceptually**: Treat the `n x m` grid as a 1D array of length `n * m` by iterating row by row.

2. **Forward pass (prefix products)**: Traverse the flattened array left to right. Maintain a running product. Before updating the running product with the current element, store the running product as the prefix value for the current position. This gives, for each cell, the product of all elements that come before it (in row-major order), modulo 12345.

3. **Backward pass (suffix products)**: Traverse right to left. Maintain a running product. For each position, multiply the already-stored prefix value by the current suffix running product. Then update the suffix running product with the current grid element.

4. **Result**: After both passes, each cell in the result matrix contains `(product of all other elements) % 12345`.

This is essentially the same algorithm as "Product of Array Except Self" (LC 238), extended to a 2D grid and with modular arithmetic.

## Complexity Analysis

Time Complexity: O(n * m) — two linear passes over all elements.
Space Complexity: O(1) extra beyond the output matrix — prefix and suffix products are computed using a single running variable each, and the result is stored in-place in the output matrix.

## Edge Cases

- **Element equals 12345 (or a multiple)**: That element is 0 mod 12345, which means every other cell's product includes a zero factor. The prefix/suffix approach handles this naturally.
- **Multiple multiples of 12345**: If two or more elements are 0 mod 12345, then every cell's product will be 0 (since every cell excludes at most one zero, but at least one zero remains).
- **Single row or single column matrix**: The algorithm works the same since we treat the matrix as a flat array.
- **Minimum size (n * m = 2)**: Each cell is simply the other cell's value mod 12345.
- **Large values**: Individual grid values can be up to 10^9, so intermediate products must use modular arithmetic at every step to avoid overflow.

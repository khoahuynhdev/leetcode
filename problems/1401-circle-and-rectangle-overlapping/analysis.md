# 1401. Circle and Rectangle Overlapping

[LeetCode Link](https://leetcode.com/problems/circle-and-rectangle-overlapping/)

Difficulty: Medium
Topics: Math, Geometry
Acceptance Rate: 56.2%

## Hints

### Hint 1

This is tagged Medium, but the hard part is not the algorithm — it is asking the
right question. There is no data structure to pick and no loop to write. Resist
the urge to enumerate candidate points (corners, edge midpoints, lattice points
inside the rectangle): any such enumeration either misses cases or needs a proof
you will not want to write in an interview. Instead, restate the problem as a
distance question about *one* well-chosen point.

### Hint 2

Two shapes overlap exactly when the **shortest distance** between the circle's
center and the rectangle is at most `radius`. So the whole problem reduces to:
given a point and an axis-aligned rectangle, which point of the rectangle is
closest to that point? Because the rectangle is axis-aligned, the x and y
dimensions are completely independent — you can solve each axis separately and
then combine.

### Hint 3

On the x-axis alone, the closest x-coordinate inside `[x1, x2]` to `xCenter` is
just `xCenter` **clamped** into that interval: `min(max(xCenter, x1), x2)`. Do
the same for y. The resulting point `(cx, cy)` is the closest point of the whole
rectangle to the center — and it is automatically correct whether the center sits
inside the rectangle, beside an edge, or diagonally past a corner. Then answer
`(xCenter-cx)² + (yCenter-cy)² <= radius²`, comparing squares so you never touch
floating point.

## Approach

The key reframing: "do the circle and rectangle share a point?" is equivalent to
"is the closest point of the rectangle to the circle's center within `radius` of
that center?"

Why the equivalence holds: the rectangle is a closed convex set. Let `P` be the
point of the rectangle minimizing distance to the center `C`.

- If `dist(C, P) <= radius`, then `P` lies in the disk and in the rectangle, so
  they overlap.
- If `dist(C, P) > radius`, then *every* point of the rectangle is farther than
  `radius` from `C` (since `P` is the minimizer), so no point of the rectangle is
  in the disk and they cannot overlap.

Finding `P` is where the axis-aligned assumption pays off. The rectangle is the
Cartesian product `[x1, x2] × [y1, y2]`, so squared distance separates:

```
dist²(C, (x, y)) = (xCenter - x)² + (yCenter - y)²
```

The two terms depend on different variables and both are non-negative, so we can
minimize each independently. Minimizing `(xCenter - x)²` over `x ∈ [x1, x2]` is
the classic clamp:

```
cx = clamp(xCenter, x1, x2) = min(max(xCenter, x1), x2)
```

and likewise `cy = clamp(yCenter, y1, y2)`. The clamp quietly handles all nine
relative positions of the center at once:

- center's x is inside `[x1, x2]` → `cx = xCenter`, that term contributes 0
- center is left of the rectangle → `cx = x1`
- center is right of the rectangle → `cx = x2`

Combined over both axes, an interior center gives `P = C` and distance 0; a
center beside an edge gives the perpendicular foot on that edge; a center
diagonally outside gives the nearest corner. No case analysis needed in the code.

Finally, compare squared distances instead of distances:

```
(xCenter - cx)² + (yCenter - cy)² <= radius²
```

Squaring is monotone on non-negative numbers, so the comparison is unchanged, and
we stay in exact integer arithmetic — no `math.Sqrt`, no epsilon tuning. With
coordinates bounded by 10⁴, each difference is at most 2·10⁴, so the sum is at
most 8·10⁸, comfortably inside `int`.

Walking through Example 2 (`radius = 1`, center `(1, 1)`, rectangle
`x1=1, y1=-3, x2=2, y2=-1`): `cx = clamp(1, 1, 2) = 1` and
`cy = clamp(1, -3, -1) = -1`, so the closest point is `(1, -1)` with squared
distance `0 + 4 = 4 > 1 = radius²` → `false`. The center is horizontally aligned
with the rectangle but two units above its top edge.

## Complexity Analysis

Time Complexity: O(1) — a fixed number of comparisons and multiplications, with
no loops and no dependence on coordinate magnitudes.
Space Complexity: O(1) — only a few integer temporaries.

## Edge Cases

- **Center inside the rectangle.** Both clamps return the center itself, giving
  squared distance 0, which is `<= radius²` for any `radius >= 1`. Correctly
  `true`. A corner-only or edge-only check would get this wrong.
- **Tangency (touching at exactly one point).** The problem counts a shared point
  as overlapping, so the comparison must be `<=`, not `<`. Example 1 is exactly
  this case: the circle touches the rectangle only at `(1, 0)`. Using `<` fails
  the very first example.
- **Center diagonally past a corner.** E.g. `radius = 1`, center `(0, 0)`,
  rectangle `(1, 1, 2, 2)`: the nearest point is the corner `(1, 1)` at squared
  distance 2 > 1 → `false`. The circle crosses both the vertical line `x = 1` and
  the horizontal line `y = 1`, so any approach that tests the two axes in
  isolation ("does the x-range overlap AND the y-range overlap?") wrongly reports
  `true`. Both axes must be combined in a single distance.
- **Center aligned with an edge but outside.** Example 2 above: the x-projection
  overlaps the rectangle while the y-distance decides the answer. The clamp gives
  the perpendicular foot on the edge, not a corner.
- **Rectangle entirely inside the circle.** The clamped point is the rectangle
  corner/edge point nearest the center, which is *within* the circle, so the
  check returns `true` — no separate containment test needed.
- **Extreme coordinates.** Center at `(-10⁴, -10⁴)` with the rectangle near
  `(10⁴, 10⁴)` yields squared distance up to 8·10⁸. Fine for Go's `int`, but
  worth noting if you port this to a 32-bit signed type where `radius² ` and the
  sum still fit, yet a careless `dist⁴`-style manipulation would not.
- **Degenerate rectangles are impossible.** Constraints guarantee `x1 < x2` and
  `y1 < y2`, so the rectangle always has positive area; the clamp would still be
  correct for a degenerate point or segment anyway.

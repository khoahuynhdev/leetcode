package main

// 1386. Cinema Seat Allocation
//
// Rows are independent, and n can be up to 1e9 while reservedSeats holds at
// most 1e4 entries, so we never iterate over rows. A row with no reservation in
// seats 2..9 always seats 2 groups (blocks 2-5 and 6-9, which are disjoint), so
// every row missing from the map contributes 2.
//
// For the affected rows we keep a 10-bit mask (bit s set == seat s reserved),
// discarding seats 1 and 10 because they belong to no block. A block is usable
// iff mask&blockMask == 0. Within a row the greedy is forced: if the two outer
// blocks are both free the row gives 2; otherwise any single free block gives 1;
// otherwise 0. The middle block can never be worth more than 1, so it only
// matters once an outer block is already broken.
//
// Time: O(len(reservedSeats)), Space: O(len(reservedSeats)).

const (
	blockLeft   = 1<<2 | 1<<3 | 1<<4 | 1<<5 // seats 2,3,4,5
	blockMiddle = 1<<4 | 1<<5 | 1<<6 | 1<<7 // seats 4,5,6,7
	blockRight  = 1<<6 | 1<<7 | 1<<8 | 1<<9 // seats 6,7,8,9
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	masks := make(map[int]int)
	for _, rs := range reservedSeats {
		row, seat := rs[0], rs[1]
		if seat < 2 || seat > 9 {
			// Seats 1 and 10 are part of no block; they cannot block anything.
			continue
		}
		masks[row] |= 1 << uint(seat)
	}

	// Rows never touched (or touched only at seats 1/10) always seat 2 groups.
	total := 2 * (n - len(masks))

	for _, mask := range masks {
		leftFree := mask&blockLeft == 0
		rightFree := mask&blockRight == 0
		switch {
		case leftFree && rightFree:
			total += 2
		case leftFree || rightFree || mask&blockMiddle == 0:
			total++
		}
	}

	return total
}

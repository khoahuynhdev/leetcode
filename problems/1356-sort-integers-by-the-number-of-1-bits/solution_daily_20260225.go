package main

import (
	"math/bits"
	"sort"
)

// sortByBits sorts integers by the number of 1-bits in their binary
// representation (ascending), breaking ties by value (ascending).
// Uses sort.Slice with a two-level custom comparator.
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		bi, bj := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if bi != bj {
			return bi < bj
		}
		return arr[i] < arr[j]
	})
	return arr
}

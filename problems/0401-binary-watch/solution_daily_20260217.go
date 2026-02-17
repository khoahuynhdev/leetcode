package main

import (
	"fmt"
	"math/bits"
)

// readBinaryWatch enumerates all valid times (h:0-11, m:0-59) and collects
// those whose total set bit count in h and m equals turnedOn.
func readBinaryWatch(turnedOn int) []string {
	var result []string
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if bits.OnesCount(uint(h))+bits.OnesCount(uint(m)) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return result
}

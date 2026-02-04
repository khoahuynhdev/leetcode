package solution

import (
	"strconv"
)

func digitCount(num string) bool {
	dict := map[string]int{}
	// if i repeat number of rept -> true
	//  else false
	for i := 0; i < len(num); i++ {
		digit := string(num[i])
		dict[digit]++
	}
	for i := 0; i < len(num); i++ {
		val, _ := strconv.Atoi(string(num[i]))
		if dict[strconv.Itoa(i)] != val {
			return false
		}
	}
	return true
}

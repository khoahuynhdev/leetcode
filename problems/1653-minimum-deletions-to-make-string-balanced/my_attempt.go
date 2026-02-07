package main

import "fmt"

func my_minimumDeletions(s string) int {
	n := len(s)
	maxKeep := 1
	str := []byte(s)
	for i := 1; i < n; i++ {
		maxValid := maxKeep
		if isValid(str[i-1], str[i]) {
			maxValid++
		} else {
			j := i - 1
			for j >= 0 && !isValid(str[j], str[i]) {
				fmt.Println("not valid at j ", j, " and i", i)
				maxValid--
				j--
			}
		}
		maxKeep = maxValid
	}
	// i - maxValid - maxKeep
	// 1 - 2        - 2
	// 2 - 3        - 3
	// 3 - 4        - 3
	return n - maxKeep
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isValid(i, j byte) bool {
	if i == 'a' && j == 'a' {
		return true
	}
	if i == 'b' && j == 'b' {
		return true
	}
	if i == 'a' && j == 'b' {
		return true
	}
	return false
}

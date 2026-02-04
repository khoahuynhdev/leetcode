package solution

// we have 2 approaches here
// using without sorting yield much better run time

// func findWinners(matches [][]int) [][]int {
//     dict := map[int]int{}
//     win, lose := []int{}, []int{}
//     for _,match := range matches {
//         w,l := match[0], match[1]
//         if _,ok := dict[w]; !ok {
//             dict[w] = 0
//         }
//         dict[l]--
//     }

//	    for k,v := range dict {
//	        if v == 0 { win = append(win, k)}
//	        if v == -1 { lose = append(lose, k)}
//	    }
//	    sort.Ints(win)
//	    sort.Ints(lose)
//	    return [][]int{win, lose}
//	}
func findWinners(matches [][]int) [][]int {
	arr := [1e5 + 1]int{}
	win, lose, m := []int{}, []int{}, 0
	for _, match := range matches {
		w, l := match[0], match[1]
		if m < w {
			m = w
		}
		if m < l {
			m = l
		}
		if arr[w] == 0 {
			arr[w]++
		}
		if arr[l] > 0 {
			arr[l] -= 2
		} else {
			arr[l]--
		}
	}
	for i := 0; i <= m; i++ {
		if arr[i] == 1 {
			win = append(win, i)
		}
		if arr[i] == -1 {
			lose = append(lose, i)
		}
	}
	return [][]int{win, lose}
}

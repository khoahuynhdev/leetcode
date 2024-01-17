package solution

func uniqueOccurrences(arr []int) bool {
	f, d := map[int]int{}, map[int]bool{}
	for _, v := range arr {
		f[v]++
	}
	for _, v := range f {
		if d[v] {
			return false
		} else {
			d[v] = true
		}
	}
	return true
}

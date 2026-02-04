package solution

// Intuition: Don't do Graph
func destCity(paths [][]string) string {
	ct := make(map[string]int)
	for _, path := range paths {
		ct[path[0]] += 1
		ct[path[1]] += 0

		if ct[path[0]] > 1 {
			delete(ct, path[0])
		}
		if ct[path[1]] > 1 {
			delete(ct, path[1])
		}
	}
	for k, v := range ct {
		if v == 0 {
			return string(k)
		}
	}
	return ""
}

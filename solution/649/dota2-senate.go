package solution

func predictPartyVictory(senate string) string {
	q := []byte(senate)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if len(q) == 0 {
			if string(cur) == "R" {
				return "Radiant"
			} else {
				return "Dire"
			}
		}
		i := 0
		for ; i < len(q); i++ {
			if i == len(q)-1 && cur == q[i] {
				if string(cur) == "R" {
					return "Radiant"
				} else {
					return "Dire"
				}
			}

			if cur != q[i] {
				break
			}
		}
		q = append(q[:i], q[i+1:]...)
		q = append(q, cur)
	}
	return "Dire"
}

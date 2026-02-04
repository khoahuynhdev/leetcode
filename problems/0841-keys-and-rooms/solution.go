package solution

// https://leetcode.com/problems/keys-and-rooms/description/?envType=study-plan-v2&envId=leetcode-75
// Intuition: simple straight forward BFS
// NOTE: just need to check the len(seen) -> number of visited rooms vs len(rooms) -> number of available rooms
func canVisitAllRooms(rooms [][]int) bool {
	seen := make(map[int]bool)
	q := make([]int, len(rooms))
	q[0] = 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			if seen[q[i]] {
				continue
			}
			seen[q[i]] = true
			for _, j := range rooms[q[i]] {
				if !seen[j] {
					q = append(q, j)
				}
			}
		}
		q = q[size:]
	}

	return len(seen) == len(rooms)
}

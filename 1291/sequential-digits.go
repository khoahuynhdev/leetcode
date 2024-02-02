package solution

func sequentialDigits(low int, high int) []int {
	result := make([]int, 0)

	for i := 1; i <= 9; i++ {
		num := i

		for j := i + 1; j <= 9; j++ {
			num = num*10 + j

			if num >= low && num <= high {
				result = append(result, num)
			}
		}
	}

	sort.Ints(result)
	return result
}

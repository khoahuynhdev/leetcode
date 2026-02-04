package solution

func FastExp(a, n int) uint64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return uint64(a)
	}
	if n < 0 {
		return 1 / FastExp(a, -n)
	}
	if n%2 == 0 {
		res := FastExp(a, n/2)
		return res * res
	}
	return FastExp(a, n-1) * uint64(a)
}

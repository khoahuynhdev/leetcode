package main

// 3348. Smallest Divisible Digit Product II
// https://leetcode.com/problems/smallest-divisible-digit-product-ii/
//
// A zero-free number only uses digits 1..9, so its digit product is always of
// the form 2^a * 3^b * 5^c * 7^d. If t has any prime factor outside {2,3,5,7}
// the answer is "-1"; otherwise t is tracked as the exponent vector that still
// has to be covered.
//
// The key tool is minDigits: the fewest digits needed to cover a leftover
// vector. Factors of 5 and 7 can only come from the digits 5 and 7, so they
// cost one digit each. Factors of 2 and 3 are packed with 8 (2^3) and 9 (3^2),
// with a single 6 replacing the tail when exactly one 2 and one 3 are left.
//
// With that oracle the answer is built greedily:
//  1. If num is zero-free and already covers t, return num.
//  2. Otherwise keep the longest zero-free prefix num[:i], raise position i to
//     the smallest digit d > num[i] whose remainder is still coverable by the
//     n-1-i free slots, and fill those slots with the smallest digits possible.
//     Scanning i from right to left gives the longest shared prefix, hence the
//     smallest result of the same length.
//  3. If nothing of length n works, use the shortest longer length that can
//     cover t and fill it the same greedy way.

// digitReq is the still-uncovered exponent vector of t over the primes 2,3,5,7.
type digitReq struct {
	e2, e3, e5, e7 int
}

// digitFactors[d] is the prime factorisation of the digit d (1 has none).
var digitFactors = [10]digitReq{
	2: {e2: 1},
	3: {e3: 1},
	4: {e2: 2},
	5: {e5: 1},
	6: {e2: 1, e3: 1},
	7: {e7: 1},
	8: {e2: 3},
	9: {e3: 2},
}

func clampZero(v int) int {
	if v < 0 {
		return 0
	}
	return v
}

func (r digitReq) isMet() bool {
	return r.e2 == 0 && r.e3 == 0 && r.e5 == 0 && r.e7 == 0
}

// divide returns what is left of r once the digit d has been contributed.
// Each prime is handled independently, so clamping at zero per prime is exact.
func (r digitReq) divide(d int) digitReq {
	f := digitFactors[d]
	return digitReq{
		e2: clampZero(r.e2 - f.e2),
		e3: clampZero(r.e3 - f.e3),
		e5: clampZero(r.e5 - f.e5),
		e7: clampZero(r.e7 - f.e7),
	}
}

// minDigits is the fewest digits in 1..9 whose product covers r.
//
// Only the digit 5 supplies a factor 5 and only 7 supplies a factor 7, so those
// exponents cost one digit each. For 2s and 3s the densest carriers are 8 and
// 9; using a 6 (which carries one of each) pays off exactly once, when the
// leftovers after that packing are one 2 and one 3.
func (r digitReq) minDigits() int {
	n := (r.e2+2)/3 + (r.e3+1)/2
	if r.e2%3 == 1 && r.e3%2 == 1 {
		n--
	}
	return n + r.e5 + r.e7
}

// smallestSuffix returns the lexicographically smallest k-digit zero-free
// string whose product covers r. It assumes r.minDigits() <= k.
func smallestSuffix(k int, r digitReq) []byte {
	out := make([]byte, k)
	for i := 0; i < k; i++ {
		free := k - i - 1
		for d := 1; d <= 9; d++ {
			if next := r.divide(d); next.minDigits() <= free {
				out[i] = byte('0' + d)
				r = next
				break
			}
		}
	}
	return out
}

func smallestNumber(num string, t int64) string {
	var need digitReq
	for t%2 == 0 {
		t /= 2
		need.e2++
	}
	for t%3 == 0 {
		t /= 3
		need.e3++
	}
	for t%5 == 0 {
		t /= 5
		need.e5++
	}
	for t%7 == 0 {
		t /= 7
		need.e7++
	}
	if t != 1 {
		// A digit product can never contain a prime above 7.
		return "-1"
	}

	n := len(num)

	// rem[i] is what t still needs after consuming the prefix num[:i].
	rem := make([]digitReq, n+1)
	rem[0] = need
	for i := 0; i < n; i++ {
		rem[i+1] = rem[i].divide(int(num[i] - '0'))
	}

	// A kept prefix must be zero-free.
	firstZero := n
	for i := 0; i < n; i++ {
		if num[i] == '0' {
			firstZero = i
			break
		}
	}

	if firstZero == n && rem[n].isMet() {
		return num
	}

	limit := firstZero
	if limit > n-1 {
		limit = n - 1
	}
	for i := limit; i >= 0; i-- {
		for d := int(num[i]-'0') + 1; d <= 9; d++ {
			next := rem[i].divide(d)
			if next.minDigits() > n-1-i {
				continue
			}
			out := make([]byte, 0, n)
			out = append(out, num[:i]...)
			out = append(out, byte('0'+d))
			out = append(out, smallestSuffix(n-1-i, next)...)
			return string(out)
		}
	}

	// Nothing of length n works: take the shortest longer length that fits.
	length := n + 1
	if m := need.minDigits(); m > length {
		length = m
	}
	return string(smallestSuffix(length, need))
}

package solution

const PrimeRK = 16777619

func HashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*PrimeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, PrimeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func IndexRabinKarp(s, substr string) int {
	// Rabin-Karp search
	hashss, pow := HashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*PrimeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= PrimeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}

func IndexByteString(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	n := len(needle)
	switch {
	case n == 0:
		return 0
	case n == len(haystack):
		if haystack == needle {
			return 0
		}
	}

	c0 := needle[0]
	c1 := needle[1]
	i := 0
	t := len(haystack) - n + 1
	fails := 0
	for i < t {
		// find the position of first char that matches c0
		if haystack[i] != c0 {
			o := IndexByteString(haystack[i+1:t], c0)
			if o < 0 {
				return -1
			}
			i += o + 1
		}
		if haystack[i+1] == c1 && haystack[i:i+n] == needle {
			return i
		}
		i++
		fails++
		if fails >= 4+i>>4 && i < t {
			// See comment in ../bytes/bytes.go.
			j := IndexRabinKarp(haystack[i:], needle)
			if j < 0 {
				return -1
			}
			return i + j
		}
	}
	return -1
}

func StrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}

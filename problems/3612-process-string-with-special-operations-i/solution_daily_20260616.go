package main

// 3612. Process String with Special Operations I
//
// Approach: straightforward left-to-right simulation. We maintain a []byte
// buffer `result` and apply each character of s:
//   - lowercase letter: append it
//   - '*': remove the last character (no-op if the buffer is empty)
//   - '#': duplicate the buffer (append a copy of it to itself)
//   - '%': reverse the buffer in place
// The constraints are tiny (len(s) <= 20), so performing each operation
// literally is both simple and fast.
func processStr(s string) string {
	result := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case '*':
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		case '#':
			result = append(result, result...)
		case '%':
			for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
				result[l], result[r] = result[r], result[l]
			}
		default:
			result = append(result, c)
		}
	}

	return string(result)
}

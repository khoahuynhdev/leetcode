package main

import "strings"

// Approach: load knowledge into a hash map, then scan s once. Characters outside
// brackets are copied directly; when a '(' is found, read up to the next ')'
// (brackets never nest), look up the key, and append its value or "?".
// Time: O(n + k), Space: O(n + k).
func evaluate(s string, knowledge [][]string) string {
	m := make(map[string]string, len(knowledge))
	for _, kv := range knowledge {
		m[kv[0]] = kv[1]
	}

	var sb strings.Builder
	sb.Grow(len(s))
	for i := 0; i < len(s); {
		if s[i] != '(' {
			sb.WriteByte(s[i])
			i++
			continue
		}
		j := i + 1
		for s[j] != ')' {
			j++
		}
		if v, ok := m[s[i+1:j]]; ok {
			sb.WriteString(v)
		} else {
			sb.WriteByte('?')
		}
		i = j + 1
	}
	return sb.String()
}

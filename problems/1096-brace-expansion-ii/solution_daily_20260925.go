package main

import "sort"

// Approach: recursive descent parsing.
// The grammar is parsed with three mutually recursive levels:
//   expr   := term (',' term)*        -> union of term sets
//   term   := factor factor*          -> cartesian concatenation of factor sets
//   factor := letter | '{' expr '}'   -> singleton set or nested expression
// Each level returns a set (map[string]struct{}), which deduplicates words
// naturally. The final set is collected and sorted.
func braceExpansionII(expression string) []string {
	pos := 0

	var parseExpr func() map[string]struct{}

	parseFactor := func() map[string]struct{} {
		if expression[pos] == '{' {
			pos++ // consume '{'
			set := parseExpr()
			pos++ // consume '}'
			return set
		}
		// Read a run of letters as a single word; concatenating single
		// letters one by one would produce the same result.
		start := pos
		for pos < len(expression) && expression[pos] >= 'a' && expression[pos] <= 'z' {
			pos++
		}
		return map[string]struct{}{expression[start:pos]: {}}
	}

	parseTerm := func() map[string]struct{} {
		result := map[string]struct{}{"": {}}
		for pos < len(expression) && (expression[pos] == '{' || (expression[pos] >= 'a' && expression[pos] <= 'z')) {
			next := parseFactor()
			product := make(map[string]struct{}, len(result)*len(next))
			for a := range result {
				for b := range next {
					product[a+b] = struct{}{}
				}
			}
			result = product
		}
		return result
	}

	parseExpr = func() map[string]struct{} {
		result := parseTerm()
		for pos < len(expression) && expression[pos] == ',' {
			pos++ // consume ','
			for w := range parseTerm() {
				result[w] = struct{}{}
			}
		}
		return result
	}

	set := parseExpr()
	words := make([]string, 0, len(set))
	for w := range set {
		words = append(words, w)
	}
	sort.Strings(words)
	return words
}

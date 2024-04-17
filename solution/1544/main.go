package main

func makeGood(s string) string {

	for i := 0; i < len(s)-1; i++ {

		// ? is same letter

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[i+1])) {

			continue

		}

		// ? NOT good string pattern

		if unicode.IsUpper(rune(s[i])) && unicode.IsLower(rune(s[i+1])) {

			return makeGood(s[:i] + s[i+2:])

		} else if unicode.IsLower(rune(s[i])) && unicode.IsUpper(rune(s[i+1])) {

			return makeGood(s[:i] + s[i+2:])

		}

	}

	return s

}
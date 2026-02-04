package solution

// Intuition: naive solution
// sort the string then compare one-by-one
// time complexity: O(nlogn)
// space complexity: O(1)
// func isAnagram(s string, t string) bool {
//  if len(s) != len(t) {
//      return false
//  }
//  strS := []rune(s)
//  strT := []rune(t)
//  sort.Slice(strS, func(i, j int) bool { //sort the string using the function
//       return strS[i] < strS[j]
//    })
// sort.Slice(strT, func(i, j int) bool { //sort the string using the function
//       return strT[i] < strT[j]
//    })

//	 return string(strS) == string(strT)
//	}

// we can leverage a hashmap store all the char in the first string
// because anagram must be match with the total chars in string -> count if total chars are match
// sort the string then compare one-by-one
// time complexity: O(n)
// space complexity: O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict := make(map[rune]int)
	for _, v := range s {
		dict[v]++
	}
	for _, v := range t {
		dict[v]--
		if dict[v] < 0 {
			return false
		}
	}
	return true
}

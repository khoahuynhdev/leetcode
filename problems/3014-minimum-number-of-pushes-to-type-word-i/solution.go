package main

// 3014. Minimum Number of Pushes to Type Word I
//
// All letters in word are distinct, so each is typed exactly once and the
// answer depends only on len(word). There are 8 usable keys (2..9), which
// gives 8 slots costing 1 push, 8 costing 2, 8 costing 3, and so on. Greedily
// filling the cheapest slots first means the letter at 0-based index i costs
// i/8 + 1 pushes; summing that over all indices gives the minimum total.
func minimumPushes(word string) int {
	pushes := 0
	for i := 0; i < len(word); i++ {
		pushes += i/8 + 1
	}
	return pushes
}

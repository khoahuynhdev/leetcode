package solution

// https://leetcode.com/problems/design-hashset/description/

type MyHashSet struct {
	bucketArray []bool
}

func Constructor() MyHashSet {
	// constraint 0 <= key <= 106
	// At most 104 calls will be made to add, remove, and contains.
	return MyHashSet{make([]bool, 1000001)}
}

func (hs *MyHashSet) Add(key int) {
	hs.bucketArray[key] = true
}

func (hs *MyHashSet) Remove(key int) {
	hs.bucketArray[key] = false
}

func (hs *MyHashSet) Contains(key int) bool {
	return hs.bucketArray[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

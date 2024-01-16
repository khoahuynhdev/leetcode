package solution

import "math/rand"

type RandomizedSet struct {
	storage map[int]int
	index   []int
}

// NOTE: Important part of this problem is the data structure
// supporting O(1) accessing
// supporting O(1) update
// supporting O(1) random access
// Array support O(1) random access
// Map support O(1) access and delete

func Constructor() RandomizedSet {
	return RandomizedSet{
		storage: map[int]int{},
		index:   []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.storage[val]
	if !ok {
		this.storage[val] = len(this.index)
		this.index = append(this.index, val)
	}
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.storage[val]
	if ok {
		idxVal := this.storage[val]
		lstIdx := len(this.index) - 1
		// fmt.Println(idxVal, lstIdx)
		this.storage[this.index[lstIdx]] = idxVal
		this.index[idxVal], this.index[lstIdx] = this.index[lstIdx], this.index[idxVal]
		delete(this.storage, val)
		this.index = this.index[:len(this.index)-1]
		// fmt.Println(this.index, this.storage)
	}
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	// fmt.Println(this.storage)
	// fmt.Println(this.index)
	return this.index[rand.Intn(len(this.index))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

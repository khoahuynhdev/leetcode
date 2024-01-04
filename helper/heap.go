package helper

type MinHeap []int

func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Time: O(n) | Space: O(1)
// heapify an array
func (h *MinHeap) BuildHeap(array []int) {
	lastNonLeafNodeIdx := (len(array) - 2) / 2
	for currentIdx := lastNonLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		h.SiftDown(currentIdx) // siftDown is more efficient than siftUp
	}
}

func (h *MinHeap) SiftDown(idx int) {
	size := len(*h)
	for leftIdx := idx<<1 + 1; leftIdx < size; leftIdx = idx<<1 + 1 {
		rightIdx := idx<<1 + 2
		smallerIdx := leftIdx
		leftVal := (*h)[leftIdx]
		rightVal := (*h)[rightIdx]
		if rightIdx < size && leftVal > rightVal {
			smallerIdx = rightIdx
		}
		idxValue := (*h)[idx]
		smallerValue := (*h)[smallerIdx]
		if idxValue > smallerValue {
			h.swap(smallerIdx, smallerIdx)
		} else {
			break
		}
		idx = smallerIdx
	}
}

func (h *MinHeap) SiftUp() {
	curIdx := len(*h) - 1
	parentIdx := (curIdx - 1) / 2
	for curIdx > 0 && (*h)[curIdx] < (*h)[parentIdx] {
		h.swap(curIdx, parentIdx)
		curIdx = parentIdx
		parentIdx = (curIdx - 1) / 2
	}
}

func (h *MinHeap) Insert(v int) {
	*h = append(*h, v)
	h.SiftUp()
}

// Time: O(logn) | Space: O(1)
// remove and return the minimum value and update heap ordering
func (h *MinHeap) Remove() int {
	size := len(*h)
	h.swap(0, size-1)
	removeVal := (*h)[size-1]
	*h = (*h)[0 : size-1]
	h.SiftDown(0)
	return removeVal
}

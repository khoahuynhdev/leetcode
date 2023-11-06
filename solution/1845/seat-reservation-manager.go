package solution

import (
	"container/heap"
)

// naive solution
// type SeatManager struct {
// 	seats []bool
// }
//
// func Constructor(n int) SeatManager {
// 	queue := make([]bool, n+1)
// 	return SeatManager{
// 		queue,
// 	}
// }
//
// func (this *SeatManager) Reserve() int {
// 	for i := 1; i < len(this.seats); i++ {
// 		if !this.seats[i] {
// 			this.seats[i] = true
// 			return i
// 		}
// 	}
// 	return -1
// }
//
// func (this *SeatManager) Unreserve(seatNumber int) {
// 	this.seats[seatNumber] = false
// }

type IntHeap []int

// utility func
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

type SeatManager struct {
	seats IntHeap
}

func Constructor(n int) SeatManager {
	h := IntHeap{}
	for i := 1; i <= n; i++ {
		h = append(h, i)
	}
	return SeatManager{h}
}

func (this *SeatManager) Reserve() int {
	res := heap.Pop(&this.seats)

	return res.(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.seats, seatNumber)
}

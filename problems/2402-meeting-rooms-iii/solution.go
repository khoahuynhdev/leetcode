package solution

type Heap []int
type BHeap [][2]int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool { return  h[i] < h[j]}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}

func (h *Heap) Push(v any) {
  *h = append(*h, v.(int))
}

func (h *Heap) Pop() any {
  old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h BHeap) Len() int { return len(h) }
func (h BHeap) Less(i, j int) bool { 
  if h[i][0] == h[j][0] {
    return h[i][1] < h[j][1]
  }
  return  h[i][0] < h[j][0]
}
func (h BHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *BHeap) Push(v any) {
  *h = append(*h, v.([2]int))
}

func (h *BHeap) Pop() any {
  old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mostBooked(n int, mts [][]int) int {
  // sort mts
  sort.Slice(mts, func (i, j int) bool {
    return mts[i][0] < mts[j][0]
  })
  bt, ans := make([]int, n), 0
  ff := &Heap{}
  nr := &BHeap{}
  for i,_ := range bt {
    heap.Push(ff, i)
  }
  for _,mt := range mts {
    for nr.Len() > 0 {
      p:=heap.Pop(nr).([2]int);
      if p[0] <= mt[0] {
        heap.Push(ff, p[1])
      } else {
        // fmt.Println("push")
        heap.Push(nr, p)
        break
      }
    }
    // if there are free rooms
    if ff.Len() > 0 {
      //  assign to free room
      rm := heap.Pop(ff).(int)
      bt[rm]++
      heap.Push(nr, [2]int{mt[1], rm})
    } else {
      slot := heap.Pop(nr).([2]int)
      bt[slot[1]]++
      slot[0] += (mt[1] - mt[0])
      heap.Push(nr, slot)
      // else assign to room end the soonest
    }
    // fmt.Println(*nr)
  }

  // fmt.Println(mts, bt)
  for i, v := range bt {
    if v > bt[ans] { ans = i}
  }
  return ans
}

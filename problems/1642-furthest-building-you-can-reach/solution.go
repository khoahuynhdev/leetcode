package solution

// Don't simulate jump, you won't get it
// use ladder for largest jumps
// use bricks for smaller ones
// if you overdo the bricks, you cannot perform the next jump -> this is the furthest you can go
type Heap []int

func (h Heap) Len() int { return len(h)}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}
func (h Heap) Less(i,j int) bool { return h[i] < h[j]}

func (h *Heap) Push(v interface{}) {
  *h = append(*h, v.(int))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)-1
    v := old[n]
    *h = old[:n]
    return v
}

func furthestBuilding(h []int, b int, l int) int {
  used := &Heap{}
  for i :=0;i<len(h);i++ {
    // end of building
    if i == len(h) - 1 { return i}
    // if can move -> just move
    if h[i] >= h[i+1] { 
      continue 
    } else {
      // if cannot move return i
      build := h[i+1] - h[i]
      if used.Len() <= l {
        heap.Push(used, build)
      }
      if used.Len() > l {
        b -= heap.Pop(used).(int)
      }
      if b < 0 {
        return i
      }
    }
  }
  return len(h) - 1
}

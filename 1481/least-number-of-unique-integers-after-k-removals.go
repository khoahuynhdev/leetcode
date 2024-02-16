package solution

// intuition: need a way to remove from least fre to max fre
// after each removal, remove the key from the hashmap
// space: O(n)
// time: O(n)

func findLeastNumOfUniqueInts(arr []int, n int) int {
  uniq, f := make(map[int]int), make([][]int, len(arr) + 1)
  for _, num := range arr {
    uniq[num]++
  }
  for k,v := range uniq {
    if f[v] != nil {
      f[v] = append(f[v], k)
    } else {
      f[v] = []int{k}
    }
  }
  for k,v := range f {
    for _,i := range v {
      if n >= k {
        delete(uniq, i)
        n -= k
      } else {
        break
      }
    }
  }
  return len(uniq)
}

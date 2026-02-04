package solution

// intuition: for counting chars, array is much better than hashmap
import (
  "sort"
  )
func frequencySort(s string) string {
    fre, bts := make([]int, 255), []byte(s)
    for i:=0;i<len(s);i++ {
        fre[s[i]]++
    }
    sort.Slice(bts, func (i, j int) bool {
        if fre[bts[i]] == fre[bts[j]] {
            return bts[i] > bts[j]
        } else {
            return fre[bts[i]] > fre[bts[j]]
        }
    })
    return string(bts)
}

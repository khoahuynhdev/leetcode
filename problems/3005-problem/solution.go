package main

func maxFrequencyElements(nums []int) int {
  f, nf, max := make([]int, 101), make([]int, 101), 0
  for _, num := range nums {
    nf[num]++
    if nf[num] - 1 > 0 {
      f[nf[num] - 1] -= nf[num]
    }
    f[nf[num]]+= nf[num]
     // fmt.Println(f[nf[num]], nf[num], num)
    if max < nf[num] { max = nf[num]}
  }
  return f[max]
}

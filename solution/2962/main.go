package main

/*
func countSubarrays(nums []int, k int) int64 {
  ans,max,f,cnt := 0,0,make([]int, len(nums)),0
  // m: max ele
  // [1,2,3,3,3]
  // [0,]
  // O((n - k)n/k)) = (n^2/k - n) = O(n^2)
  // q: can we get number of m at any range? -> yes
  for _,n := range nums {
    if max < n { max = n}
  }
  for i:=0;i<len(f);i++ {
    if nums[i] == max { cnt++ }
    f[i] = cnt
  }
  for i:=k; i <= len(nums);i++ {
    for j:=0;j<=len(nums) - i;j++ {
      var l int
      if j == 0 { l = 0 } else { l = f[j - 1]}
      r := f[j + i - 1]
      if r - l >= k { 
        ans++ 
      }
    }
  }
  return int64(ans)
}
*/


/*
## Approach

- Use sliding window approach to keep track of the current subarray's numbers and the occurences of maximum number.
- Each move to the right of the array: If it is the maximum number, add one to its occurences. While it has greater than or equal k occurences, then the window have to shift the left side and remove somes number repeatly until it pass another maximum number then only k - 1 occurences left.
Now the subarray [left - 1 .. right] has exactly k occurences of maximum number, so these subarrays [0..right], [1..right], .., [left - 1 .. right] have at least k occurences, add the number to result.

## Complexity

- Time complexity: O(n)O(n)O(n)
- Space complexity: O(1)O(1)O(1)

*/
func countSubarrays(nums []int, k int) int64 {
    m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}

	var res int64 = 0
	cnt := 0
	l := 0
	for _, num := range nums {
		if num == m {
			cnt++
		}

		for cnt >= k {
			if nums[l] == m {
				cnt--
			}
			l++
		}

		res += int64(l)
	}

	return res
}

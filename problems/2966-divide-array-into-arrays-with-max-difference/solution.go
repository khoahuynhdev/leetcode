package solution

func divideArray(nums []int, k int) [][]int {
    ans := make([][]int, 0)
    sort.Ints(nums)
    for i:=0;i<len(nums);i+=3 {
        if nums[i+2] - nums[i+1] <= k && nums[i+2] - nums[i] <= k {
            ans = append(ans,[]int{nums[i], nums[i+1], nums[i+2]})
        } else {
            return [][]int{}
        }
    }
    return ans
}

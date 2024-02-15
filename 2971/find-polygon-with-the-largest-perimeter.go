package solution

func largestPerimeter(nums []int) int64 {
    sort.Ints(nums)
    preSum := make([]int, len(nums))
    preSum[0] = nums[0]
    for i:=1;i<len(nums);i++{
        preSum[i] = preSum[i-1] + nums[i]
    }
    var ans int64 = -1
    for i:=1;i<len(nums);i++{
        if nums[i] < preSum[i-1] {
            if i == len(nums) - 1 {
                return int64(preSum[i])
            } else {
                if i >=2 && ans < int64(preSum[i]) { ans = int64(preSum[i])}
            }
        } 
    }
    return ans
}

package main
func subarraysWithKDistinct(nums []int, k int) int {
    return atMostKDistinct(nums, k) - atMostKDistinct(nums, k-1)
}

func atMostKDistinct(nums []int, k int) int {
    myMap := make(map[int]int)
    left, count, result := 0, 0, 0

    for right, num := range nums {
        if myMap[num] == 0 {
            count++
        }
        myMap[num]++

        for count > k {
            myMap[nums[left]]--
            if myMap[nums[left]] == 0 {
                count--
            }
            left++
        }

        result += right - left + 1
    }

    return result
}

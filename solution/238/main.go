package main

func productExceptSelf(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    product := 1
    zeros := 0

    for _, num := range nums {
        if num == 0 {
            zeros++
            continue
        }
        product *= num
    }

    if zeros == 1 {
        for i, num := range nums {
            if num == 0 {
                ans[i] = product
            }
        }
    } else if zeros == 0 {
        for i, num := range nums {
            ans[i] = product / num
        }
    }

    return ans
}

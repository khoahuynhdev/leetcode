package main

func trap(height []int) int {
    n := len(height)
    res := 0
    var st []int

    for r := 0; r < n; r++ {
        for len(st) > 0 && height[st[len(st) - 1]] <= height[r] {
            m := st[len(st) - 1] 
            st = st[:len(st) - 1]

            if len(st) > 0 {
                l := st[len(st) - 1]
                d := min(height[l], height[r])
                k := r - 1 - (l + 1) + 1
                res += k * (d - height[m])
            }
        }
        st = append(st, r)
    }

    return res
}

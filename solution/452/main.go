package main
func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        if points[i][1] == points[j][1] {
            return points[i][0] < points[j][0]
        }
        return points[i][1] < points[j][1]
    })

    res := 1
    pos := points[0][1]
    for _, point := range points {
        if point[0] <= pos {
            continue
        }

        res++
        pos = point[1]
    }

    return res
}

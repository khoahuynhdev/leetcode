package solution

func cherryPickup(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    mem := make([][][]int, m)
    for i := range mem {
        mem[i] = make([][]int, n)
        for j := range mem[i] {
            mem[i][j] = make([]int, n)
            for k := range mem[i][j] {
                mem[i][j][k] = -1
            }
        }
    }

    return cherryPick(grid, 0, 0, n-1, mem)
}

func cherryPick(grid [][]int, x, y1, y2 int, mem [][][]int) int {
    if x == len(grid) {
        return 0
    }
    if y1 < 0 || y1 == len(grid[0]) || y2 < 0 || y2 == len(grid[0]) {
        return 0
    }
    if mem[x][y1][y2] != -1 {
        return mem[x][y1][y2]
    }

    currRow := grid[x][y1]
    if y1 != y2 {
        currRow += grid[x][y2]
    }

    maxCherries := 0
    for d1 := -1; d1 <= 1; d1++ {
        for d2 := -1; d2 <= 1; d2++ {
            cherries := currRow + cherryPick(grid, x+1, y1+d1, y2+d2, mem)
            if cherries > maxCherries {
                maxCherries = cherries
            }
        }
    }

    mem[x][y1][y2] = maxCherries
    return maxCherries
}

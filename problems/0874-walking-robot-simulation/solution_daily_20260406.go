package main

// Walking Robot Simulation
// Simulate the robot step-by-step using direction vectors and a hash set of obstacles.
// Track the maximum squared Euclidean distance after every step.

type point struct {
	x, y int
}

func robotSim(commands []int, obstacles [][]int) int {
	// Direction vectors: North, East, South, West
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}

	// Build obstacle set
	obstacleSet := make(map[point]bool, len(obstacles))
	for _, o := range obstacles {
		obstacleSet[point{o[0], o[1]}] = true
	}

	x, y := 0, 0
	dir := 0 // 0=North, 1=East, 2=South, 3=West
	maxDist := 0

	for _, cmd := range commands {
		switch {
		case cmd == -2: // turn left
			dir = (dir + 3) % 4
		case cmd == -1: // turn right
			dir = (dir + 1) % 4
		default: // move forward cmd steps
			for i := 0; i < cmd; i++ {
				nx, ny := x+dx[dir], y+dy[dir]
				if obstacleSet[point{nx, ny}] {
					break
				}
				x, y = nx, ny
				dist := x*x + y*y
				if dist > maxDist {
					maxDist = dist
				}
			}
		}
	}

	return maxDist
}

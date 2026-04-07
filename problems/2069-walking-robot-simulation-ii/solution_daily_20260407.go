package main

// Walking Robot Simulation II
// Model the robot's position as an index along the grid perimeter.
// Use modular arithmetic in Step() for O(1) per call.
// Convert perimeter index to (x, y) and direction on demand.

type Robot struct {
	w, h  int
	pos   int
	moved bool
}

func Constructor(width int, height int) Robot {
	return Robot{w: width, h: height}
}

func (r *Robot) Step(num int) {
	p := 2*(r.w-1) + 2*(r.h-1)
	r.pos = (r.pos + num) % p
	r.moved = true
}

func (r *Robot) GetPos() []int {
	x, y := r.getXY()
	return []int{x, y}
}

func (r *Robot) GetDir() string {
	pos := r.pos
	w, h := r.w, r.h
	if pos == 0 {
		if r.moved {
			return "South"
		}
		return "East"
	}
	if pos < w {
		return "East"
	}
	if pos < w+h-1 {
		return "North"
	}
	if pos < 2*w+h-2 {
		return "West"
	}
	return "South"
}

func (r *Robot) getXY() (int, int) {
	pos := r.pos
	w, h := r.w, r.h
	if pos < w {
		return pos, 0
	}
	pos -= w
	if pos < h-1 {
		return w - 1, pos + 1
	}
	pos -= h - 1
	if pos < w-1 {
		return w - 2 - pos, h - 1
	}
	pos -= w - 1
	return 0, h - 2 - pos
}

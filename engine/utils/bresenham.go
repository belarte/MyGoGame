package utils

// Line computes a straight line between two Coords.
// It uses the Bresenham line drawing algorithm.
// The return line does NOT contain the 'from' point.
func Line(from, to Coord) (result []Coord) {
	dx, dy, sx, sy := getParameters(from, to)
	err := dx - dy

	for {
		if from.X == to.X && from.Y == to.Y {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			from.X += sx
		}
		if e2 < dx {
			err += dx
			from.Y += sy
		}

		result = append(result, Coord{from.X, from.Y})
	}

	return
}

func getParameters(from, to Coord) (dx, dy, sx, sy int) {
	dx = to.X - from.X
	if dx < 0 {
		dx = -dx
	}
	dy = to.Y - from.Y
	if dy < 0 {
		dy = -dy
	}

	if from.X < to.X {
		sx = 1
	} else {
		sx = -1
	}
	if from.Y < to.Y {
		sy = 1
	} else {
		sy = -1
	}

	return
}

// Circle compute the circle centered on given Coord with given radius.
// It uses the Midpoint circle drawing algorithm.
func Circle(centre Coord, radius int) (result []Coord) {
	x := radius
	y := 0
	decisionOver2 := 1 - x

	// buffer is required to qvoid duplicates
	buffer := make(map[Coord]bool)

	for x >= y {
		addCoordToCircle(x, y, centre, buffer)
		y++
		if decisionOver2 <= 0 {
			decisionOver2 += 2*y + 1
		} else {
			x--
			decisionOver2 += 2*(y-x) + 1
		}
	}

	for key := range buffer {
		result = append(result, key)
	}
	return
}

func addCoordToCircle(x, y int, centre Coord, buffer map[Coord]bool) {
	buffer[Coord{x + centre.X, y + centre.Y}] = true
	buffer[Coord{y + centre.X, x + centre.Y}] = true
	buffer[Coord{-x + centre.X, y + centre.Y}] = true
	buffer[Coord{-y + centre.X, x + centre.Y}] = true
	buffer[Coord{-x + centre.X, -y + centre.Y}] = true
	buffer[Coord{-y + centre.X, -x + centre.Y}] = true
	buffer[Coord{x + centre.X, -y + centre.Y}] = true
	buffer[Coord{y + centre.X, -x + centre.Y}] = true
}

package utils

func Bresenham(from, to Coord) (result []Coord) {
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

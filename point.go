package dollar

import "math"

type Point struct {
	X, Y float64
}

func PathLength(points []Point) int {
	return 0
}

func Distance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

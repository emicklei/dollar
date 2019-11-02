package dollar

import "math"

type DollarRecognizer struct {
	Unistrokes []Unistroke
}

func Resample(points []Point, n int) []Point {
	I := float64(PathLength(points) / (n - 1)) // interval length
	D := 0.0
	newPoints := []Point{}
	for i := 1; i < len(points); i++ {
		d := Distance(points[i-1], points[i])
		if (D + d) >= I {
			qx := points[i-1].X + ((I-D)/d)*(points[i].X-points[i-1].X)
			qy := points[i-1].Y + ((I-D)/d)*(points[i].Y-points[i-1].Y)
			q := Point{X: qx, Y: qy}
			newPoints = append(newPoints, q)
			points[i] = q
			D = 0.0
		} else {
			D += d
		}
	}
	if len(newPoints) == n-1 {
		newPoints[len(newPoints)] = Point{X: points[len(points)-1].X, Y: points[len(points)-1].Y}
	}
	return newPoints
}

func IndicativeAngle(points []Point) float64 {
	c := Centroid(points)
	return math.Atan2(c.Y-points[0].Y, c.X-points[0].X)
}

func RotateBy(points []Point, radians float64) []Point {
	c := Centroid(points)
	cos := math.Cos(radians)
	sin := math.Sin(radians)
	newpoints := []Point{}
	for _, each := range points {
		qx := (each.X-c.X)*cos - (each.Y-c.Y)*sin + c.X
		qy := (each.X-c.X)*sin - (each.Y-c.Y)*cos + c.Y
		newpoints = append(newpoints, Point{X: qx, Y: qy})
	}
	return newpoints
}

func ScaleTo(points []Point, size float64) []Point {
	B := BoundingBox(points)
	newpoints := []Point{}
	for _, each := range points {
		qx := each.X * (size / B.Width)
		qy := each.Y * (size / B.Height)
		newpoints = append(newpoints, Point{X: qx, Y: qy})
	}
	return newpoints
}

func TranslateTo(points []Point, pt Point) []Point {
	c := Centroid(points)
	newpoints := []Point{}
	for _, each := range points {
		qx := each.X + pt.X - c.X
		qy := each.Y + pt.Y - c.Y
		newpoints = append(newpoints, Point{X: qx, Y: qy})
	}
	return newpoints
}

// Bug?
func Vectorize(points []Point) {
	//	sum := 0.0
	//	vector := []float64{}
	//	for _, each := range points {
	//		vector[len(length)] =
	//	}
}

func OptimalCosineDistance(v1 []float64, v2 []float64) float64 {
	a := 0.0
	b := 0.0
	for i := 0; i < len(v1); i += 2 {
		a += v1[i]*v2[i] + v1[i+1]*v2[i+1]
		b += v1[i]*v2[i+1] - v1[i+1]*v2[i]
	}
	angle := math.Atan(b / a)
	return math.Acos(a*math.Cos(angle) + b*math.Sin(angle))
}

func DistanceAtBestAngle(points []Point, T Unistroke, a, b, threshold float64) {}

func Centroid(points []Point) Point {
	x := 0.0
	y := 0.0
	for _, each := range points {
		x += each.X
		y += each.Y
	}
	x /= float64(len(points))
	y /= float64(len(points))
	return Point{X: x, Y: y}
}

func BoundingBox(points []Point) Rectangle {
	minX, maxX, minY, maxY := math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64
	for _, each := range points {
		minX = math.Min(minX, each.X)
		minY = math.Min(minY, each.Y)
		maxX = math.Max(maxX, each.X)
		maxY = math.Max(maxY, each.Y)
	}
	return Rectangle{X: minX, Y: minY, Width: maxX - minX, Height: maxY - minY}
}

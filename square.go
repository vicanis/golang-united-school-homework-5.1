package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	return Point{
		x: s.start.x + int(s.a),
		y: s.start.y + int(s.a),
	}
}

func (s Square) Area() uint {
	return uint(math.Pow(float64(s.a), 2))
}

func (s Square) Perimeter() uint {
	return uint(s.a * 4)
}

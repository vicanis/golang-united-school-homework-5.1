package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (receiver) End() Point {
	// implement me
}

func (receiver) Area() uint {
	// implement me
}

func (receiver) Perimeter() uint {
	// implement me
}

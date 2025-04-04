package shapes

// Shape is an interface for 2D shapes that can calculate their area.
type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}

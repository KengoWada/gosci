package shapes

import "fmt"

var ErrNonPositiveSide = fmt.Errorf("length and width must be greater than zero")

// Rectangle represents a rectangle.
type Rectangle struct {
	Length float64
	Width  float64
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() (float64, error) {
	if r.Length <= 0 || r.Width <= 0 {
		return 0, ErrNonPositiveSide
	}
	return r.Length * r.Width, nil
}

// Perimeter calculates the perimeter of a rectangle.
func (r Rectangle) Perimeter() (float64, error) {
	if r.Length <= 0 || r.Width <= 0 {
		return 0, ErrNonPositiveSide
	}
	return 2 * (r.Length + r.Width), nil
}

package shapes

import (
	"fmt"
	"math"
)

var ErrNonPositiveRadius = fmt.Errorf("radius must be greater than zero")

// Circle represents a circle.
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle.
func (c Circle) Area() (float64, error) {
	if c.Radius <= 0 {
		return 0, ErrNonPositiveRadius
	}
	return math.Pi * c.Radius * c.Radius, nil
}

// Perimeter calculates the circumference of a circle.
func (c Circle) Perimeter() (float64, error) {
	if c.Radius <= 0 {
		return 0, ErrNonPositiveRadius
	}
	return 2 * math.Pi * c.Radius, nil
}

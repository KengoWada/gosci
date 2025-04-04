package shapes

import "fmt"

var (
	ErrNonPositiveDimension = fmt.Errorf("dimensions must be greater than zero")
	ErrInvalidTriangleSides = fmt.Errorf("the given sides cannot form a triangle")
)

// Triangle represents a triangle.
type Triangle struct {
	SideA  float64
	SideB  float64
	SideC  float64
	Base   float64
	Height float64
}

// Area calculates the area of a triangle.
func (t Triangle) Area() (float64, error) {
	if t.Base <= 0 || t.Height <= 0 {
		return 0, ErrNonPositiveDimension
	}
	return 0.5 * t.Base * t.Height, nil
}

// Perimeter calculates the perimeter of a triangle.
func (t Triangle) Perimeter() (float64, error) {
	if t.SideA <= 0 || t.SideB <= 0 || t.SideC <= 0 {
		return 0, ErrNonPositiveDimension
	}

	// Triangle Inequality Theorem check
	if t.isTriangle() {
		return 0, ErrInvalidTriangleSides
	}

	return t.SideA + t.SideB + t.SideC, nil
}

// isTriangle checks that the provided sides make a valid triangle.
func (t Triangle) isTriangle() bool {
	return t.SideA+t.SideB <= t.SideC || t.SideA+t.SideC <= t.SideB || t.SideB+t.SideC <= t.SideA
}

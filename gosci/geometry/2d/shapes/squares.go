package shapes

// Square represents a square.
type Square struct {
	Side float64
}

// Area calculates the area of a square.
func (s Square) Area() (float64, error) {
	if s.Side <= 0 {
		return 0, ErrNonPositiveSide
	}
	return s.Side * s.Side, nil
}

// Perimeter calculates the perimeter of a square.
func (s Square) Perimeter() (float64, error) {
	if s.Side <= 0 {
		return 0, ErrNonPositiveSide
	}
	return 4 * s.Side, nil
}

package shapes_test

import (
	"testing"

	geometry2d "github.com/KengoWada/gosci-lib/gosci/geometry/2d/shapes"
	"github.com/stretchr/testify/assert"
)

type rectangleTestData struct {
	name      string
	rectangle geometry2d.Rectangle
	want      float64
	wantErr   bool
}

func TestRectangleArea(t *testing.T) {
	tests := []rectangleTestData{
		{
			name:      "Valid length and width",
			rectangle: geometry2d.Rectangle{Length: 2, Width: 4},
			want:      8,
			wantErr:   false,
		},
		{
			name:      "Small length and width",
			rectangle: geometry2d.Rectangle{Length: 0.1, Width: 0.3},
			want:      0.03,
			wantErr:   false,
		},
		{
			name:      "Large length and width",
			rectangle: geometry2d.Rectangle{Length: 1000, Width: 10000},
			want:      10000000,
			wantErr:   false,
		},
		{
			name:      "Negative length",
			rectangle: geometry2d.Rectangle{Length: -5, Width: 4},
			wantErr:   true,
		},
		{
			name:      "Zero width",
			rectangle: geometry2d.Rectangle{Length: 5, Width: 0},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rectangle.Area()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, tt.want, got, 1e-9)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []rectangleTestData{
		{
			name:      "Valid length and width",
			rectangle: geometry2d.Rectangle{Length: 2, Width: 4},
			want:      12,
			wantErr:   false,
		},
		{
			name:      "Small length and width",
			rectangle: geometry2d.Rectangle{Length: 0.1, Width: 0.3},
			want:      0.8,
			wantErr:   false,
		},
		{
			name:      "Large length and width",
			rectangle: geometry2d.Rectangle{Length: 1000, Width: 10000},
			want:      22000,
			wantErr:   false,
		},
		{
			name:      "Negative length",
			rectangle: geometry2d.Rectangle{Length: -5, Width: 4},
			wantErr:   true,
		},
		{
			name:      "Zero width",
			rectangle: geometry2d.Rectangle{Length: 5, Width: 0},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rectangle.Perimeter()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, tt.want, got, 1e-9)
			}
		})
	}
}

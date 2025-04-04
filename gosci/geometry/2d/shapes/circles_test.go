package shapes_test

import (
	"math"
	"testing"

	geometry2d "github.com/KengoWada/gosci-lib/gosci/geometry/2d/shapes"
	"github.com/stretchr/testify/assert"
)

type circleTestData struct {
	name    string
	circle  geometry2d.Circle
	want    float64
	wantErr bool
}

func TestCircleArea(t *testing.T) {
	tests := []circleTestData{
		{
			name:    "Valid Radius",
			circle:  geometry2d.Circle{Radius: 5},
			want:    math.Pi * 25,
			wantErr: false,
		},
		{
			name:    "Small Radius",
			circle:  geometry2d.Circle{Radius: 0.1},
			want:    math.Pi * 0.01,
			wantErr: false,
		},
		{
			name:    "Large Radius",
			circle:  geometry2d.Circle{Radius: 1000},
			want:    math.Pi * 1000000,
			wantErr: false,
		},
		{
			name:    "Negative Radius",
			circle:  geometry2d.Circle{Radius: -5},
			wantErr: true,
		},
		{
			name:    "Zero Radius",
			circle:  geometry2d.Circle{Radius: 0},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.circle.Area()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, tt.want, got, 1e-9)
			}
		})
	}
}

func TestCirclePerimeter(t *testing.T) {
	tests := []circleTestData{
		{
			name:    "Valid Radius",
			circle:  geometry2d.Circle{Radius: 5},
			want:    math.Pi * 10,
			wantErr: false,
		},
		{
			name:    "Small Radius",
			circle:  geometry2d.Circle{Radius: 0.1},
			want:    math.Pi * 0.2,
			wantErr: false,
		},
		{
			name:    "Large Radius",
			circle:  geometry2d.Circle{Radius: 1000},
			want:    math.Pi * 2000,
			wantErr: false,
		},
		{
			name:    "Negative Radius",
			circle:  geometry2d.Circle{Radius: -5},
			wantErr: true,
		},
		{
			name:    "Zero Radius",
			circle:  geometry2d.Circle{Radius: 0},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.circle.Perimeter()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, tt.want, got, 1e-9)
			}
		})
	}
}

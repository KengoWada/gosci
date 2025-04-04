package shapes_test

import (
	"testing"

	geometry2d "github.com/KengoWada/gosci-lib/gosci/geometry/2d/shapes"
	"github.com/stretchr/testify/assert"
)

type triangleTestData struct {
	name     string
	triangle geometry2d.Triangle
	want     float64
	wantErr  bool
}

func TestTriangleArea(t *testing.T) {
	tests := []triangleTestData{
		{
			name:     "Valid triangle",
			triangle: geometry2d.Triangle{Base: 5, Height: 8},
			want:     0.5 * 40,
			wantErr:  false,
		},
		{
			name:     "Small triangle",
			triangle: geometry2d.Triangle{Base: 0.1, Height: 0.5},
			want:     0.5 * 0.05,
			wantErr:  false,
		},
		{
			name:     "Large triangle",
			triangle: geometry2d.Triangle{Base: 100, Height: 1000},
			want:     0.5 * 100000,
			wantErr:  false,
		},
		{
			name:     "Negative side triangle",
			triangle: geometry2d.Triangle{Base: -100, Height: 1000},
			wantErr:  true,
		},
		{
			name:     "Zero side triangle",
			triangle: geometry2d.Triangle{Base: 100, Height: 0},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.triangle.Area()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, tt.want, got, 1e-9)
			}
		})
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tests := []triangleTestData{
		{
			name:     "Valid triangle",
			triangle: geometry2d.Triangle{SideA: 3, SideB: 5, SideC: 7},
			want:     15,
			wantErr:  false,
		},
		{
			name:     "Invalid triangle",
			triangle: geometry2d.Triangle{SideA: 3, SideB: 5, SideC: 9},
			wantErr:  true,
		},
		{
			name:     "Small triangle",
			triangle: geometry2d.Triangle{SideA: 0.3, SideB: 0.5, SideC: 0.7},
			want:     1.5,
			wantErr:  false,
		},
		{
			name:     "Large triangle",
			triangle: geometry2d.Triangle{SideA: 20000, SideB: 15000, SideC: 10000},
			want:     45000,
			wantErr:  false,
		},
		{
			name:     "Negative side triangle",
			triangle: geometry2d.Triangle{SideA: 3, SideB: -5, SideC: 7},
			wantErr:  true,
		},
		{
			name:     "Zero side triangle",
			triangle: geometry2d.Triangle{SideA: 3, SideB: 5, SideC: 0},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.triangle.Perimeter()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, tt.want, got, 1e-9)
			}
		})
	}
}

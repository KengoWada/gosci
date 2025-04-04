package shapes_test

import (
	"testing"

	geometry2d "github.com/KengoWada/gosci-lib/gosci/geometry/2d/shapes"
	"github.com/stretchr/testify/assert"
)

type squareTestData struct {
	name    string
	square  geometry2d.Square
	want    float64
	wantErr bool
}

func TestSquareArea(t *testing.T) {
	tests := []squareTestData{
		{
			name:    "Valid side",
			square:  geometry2d.Square{Side: 5},
			want:    25,
			wantErr: false,
		},
		{
			name:    "Large side",
			square:  geometry2d.Square{Side: 0.1},
			want:    0.01,
			wantErr: false,
		},
		{
			name:    "Large side",
			square:  geometry2d.Square{Side: 10000},
			want:    100000000,
			wantErr: false,
		},
		{
			name:    "Negative side",
			square:  geometry2d.Square{Side: -5},
			wantErr: true,
		},
		{
			name:    "Zero side",
			square:  geometry2d.Square{Side: 0},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.square.Area()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, tt.want, got, 1e-9)
			}
		})
	}
}

func TestSquarePerimeter(t *testing.T) {
	tests := []squareTestData{
		{
			name:    "Valid side",
			square:  geometry2d.Square{Side: 5},
			want:    20,
			wantErr: false,
		},
		{
			name:    "Small side",
			square:  geometry2d.Square{Side: 0.1},
			want:    0.4,
			wantErr: false,
		},
		{
			name:    "Large side",
			square:  geometry2d.Square{Side: 10000},
			want:    40000,
			wantErr: false,
		},
		{
			name:    "Negative side",
			square:  geometry2d.Square{Side: -5},
			wantErr: true,
		},
		{
			name:    "Zero side",
			square:  geometry2d.Square{Side: 0},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.square.Perimeter()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, tt.want, got, 1e-9)
			}
		})
	}
}

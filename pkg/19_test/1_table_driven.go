package test

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
		wantErr  bool
	}{
		{"positive numbers", 1, 2, 3, false},
		{"negative numbers", -1, -1, -2, false},
		{"zero", 0, 0, 0, false},
		{"overflow case", math.MaxInt64, 1, 0, true}, // 假設有 error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.a, tt.b)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.expected, got)
		})
	}
}

func Sum(a, b int) (int, error) {
	if math.MaxInt64-a < b {
		return 0, errors.New("overflow")
	}
	return a + b, nil
}

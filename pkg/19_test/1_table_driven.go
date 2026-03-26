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
	}{
		{"positive numbers", 1, 2, 3},
		{"negative numbers", -1, -1, -2},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.a, tt.b)
			require.NoError(t, err)
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

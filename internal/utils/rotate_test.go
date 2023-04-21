package utils

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		name      string
		input     []string
		positions int
		expected  []string
	}{
		{
			name:      "Positive positions",
			input:     []string{"A", "B", "C", "D", "E"},
			positions: 2,
			expected:  []string{"D", "E", "A", "B", "C"},
		},
		{
			name:      "No rotation",
			input:     []string{"A", "B", "C", "D", "E"},
			positions: 0,
			expected:  []string{"A", "B", "C", "D", "E"},
		},
		{
			name:      "Full rotation",
			input:     []string{"A", "B", "C", "D", "E"},
			positions: 5,
			expected:  []string{"A", "B", "C", "D", "E"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Rotate(tc.input, tc.positions)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

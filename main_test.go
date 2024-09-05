package main

import (
	"math"
	"testing"
)

// Helper function to check for float64 equality with a tolerance
func float64Equal(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

// Test cases for linear regression and correlation functionsfunc TestLinearRegression(t *testing.T) {
func TestLinearRegression(t *testing.T) {
	tests := []struct {
		name              string
		data              string
		expectedSlope     float64
		expectedIntercept float64
	}{
		{
			name:              "Simple positive slope",
			data:              "1\n2\n3\n4\n5",
			expectedSlope:     1.0,
			expectedIntercept: 1.0, // Corrected intercept
		},
		{
			name:              "Simple negative slope",
			data:              "5\n4\n3\n2\n1",
			expectedSlope:     -1.0,
			expectedIntercept: 5.0, // Corrected intercept
		},
		{
			name:              "Horizontal line",
			data:              "1\n1\n1\n1\n1",
			expectedSlope:     0.0,
			expectedIntercept: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slope, intercept := InterceptAndSlope(tt.data)
			if !float64Equal(slope, tt.expectedSlope, 1e-6) {
				t.Errorf("InterceptAndSlope() got slope = %v, want %v", slope, tt.expectedSlope)
			}
			if !float64Equal(intercept, tt.expectedIntercept, 1e-6) {
				t.Errorf("InterceptAndSlope() got intercept = %v, want %v", intercept, tt.expectedIntercept)
			}
		})
	}
}

func TestCorrelation(t *testing.T) {
	tests := []struct {
		name      string
		data      string
		expectedR float64
	}{
		{
			name:      "Perfect positive correlation",
			data:      "1\n2\n3\n4\n5",
			expectedR: 1.0,
		},
		{
			name:      "Perfect negative correlation",
			data:      "5\n4\n3\n2\n1",
			expectedR: -1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := correlation(tt.data)
			if !float64Equal(r, tt.expectedR, 1e-10) {
				t.Errorf("correlation() got = %v, want %v", r, tt.expectedR)
			}
		})
	}
}

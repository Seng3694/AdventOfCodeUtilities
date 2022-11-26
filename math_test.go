package aocutil

import (
	"fmt"
	"testing"
)

func TestClamp(t *testing.T) {
	tests := []struct {
		value, lower, upper, expected int32
	}{
		{10, 0, 20, 10},
		{10, 15, 20, 15},
		{10, 0, 5, 5},

		{-5, -10, 10, -5},
		{-15, -10, 10, -10},
		{25, -10, 10, 10},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v,%v,%v", test.value, test.lower, test.upper)
		t.Run(name, func(t *testing.T) {
			actual := Clamp(test.value, test.lower, test.upper)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}
}

func TestClampLower(t *testing.T) {
	tests := []struct {
		value, lower, expected int32
	}{
		{10, 0, 10},
		{10, 15, 15},

		{-5, -10, -5},
		{-15, -10, -10},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v,%v", test.value, test.lower)
		t.Run(name, func(t *testing.T) {
			actual := ClampLower(test.value, test.lower)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}
}

func TestClampUpper(t *testing.T) {
	tests := []struct {
		value, upper, expected int32
	}{
		{10, 20, 10},
		{30, 15, 15},

		{-5, -10, -10},
		{-15, -10, -15},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v,%v", test.value, test.upper)
		t.Run(name, func(t *testing.T) {
			actual := ClampUpper(test.value, test.upper)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}
}

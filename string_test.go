package aocutil

import (
	"fmt"
	"testing"
)

func TestIsHexChar(t *testing.T) {
	validHexChars := "abcdefABCDEF0123456789"

	for _, c := range validHexChars {
		name := fmt.Sprintf("%v", c)
		t.Run(name, func(t *testing.T) {
			actual := IsHexChar(c)
			if !actual {
				t.Errorf("got %v, expected %v", actual, true)
			}
		})
	}

	someInvalidHexChars := "ghijklmnopqrstuvwxyz /?!-+#"
	for _, c := range someInvalidHexChars {
		name := fmt.Sprintf("%v", c)
		t.Run(name, func(t *testing.T) {
			actual := IsHexChar(c)
			if actual {
				t.Errorf("got %v, expected %v", actual, false)
			}
		})
	}
}

func TestIsHexString(t *testing.T) {
	tests := []struct {
		text     string
		expected bool
	}{
		{"dead", true},
		{"1234567890", true},
		{"D34D", true},
		{"5fB0", true},
		{"nope", false},
		{"Hello World", false},
		{"0H4edll6oWoArld", false},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.text)
		t.Run(name, func(t *testing.T) {
			actual := IsHexString(test.text)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}

}

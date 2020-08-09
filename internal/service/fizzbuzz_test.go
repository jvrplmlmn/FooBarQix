package service

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},

		// Check boundaries
		{-7, "-7"},
		{-5, "-5"},
		{-3, "-3"},
		{-1, "-1"},
		{0, "0"},
		{100, "Buzz"},
		{101, "101"},
		{333, "333"},
	}

	s := NewFizzBuzz()
	for _, tt := range tests {
		name := func(in int) string {
			return fmt.Sprintf("%d", tt.in)
		}(tt.in)
		t.Run(name, func(t *testing.T) {
			if gotResult := s.CalculateForNumber(tt.in); gotResult != tt.out {
				t.Errorf("CalculateForNumber() = %v, want %v", gotResult, tt.out)
			}
		})
	}
}

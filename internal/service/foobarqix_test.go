package service

import (
	"fmt"
	"testing"
)

func TestFooBarQixService(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		// First 15 digits
		{1, "1"},
		{2, "2"},
		{3, "FooFoo"},
		{4, "4"},
		{5, "BarBar"},
		{6, "Foo"},
		{7, "QixQix"},
		{8, "8"},
		{9, "Foo"},
		{10, "Bar"},
		{11, "11"},
		{12, "Foo"},
		{13, "Foo"},
		{14, "Qix"},
		{15, "FooBarBar"},

		// Ambiguous cases
		{21, "FooQix"},
		{33, "FooFooFoo"},
		{51, "FooBar"},
		{53, "BarFoo"},

		{33, "FooFooFoo"},
		{35, "BarQixFooBar"},
		{37, "FooQix"},
		{53, "BarFoo"},
		{55, "BarBarBar"},
		{57, "FooBarQix"},
		{73, "QixFoo"},
		{75, "FooBarQixBar"},
		{77, "QixQixQix"},

		// Check boundaries
		{-7, "-7"},
		{-5, "-5"},
		{-3, "-3"},
		{-1, "-1"},
		{0, "0"},
		{100, "Bar"},
		{101, "101"},
		{333, "333"},
	}

	s := NewFooBarQix()
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

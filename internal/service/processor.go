package service

import (
	"strconv"

	"github.com/jvrplmlmn/FooBarQuix/internal/rules"
)

type Processor interface {
	CalculateForNumber(number int) string
}

type Service struct {
	lowerBound  int
	higherBound int
	rules       rules.Rules
}

func NewService(lowerBound, higherBound int, rules ...rules.Rule) *Service {
	return &Service{
		lowerBound:  lowerBound,
		higherBound: higherBound,
		rules:       rules,
	}
}

func (s *Service) CalculateForNumber(number int) (result string) {
	// If the number is out of bounds there is no need to transform it
	if number < s.lowerBound || number > s.higherBound {
		return strconv.Itoa(number)
	}

	// Process the rules sequentially
	for _, rule := range s.rules {
		if applies, text := rule.Apply(number); applies {
			result += text
		}
	}
	// If no rule applied, then return the number as-is
	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}

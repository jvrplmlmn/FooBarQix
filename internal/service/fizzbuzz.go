package service

import "github.com/jvrplmlmn/FooBarQuix/internal/rules"

func NewFizzBuzz() *Service {
	return NewService(1, 100,
		rules.NewDivisible(3, "Fizz"),
		rules.NewDivisible(5, "Buzz"),
	)
}

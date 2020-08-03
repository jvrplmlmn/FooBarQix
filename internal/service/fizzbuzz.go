package service

import "github.com/jvrplmlmn/FooBarQuix/internal/rules"

func NewFizzBuzzService() *Service {
	return NewService(1, 100,
		rules.NewDivisible(3, "Fizz"),
		rules.NewDivisible(5, "Buzz"),
	)
}

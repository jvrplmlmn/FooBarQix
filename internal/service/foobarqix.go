package service

import "github.com/jvrplmlmn/FooBarQuix/internal/rules"

func NewFooBarQix() *Service {
	return NewService(1, 100,
		rules.NewDivisible(3, "Foo"),
		rules.NewDivisible(5, "Bar"),
		rules.NewDivisible(7, "Qix"),
		rules.NewContains(rules.ContainsMap{
			'3': "Foo",
			'5': "Bar",
			'7': "Qix",
		}),
	)
}

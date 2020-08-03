package rules

type Rules []Rule

type Rule interface {
	Apply(number int) (applies bool, value string)
}

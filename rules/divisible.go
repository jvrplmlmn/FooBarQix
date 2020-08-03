package rules

type Divisible struct {
	factor int
	text   string
}

func NewDivisible(factor int, text string) *Divisible {
	return &Divisible{factor: factor, text: text}
}

func (d *Divisible) Apply(number int) (bool, string) {
	if d.factor == 0 {
		return false, ""
	}
	if number%d.factor == 0 {
		return true, d.text
	}
	return false, ""
}

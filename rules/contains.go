package rules

import "strconv"

type ContainsMap map[rune]string

type Contains struct {
	relationship ContainsMap
}

func NewContains(runeRelationship ContainsMap) *Contains {
	return &Contains{relationship: runeRelationship}
}

func (c *Contains) Apply(number int) (bool, string) {
	applies, value := false, ""
	for _, digit := range strconv.Itoa(number) {
		if val, ok := c.relationship[digit]; ok {
			value += val
			applies = true
		}
	}
	return applies, value
}

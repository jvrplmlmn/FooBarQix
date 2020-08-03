package rules

import (
	"fmt"
	"testing"
)

func TestContains_Apply(t *testing.T) {
	const (
		empty = ""
		foo   = "Foo"
		bar   = "Bar"
		qix   = "Qix"
	)

	containsMap := ContainsMap{
		'3': foo,
		'5': bar,
		'7': qix,
	}

	tests := []struct {
		number      int
		wantApplies bool
		wantValue   string
	}{
		// Applies when it is a exact match
		{3, true, foo},
		{5, true, bar},
		{7, true, qix},

		// Applies when the digit matches and is repeated
		{33, true, foo + foo},
		{55, true, bar + bar},
		{77, true, qix + qix},

		// Applies when there is a digit match but there is other non matching numbers
		{13, true, foo},
		{31, true, foo},
		{313, true, foo + foo},
		{15, true, bar},
		{51, true, bar},
		{515, true, bar + bar},
		{17, true, qix},
		{71, true, qix},
		{717, true, qix + qix},

		// Applies when there is multiple digit matches
		{35, true, foo + bar},
		{57, true, bar + qix},
		{73, true, qix + foo},
		{357, true, foo + bar + qix},

		// Does not apply because is does not contain a matching digit
		{0, false, empty},
		{2, false, empty},
		{4, false, empty},
	}
	for _, tt := range tests {
		name := func(applies bool, number int) string {
			if applies {
				return fmt.Sprintf("Applies to %d", number)
			} else {
				return fmt.Sprintf("Does not apply to %d", number)
			}
		}(tt.wantApplies, tt.number)
		t.Run(name, func(t *testing.T) {
			c := NewContains(containsMap)
			gotApplies, gotValue := c.Apply(tt.number)
			if gotApplies != tt.wantApplies {
				t.Errorf("Apply() gotApplies = %v, wantApplies %v", gotApplies, tt.wantApplies)
			}
			if gotValue != tt.wantValue {
				t.Errorf("Apply() gotValue = %v, wantApplies %v", gotValue, tt.wantValue)
			}
		})
	}
}

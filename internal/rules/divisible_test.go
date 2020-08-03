package rules

import (
	"fmt"
	"testing"
)

func TestDivisible_Apply(t *testing.T) {
	const (
		empty = ""
		foo   = "Foo"
		bar   = "Bar"
		qix   = "Qix"
	)

	type fields struct {
		factor int
		text   string
	}

	fieldsDiv0 := fields{0, "this should not panic"}
	fieldsDiv3 := fields{3, foo}
	fieldsDiv5 := fields{5, bar}
	fieldsDiv7 := fields{7, qix}
	tests := []struct {
		number      int
		fields      fields
		wantApplies bool
		wantValue   string
	}{

		{0, fieldsDiv0, false, empty},
		{1, fieldsDiv0, false, empty},

		{0, fieldsDiv3, true, foo},
		{0, fieldsDiv5, true, bar},
		{0, fieldsDiv7, true, qix},

		{3, fieldsDiv3, true, foo},
		{5, fieldsDiv5, true, bar},
		{7, fieldsDiv7, true, qix},

		{6, fieldsDiv3, true, foo},
		{10, fieldsDiv5, true, bar},
		{14, fieldsDiv7, true, qix},

		{15, fieldsDiv3, true, foo},
		{15, fieldsDiv5, true, bar},

		{105, fieldsDiv3, true, foo},
		{105, fieldsDiv5, true, bar},
		{105, fieldsDiv7, true, qix},

		{74, fieldsDiv3, false, empty},
		{74, fieldsDiv5, false, empty},
		{74, fieldsDiv7, false, empty},
	}
	for _, tt := range tests {
		name := func(applies bool, number int, factor int) string {
			if applies {
				return fmt.Sprintf("%d is divisible by %d", number, factor)
			}
			return fmt.Sprintf("%d is not divisible by %d", number, factor)
		}(tt.wantApplies, tt.number, tt.fields.factor)

		t.Run(name, func(t *testing.T) {
			d := NewDivisible(tt.fields.factor, tt.fields.text)
			gotApplies, gotValue := d.Apply(tt.number)
			if gotApplies != tt.wantApplies {
				t.Errorf("Apply() gotApplies = %v, want %v", gotApplies, tt.wantApplies)
			}
			if gotValue != tt.wantValue {
				t.Errorf("Apply() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := newDollar(5)
	five.times(2)
	if five.amount != 10 {
		t.Errorf("expect %d actual %d", 10, five.amount)
	}
}

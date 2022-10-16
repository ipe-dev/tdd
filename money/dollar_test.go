package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := newDollar(5)
	product := five.times(2)
	if product.amount != 10 {
		t.Errorf("expect %d actual %d", 10, product.amount)
	}
	product = five.times(3)
	if product.amount != 15 {
		t.Errorf("expect %d actual %d", 15, five.amount)
	}
}

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

func TestEquality(t *testing.T) {
	if !newDollar(5).equals(newDollar(5)) {
		t.Error("expect equal but actual not equal")
	}
	if newDollar(5).equals(newDollar(6)) {
		t.Error("expect not equal but actual equal")
	}

}

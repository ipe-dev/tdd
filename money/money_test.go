package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if five.times(2) != *NewDollar(10) {
		t.Errorf("expect %d actual %d", 10, five.amount)
	}
	if five.times(3) != *NewDollar(15) {
		t.Errorf("expect %d actual %d", 15, five.amount)
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).equals(NewDollar(5).Money) {
		t.Error("expect 5dollar equal 5dollar but actual not equal")
	}
	if NewDollar(5).equals(NewDollar(6).Money) {
		t.Error("expect 6dollar not equal 6dollar but actual equal")
	}
	if !NewFranc(5).equals(NewFranc(5).Money) {
		t.Error("expect 5franc equal 5franc but actual not equal")
	}
	if NewFranc(5).equals(NewFranc(6).Money) {
		t.Error("expect 6franc not equal 5franc but actual equal")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if five.times(2) != *NewFranc(10) {
		t.Errorf("expect %d actual %d", 10, five.amount)
	}
	if five.times(3) != *NewFranc(15) {
		t.Errorf("expect %d actual %d", 15, five.amount)
	}
}

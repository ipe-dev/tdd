package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if five.Times(2) != NewDollar(10) {
		t.Errorf("expect %d actual %d", 10, five.amount)
	}
	if five.Times(3) != NewDollar(15) {
		t.Errorf("expect %d actual %d", 15, five.amount)
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).Equals(NewDollar(5)) {
		t.Error("expect 5dollar equal 5dollar but actual not equal")
	}
	if NewDollar(5).Equals(NewDollar(6)) {
		t.Error("expect 6dollar not equal 6dollar but actual equal")
	}
	if !NewFranc(5).Equals(NewFranc(5)) {
		t.Error("expect 5franc equal 5franc but actual not equal")
	}
	if NewFranc(5).Equals(NewFranc(6)) {
		t.Error("expect 6franc not equal 5franc but actual equal")
	}
	if NewDollar(5).Equals(NewFranc(5)) {
		t.Error("expect 5dollar not equal 5franc but actual equal")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if five.Times(2) != NewFranc(10) {
		t.Errorf("expect %d actual %d", 10, five.amount)
	}
	if five.Times(3) != NewFranc(15) {
		t.Errorf("expect %d actual %d", 15, five.amount)
	}
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.plus(five)
	bank := new(Bank) 
	reduced := bank.reduce(sum,"USD")
	if NewDollar(10) != reduced {
		t.Errorf("expect true actual false")
	}
}
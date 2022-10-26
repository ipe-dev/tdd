package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	t.Run("$5=$5", func(t *testing.T) {
		assert.True(t, NewDollar(5).Equals(NewDollar(5)))

	})
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
	reduced := bank.reduce(sum, "USD")
	assert.Equal(t, NewDollar(10), reduced)
}

func TestPulsReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.plus(five)
	sum := result.(Sum)
	t.Run("augendとfiveが等しいこと", func(t *testing.T) {
		assert.Equal(t,five, sum.augend)
	})

	t.Run("addendとfiveが等しいこと", func(t *testing.T) {
		assert.Equal(t, five, sum.addend)
	})
}

// func TestReduceSum(t *testing.T) {
// 	sum := NewSum(NewDollar(3), NewDollar(4))
// 	bank := NewBank()
// 	result := bank.reduce(sum, "USD")
// 	assert.Equal(t, NewDollar(7), result)

// }

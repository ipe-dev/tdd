package money

type Sum struct {
	augend Money
	addend Money
}

func NewSum(augend Money, addend Money) Sum {
	return Sum {
		augend: augend,
		addend: addend,
	}
}
func (s Sum) Plus(addend Money) Expression {
	return NewDollar(2)
}
func (s Sum) Reduce(to string) Money {
	amount := s.augend.amount + s.addend.amount
	return NewMoney(amount, to)
}
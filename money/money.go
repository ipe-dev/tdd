package money

type MoneyInterface interface {
	Times(int) Money
}

type Money struct {
	amount int
	name string
}

func (m Money) Equals(money Money) bool {
	return m.amount == money.amount && m.name == money.name
}

func NewDollar(amount int) Money {
	return Money{
		amount: amount,
		name: "dollar",
	}
}

func (m Money) Times(multiplier int) Money {
	// これがコピーを返すっていう意味か！
	return Money{
		amount: m.amount * multiplier,
		name: m.name,
	}
}
func NewFranc(amount int) Money {
	return Money {
		amount: amount,
		name: "franc",
	}
}
package money

type MoneyInterface interface {
	Equals() bool
}

type Money struct {
	amount int
	name string
}

func (m *Money) Equals(money Money) bool {
	return m.amount == money.amount && m.name == money.name
}

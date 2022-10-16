package money


type Money struct {
	amount int
}

func (m *Money) equals(money Money) bool {
	return m.amount == money.amount
}

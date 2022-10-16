package money

type Dollar struct {
	amount int
}
func newDollar(amount int) *Dollar {
	d := new(Dollar)
	d.amount = amount
	return d
}

func (d *Dollar) times(multiplier int) Dollar {
	// これがコピーを返すっていう意味か！
	return *newDollar(d.amount * multiplier)
}
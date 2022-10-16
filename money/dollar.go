package money

type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	d := new(Dollar)
	d.amount = amount
	return d
}

func (d *Dollar) times(multiplier int) Dollar {
	// これがコピーを返すっていう意味か！
	return *NewDollar(d.amount * multiplier)
}

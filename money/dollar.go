package money

type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	d := new(Dollar)
	d.amount = amount
	d.name = "dollar"
	return d
}

func (d *Dollar) Times(multiplier int) Dollar {
	// これがコピーを返すっていう意味か！
	return *NewDollar(d.amount * multiplier)
}

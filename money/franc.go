package money

type Franc struct {
	amount int
}

func newFranc(amount int) *Franc {
	f := new(Franc)
	f.amount = amount
	return f
}

func (f *Franc) times(multiplier int) Franc {
	return *newFranc(f.amount * multiplier)
}

func (f *Franc) equals(franc *Franc) bool {
	return f.amount == franc.amount
}

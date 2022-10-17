package money

type Franc struct {
	Money 
}

func NewFranc(amount int) *Franc {
	f := new(Franc)
	f.amount = amount
	f.name = "franc"
	return f
}

func (f *Franc) times(multiplier int) Franc {
	return *NewFranc(f.amount * multiplier)
}

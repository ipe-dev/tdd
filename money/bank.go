package money

type Bank struct {
	rates map[Pair]int
}

func (b *Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(*b, to)
}
func NewBank() Bank {
	return Bank{}
}

func (b *Bank)AddRate(from string, to string, rate int) {
	b.rates = make(map[Pair]int)
	b.rates[NewPair(from, to)] = rate
}
func (b Bank) Rate(from string, to string) int {
	if (from == to) {
		return 1
	}
	return b.rates[NewPair(from, to)]
}

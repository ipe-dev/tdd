package money
type Bank struct {
	
}
func (b *Bank) reduce(source Expression, to string) Money {
	sum := source.(Sum)
	return sum.reduce(to)
}
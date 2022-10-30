package money

type Expression interface {
	Plus(addend Money) Expression
	Reduce(to string) Money
}

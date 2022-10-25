package money

type Expression interface {
	plus(addend Money) Expression
}

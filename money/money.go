package money

type Accessor interface {
	Amount() int
	Currency() string
}

type Money struct {
	amount int
	currency string
}

func NewMoney(amount int, currency string) Money {
	return Money{
		amount: amount,
		currency: currency,
	}
}
func (m Money) Amount() int {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Equals(money Money) bool {
	return m.amount == money.amount && m.Currency() == money.Currency()
}

func NewDollar(amount int) Money {
	return Money{
		amount: amount,
		currency: "USD",
	}
}

func (m Money) Times(multiplier int) Expression {
	// これがコピーを返すっていう意味か！
	return Money{
		amount: m.amount * multiplier,
		currency: m.currency,
	}
}
func NewFranc(amount int) Money {
	return Money {
		amount: amount,
		currency: "CHF",
	}
}

func (m Money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}
func (m Money) Reduce (bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.amount / rate, to)
}
package shop

type Itemer interface {
	GetPrice() int
	GetName() string
	GetCount() int
	CountMinus()
	GetStatusItem() bool
}

type Useer interface {
	GetName() string
	GetCash() int
	CashMinus(int)
	GetStatusUser() bool
}

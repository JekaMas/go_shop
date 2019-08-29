package users

type user struct {
	name string
	cash int
}

func NewUser(name string, cash int) *user {
	return &user{
		name: name,
		cash: cash,
	}
}

func (u user) GetName() string {
	return u.name
}

func (u user) GetCash() int {
	return u.cash
}

func (u *user) CashMinus(sum int) {
	u.cash = u.cash - sum
}

func (u user) GetStatusUser() bool {
	return false
}

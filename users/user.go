package users

type user struct {
	Name string
	Cash int
}

func New(name string, cash int) *user {
	return &user{
		Name: name,
		Cash: cash,
	}
}

func (u user) GetName() string {
	return u.Name
}

func (u user) GetCash() int {
	return u.Cash
}

func (u *user) CashMinus(sum int) {
	u.Cash = u.Cash - sum
}

func (u user) GetStatusUser() bool {
	return false
}

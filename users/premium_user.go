package users

type premUser struct {
	user
}

func NewPrem(name string, cash int) *premUser {
	return &premUser{user{
		Name: name,
		Cash: cash,
	}}
}

func (pu premUser) GetStatusUser() bool {
	return true
}

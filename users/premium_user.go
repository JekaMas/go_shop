package users

type premUser struct {
	user
}

func NewPremUser(name string, cash int) *premUser {
	return &premUser{user{name, cash}}
}

func (pu premUser) GetStatusUser() bool {
	return true
}

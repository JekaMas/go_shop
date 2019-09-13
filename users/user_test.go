package users

import "testing"

func TestUser_GetName(t *testing.T) {

	t.Run("test GetName mortal user", func(t *testing.T) {
		user := user{"semen", 100}

		got := user.GetName()
		want := "semen"

		if got != want {
			t.Errorf("неправильно возвращает имя обычного смертного. Надо %q, возвращает %q", want, got)
		}
	})

	t.Run("test GetName prem user", func(t *testing.T) {
		user := premUser{user{"semen", 100}}

		got := user.GetName()
		want := "semen"

		if got != want {
			t.Errorf("неправильно возвращает имя прем смертного. Надо %q, возвращает %q", want, got)
		}
	})
}

func TestUser_GetCash(t *testing.T) {

	t.Run("test GetCash mortal user", func(t *testing.T) {
		user := user{"semen", 100}

		got := user.GetCash()
		want := 100

		if got != want {
			t.Errorf("неправильно возвращает кэш обычного смертного. Надо %d, возвращает %d", want, got)
		}
	})

	t.Run("test GetCash prem user", func(t *testing.T) {
		user := premUser{user{"semen", 100}}

		got := user.GetCash()
		want := 100

		if got != want {
			t.Errorf("неправильно возвращает кэш прем смертного. Надо %d, возвращает %d", want, got)
		}
	})
}
func TestUser_CashMinus(t *testing.T) {

	t.Run("CashMinus mortal user", func(t *testing.T) {
		user := user{"semen", 100}
		sum := 10

		user.CashMinus(sum)
		want := 90

		if user.GetCash() != want {
			t.Errorf("неправильно снимает деньги у простого смертного. Надо снять %d и получить %d, возвращает %d", sum, want, user.GetCash())
		}
	})

	t.Run("CashMinus prem user", func(t *testing.T) {
		user := premUser{user{"semen", 100}}
		sum := 10

		user.CashMinus(sum)
		want := 90

		if user.GetCash() != want {
			t.Errorf("неправильно снимает деньги у прем смертного. Надо снять %d и получить %d, возвращает %d", sum, want, user.GetCash())
		}
	})
}
func TestUser_GetStatusUser(t *testing.T) {

	t.Run("GetStatus mortal user", func(t *testing.T) {
		user := user{"semen", 100}

		got := user.GetStatusUser()
		want := false

		if got != want {
			t.Errorf("неправильно возвращает статус обычного смертного. Надо %t, возвращает %t", want, got)
		}
	})

	t.Run("GetStatus prem user", func(t *testing.T) {
		user := premUser{user{"semen", 100}}

		got := user.GetStatusUser()
		want := true

		if got != want {
			t.Errorf("неправильно возвращает статус обычного смертного. Надо %t, возвращает %t", want, got)
		}
	})
}

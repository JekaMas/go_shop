package shop

import (
	"github.com/JekaMas/shop/items"
	"github.com/JekaMas/shop/users"
	"reflect"
	"testing"
)

func TestShopuserUpdate(t *testing.T) {

	t.Run("delete user", func(t *testing.T) {

	})

	t.Run("delete prem user", func(t *testing.T) {

	})

	t.Run("update user", func(t *testing.T) {

	})

	t.Run("update prem user", func(t *testing.T) {

	})

	t.Run("create user", func(t *testing.T) {

	})

	t.Run("create prem user", func(t *testing.T) {

	})
}

func TestNewShop(t *testing.T) {
	t.Run("empty shop, nil users and nil items", func(t *testing.T) {
		shop := newShop(nil, nil)

		if len(shop.items) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.items)
		}

		if len(shop.users) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.items)
		}

		if len(shop.littleCache) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.items)
		}

		if len(shop.bigCache) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.items)
		}
	})

	t.Run("empty shop, empty users and empty items", func(t *testing.T) {
		shop := newShop([]Itemer{}, []Useer{})

		if len(shop.items) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.items)
		}

		if len(shop.users) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.items)
		}

		if len(shop.littleCache) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.items)
		}

		if len(shop.bigCache) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.items)
		}
	})

	t.Run("check items", func(t *testing.T) {
		shop := newTestShop()

		expected := map[string]Itemer{
			"apple": &items.Item{"apple", 20, 10},
			"banana": &items.PremItem{items.Item{"banana", 20, 20}},
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("usersMap формируется неправильно. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("check users", func(t *testing.T) {
		shop := newTestShop()

		expected := map[string]Useer{
			"dasha": users.New("dasha", 10),
			"masha": users.NewPrem("masha", 500),
		}

		if !reflect.DeepEqual(expected, shop.users) {
			t.Errorf("usersMap формируется неправильно. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("check shop", func(t *testing.T) {
		expectedShop := &shop{
			users: map[string]Useer{
				"dasha": users.New("dasha", 10),
				"masha": users.NewPrem("masha", 500),
			},
			items: map[string]Itemer{
				"apple": &items.Item{"apple", 20, 10},
				"banana": &items.PremItem{items.Item{"banana", 20, 20}},
			},
		}

		shop := newTestShop()

		if !reflect.DeepEqual(expectedShop.items, shop.items) {
			t.Errorf("usersMap формируется неправильно. expected %v, got %v", expectedShop.items, shop.items)
		}

		if !reflect.DeepEqual(expectedShop.users, shop.users) {
			t.Errorf("usersMap формируется неправильно. expected %v, got %v", expectedShop.users, shop.items)
		}

		if len(shop.littleCache) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.littleCache)
		}

		if len(shop.bigCache) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.bigCache)
		}
	})
}

func newTestShop() *shop {
	shopItems := []Itemer{
		&items.Item{"apple", 20, 10},
		&items.PremItem{items.Item{"banana", 20, 20}},
	}

	shopUsers := []Useer{
		users.New("dasha", 10),
		users.NewPrem("masha", 500),
	}

	return newShop(shopItems, shopUsers)
}
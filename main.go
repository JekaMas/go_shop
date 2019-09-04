package main

import (
	"github.com/JekaMas/shop/items"
	"github.com/JekaMas/shop/shop"
	"github.com/JekaMas/shop/users"
)

func main() {
	shopItems := []shop.Itemer{
		items.NewItem("apple", 20, 10),
		items.NewItem("tea", 20, 50),
		items.NewItem("salt", 0, 80),

		items.NewPremItem("banana", 20, 20),
		items.NewPremItem("potato", 20, 20),
	}

	shopUsers := []shop.Useer{
		users.NewUser("dasha", 10),
		users.NewUser("natasha", 500),
		users.NewUser("kirill", -1000),

		users.NewPremUser("masha", 1000),
	}

	//создает новый шоп и тут же обновляет его,
	// убирая все товары с невозможной стоимостью
	bigCacheMap := shop.NewBigCache(100)
	//bigCacheMap := &shop.BigCacheStub{}
	shop := shop.NewShop(shopItems, shopUsers, bigCacheMap)

	err := shop.Buy([]string{"apple", "banana"}, "masha")
	shop.Intrigue(err, "masha")

	err = shop.Buy([]string{"apple", "banana"}, "natasha")
	shop.Intrigue(err, "natasha")
}

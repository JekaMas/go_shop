package main

import (
	"fmt"

	"github.com/JekaMas/shop/items"
	"github.com/JekaMas/shop/shop"
	"github.com/JekaMas/shop/users"
)

func main() {
	fmt.Println("Hello World")

	items := []shop.Itemer{
		&items.Item{"apple", 20, 10},
		&items.Item{"tea", 20, 50},
		&items.Item{"salt", 0, 80},

		&items.PremItem{items.Item{"banana", 20, 20}},
		&items.PremItem{items.Item{"potato", 20, 20}},
	}

	users := []shop.Useer{
		users.New("dasha", 10),
		users.New("natasha", 500),
		users.New("kirill", -1000),

		users.NewPrem("masha", 500),
	}

	//создает новый шоп и тут же обновляет его,
	// убирая все товары с невозможной стоимостью
	shop := shop.NewShopPrepare(items, users)

	err := shop.Buy([]string{"apple", "banana"}, "masha")
	shop.Intrigue(err, "masha")

	err = shop.Buy([]string{"apple", "banana"}, "natasha")
	shop.Intrigue(err, "natasha")

}

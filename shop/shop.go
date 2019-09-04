package shop

import (
	"errors"
	"fmt"
	"sort"
)

type Shop struct {
	users       map[string]Useer
	items       map[string]Itemer
	littleCache map[string]map[string]int
	bigCache    cacher
}

func NewShop(items []Itemer, users []Useer, bigCache cacher) *Shop {
	shop := newShopPrepare(items, users)
	shop.bigCache = bigCache
	shop.newShopUpdate(items, users)

	return shop
}

func newShopPrepare(items []Itemer, users []Useer) *Shop {
	usersMap := make(map[string]Useer, len(items))
	itemsMap := make(map[string]Itemer, len(users))

	for _, user := range users {
		usersMap[user.GetName()] = user
	}
	for _, item := range items {
		itemsMap[item.GetName()] = item
	}

	littleCacheMap := make(map[string]map[string]int)

	return &Shop{
		users:       usersMap,
		items:       itemsMap,
		littleCache: littleCacheMap,
	}
}

func (s *Shop) newShopUpdate(items []Itemer, users []Useer) {
	//апдейт юзеров
	s.userUpdate(users)
	//апдейт предметов
	s.itemsUpdate(items)
	//сносим мапу, которая с длинными заказами
	s.bigCache.Clear()
	//редактируем мапу, которая с короткими заказами
	s.editLittleMap(items)
}

const deleteKey = -1

func (s *Shop) userUpdate(users []Useer) {
	for _, user := range users {
		_, ok := s.users[user.GetName()]
		if ok {
			delete(s.users, user.GetName())
			if user.GetCash() != deleteKey {
				s.users[user.GetName()] = user
			}
		} else {
			s.users[user.GetName()] = user
		}
	}
}

func (s *Shop) itemsUpdate(items []Itemer) {
	for _, item := range items {
		_, ok := s.items[item.GetName()]
		if ok {
			delete(s.items, item.GetName())

		}
		if item.GetPrice() > deleteKey {
			s.items[item.GetName()] = item
		}
	}
}

func (s *Shop) editLittleMap(items []Itemer) {
	for _, item := range items {
		_, ok := s.littleCache[item.GetName()]
		if ok {
			delete(s.littleCache, item.GetName())
		}
	}

	for _, item1 := range s.littleCache {
		for _, item2 := range items {
			_, ok := item1[item2.GetName()]
			if ok {
				delete(item1, item2.GetName())
			}
		}
	}
}

func (s *Shop) Buy(order []string, user string) error {
	//сначала чекаем существует ли товар, и есть ли он на складе
	err := s.checkItems(order)
	if err != nil {
		return err
	}

	//если товаров больше, чем два, делаем так
	if len(order) > 2 {
		err = s.buyBigPrice(order, user)

		return err
	}

	//если товаров меньше двух, то идем в мапу-мапу
	err = s.buyLittlePrice(order, user)
	return err
}

func (s *Shop) checkItems(order []string) error {
	for _, item := range order {
		_, ok := s.items[item]
		if !ok {
			return errors.New("товара не существует")
		}
		if s.items[item].GetCount() <= 0 {
			return errors.New("товар закончился")
		}
	}

	return nil
}

func (s *Shop) buyBigPrice(order []string, user string) error {
	//делаем хэш
	hash := makeHash(order)
	//если хэш найден, проводим оплату
	sum, ok := s.bigCache.Get(hash)
	if ok {
		err := s.payment(sum, user, order)
		if err == nil { //если деньги есть, делаем доставку
			s.delivery(order)
		}

		return err
	}
	//если хэша нет, собираем заказ вручную
	sum = s.collectOrder(order, user) //собираем заказ
	err := s.payment(sum, user, order)
	//делаем кэш. Даже если покупка не состоится (денег нет), все равно делаем
	s.bigCache.Set(hash, sum)
	//если ошибок нет, списываем товары
	if err == nil {
		s.delivery(order)
	}
	return err
}

func (s *Shop) collectOrder(order []string, user string) int {
	var sum int

	for _, item := range order {
		sum += s.items[item].GetPrice()
	}

	return sum
}

func (s *Shop) payment(sum int, user string, order []string) error {
	//смотрим статус пользователя - вип он или кто?
	//рассчитываем для него конечную стоимость исходя из статуса
	statusUser := s.users[user].GetStatusUser()
	//если тру, значит вип пользователь
	if statusUser == true {
		for _, item := range order {
			statusItem := s.items[item].GetStatusItem()
			if statusItem == false {
				sum = sum - s.items[item].GetPrice()
				newPriceItem := s.items[item].GetPrice() * 90 / 100
				sum = sum + newPriceItem
			}
		}
	}

	//если фолс, то пользователь обычный
	if statusUser == false {
		for _, item := range order {
			statusItem := s.items[item].GetStatusItem()
			if statusItem == true {
				sum = sum - s.items[item].GetPrice()
				newPriceItem := s.items[item].GetPrice() * 150 / 100
				sum = sum + newPriceItem
			}
		}
	}

	//проверяем, есть ли деньги у пользователя
	if s.users[user].GetCash() < sum {
		return errors.New("денег нет, но вы держитесь")
	}

	s.users[user].CashMinus(sum)
	return nil
}

func (s *Shop) delivery(order []string) {
	for _, item := range order {
		shopItem := s.items[item]
		shopItem.CountMinus()
		s.items[item] = shopItem
	}
}

func (s *Shop) buyLittlePrice(order []string, user string) error {
	//проверяем кэш
	sum, ok := s.checkLittleCache(order)
	if ok { //если там что-то лежит, покупаем (чекаем, вдруг нет денег)
		err := s.payment(sum, user, order)
		if err == nil { //если деньги есть, делаем доставку
			s.delivery(order)
		}
		return err
	}
	//если хэша нет, придется собирать заказ вручную
	sum = s.collectOrder(order, user) //собираем заказ
	err := s.payment(sum, user, order)
	//делаем кэш. Даже если покупка не состоится (денег нет), все равно делаем
	s.writeLittleCache(order, sum)
	//если ошибок нет, списываем товары
	if err == nil {
		s.delivery(order)
	}

	return err
}

func (s *Shop) checkLittleCache(order []string) (int, bool) {
	mapStep, ok := s.littleCache[getItem(order, 0)]
	if !ok {
		return 0, false
	}

	_, ok = mapStep[getItem(order, 1)]
	if !ok {
		return 0, false
	}

	sum := mapStep[getItem(order, 1)]

	return sum, true
}

func (s *Shop) writeLittleCache(order []string, sum int) {
	item := getItem(order, 0)
	mapStep1, ok := s.littleCache[item]
	if !ok {
		mapStep1 = make(map[string]int)
		s.littleCache[item] = mapStep1
	}

	item = getItem(order, 1)
	_, ok = mapStep1[item]
	if !ok {
		mapStep1[item] = sum
	}
}

func (s *Shop) Intrigue(err error, user string) {
	if err == nil {
		fmt.Println("оплата прошла успешно, ")
	} else {
		fmt.Println(err)
	}

	fmt.Println("Баланс пользователя:", s.users[user])

	itemString := "Товары на складе:\n"
	for itemName, item := range s.items {
		itemString += fmt.Sprintf("\titemName %q  \tcount %d  \tprice %d\n", itemName, item.GetCount(), item.GetPrice())
	}
	fmt.Print(itemString)
}

func getItem(order []string, index int) string {
	item := ""
	if index < len(order) {
		item = order[index]
	}
	return item
}

func makeHash(order []string) string {
	sort.Strings(order)

	var hash string
	for _, item := range order {
		hash += item
	}

	return hash
}
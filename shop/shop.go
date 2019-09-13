package shop

import (
	"errors"
	"fmt"
	"sort"
)

const deletekey = -1

type shop struct {
	users       map[string]Useer
	items       map[string]Itemer
	littleCache map[string]map[string]int
	bigCache    map[string]int
}

type shopUpdater struct{}

func (su *shopUpdater) newShopUpdate(s *shop, items []Itemer, users []Useer) {
	//апдейт юзеров
	s.userUpdate(users)
	//апдейт предметов
	s.itemsUpdate(items)
	//сносим мапу, которая с длинными заказами
	s.bigCache = make(map[string]int)
	//редактируем мапу, которая с короткими заказами
	s.editLittleMap(items)
}

func (s *shop) userUpdate(users []Useer) {
	for _, user := range users {
		_, ok := s.users[user.GetName()]
		if ok {
			delete(s.users, user.GetName())
			if user.GetCash() != deletekey {
				s.users[user.GetName()] = user
			}
		} else {
			s.users[user.GetName()] = user
		}
	}
}

func (s *shop) itemsUpdate(items []Itemer) {
	for _, item := range items {
		_, ok := s.items[item.GetName()]
		if ok {
			delete(s.items, item.GetName())

		}
		if item.GetPrice() > deletekey {
			s.items[item.GetName()] = item
		}
	}
}

func (s *shop) editLittleMap(items []Itemer) {
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

func (s *shop) Buy(order []string, user string) error {
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

func (s *shop) checkItems(order []string) error {
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

func (s *shop) buyBigPrice(order []string, user string) error {
	//делаем хэш
	hash := makeHash(order)
	//если хэш найден, проводим оплату
	sum, ok := s.checkBigCache(hash)
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
	s.writeBigCache(hash, sum)
	//если ошибок нет, списываем товары
	if err == nil {
		s.delivery(order)
	}
	return err
}

func (s *shop) checkBigCache(hash string) (int, bool) {
	sum, ok := s.bigCache[hash]

	return sum, ok
}

func (s *shop) collectOrder(order []string, user string) int {
	var sum int

	for _, item := range order {
		sum += s.items[item].GetPrice()
	}

	return sum
}

func (s *shop) payment(sum int, user string, order []string) error {
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

func (s *shop) writeBigCache(hash string, sum int) {
	s.bigCache[hash] = sum
}

func (s *shop) delivery(order []string) {
	for _, item := range order {
		shopItem := s.items[item]
		shopItem.CountMinus()
		s.items[item] = shopItem
	}
}

func (s *shop) buyLittlePrice(order []string, user string) error {
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

func (s *shop) checkLittleCache(order []string) (int, bool) {
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

func (s *shop) writeLittleCache(order []string, sum int) {
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

func (s *shop) Intrigue(err error, user string) {
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

func NewShopPrepare(items []Itemer, users []Useer) *shop {
	shop := newShop(items, users)

	shopUpdater := &shopUpdater{}
	shopUpdater.newShopUpdate(shop, items, users)

	return shop
}

func newShop(items []Itemer, users []Useer) *shop {
	usersMap := make(map[string]Useer, len(users))
	itemsMap := make(map[string]Itemer, len(items))

	for _, user := range users {
		usersMap[user.GetName()] = user
	}
	for _, item := range items {
		itemsMap[item.GetName()] = item
	}

	littleCacheMap := make(map[string]map[string]int)
	bigCacheMap := make(map[string]int)

	return &shop{
		users:       usersMap,
		items:       itemsMap,
		littleCache: littleCacheMap,
		bigCache:    bigCacheMap,
	}
}

func makeHash(order []string) string {
	sort.Strings(order)

	var hash string
	for _, item := range order {
		hash += item
	}

	return hash
}

func getItem(order []string, index int) string {
	item := ""
	if index < len(order) {
		item = order[index]
	}
	return item
}

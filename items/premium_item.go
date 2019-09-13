package items

type PremItem struct {
	Item
}

func (pi PremItem) GetPrice() int {

	return pi.Price * 150 / 100
}

func (pi PremItem) GetStatusItem() bool {

	return true
}

package items

type premItem struct {
	item
}

func NewPremItem(name string, count int, price int) *premItem {
	return &premItem{
		item{
			name:  name,
			count: count,
			price: price,
		},
	}
}

func (pi premItem) GetPrice() int {
	return pi.price * 150 / 100
}

func (pi premItem) GetStatusItem() bool {
	return true
}

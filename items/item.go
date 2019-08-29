package items

type item struct {
	name  string
	count int
	price int
}

func NewItem(name  string, count int, price int) *item {
	return &item{
		name: name,
		count: count,
		price: price,
	}
}

func (i item) GetPrice() int {
	return i.price
}

func (i item) GetName() string {
	return i.name
}

func (i item) GetCount() int {
	return i.count
}

func (i *item) CountMinus() {
	i.count--
}

func (i item) GetStatusItem() bool {
	return false
}

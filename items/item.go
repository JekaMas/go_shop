package items

type Item struct {
	Name  string
	Count int
	Price int
}

func (i Item) GetPrice() int {

	return i.Price
}

func (i Item) GetName() string {
	return i.Name
}

func (i Item) GetCount() int {
	return i.Count
}

func (i *Item) CountMinus() {
	i.Count--
}

func (i Item) GetStatusItem() bool {
	return false
}

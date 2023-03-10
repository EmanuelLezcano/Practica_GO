package invoiceitem

//Item contains of information of a invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

//New returns a new Item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

//Getter of value of Item
/* func (i Item) Value() float64 {
	return i.value
} */

//Type Items
type Items []Item

//Constructor for Items
func NewItems(items ...Item) Items {
	var is Items
	for _, item := range items {
		is = append(is, item)
	}
	return is
}

//Method for type Items
func (is Items) Total() float64 {
	var total float64
	for _, item := range is {
		total += item.value
	}
	return total
}

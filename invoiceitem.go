package invoiceitem

//Item contains of information of
type Item struct {
	id      uint
	product string
	value   float64
}

//New returns a new itme
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value getter of Item.Value
func (i Item) Value() float64 {
	return i.value
}

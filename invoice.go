package invoice

import (
	"github.com/AJRDRGZ/composition/pkg/customer"
	"github.com/AJRDRGZ/composition/pkg/invoiceitem"
)

//Invoice is the struct of a invoice
type Invoice struct {
	country string
	city string
	total float64
	client customer.Customer
	items []invoiceitem.item
}

func new(country, city string, cliente customer.Customer, items[]
invoiceitem.Item) Invoice {
	return Invoice{}
}
// SetTotal is de the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
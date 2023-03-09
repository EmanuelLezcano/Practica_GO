package invoice

import (
	"curso_go/composition/customer"
	"curso_go/composition/invoiceitem"
)

type Invoice struct {
	country string
	city    string
	total   string
	client  customer.Customer
	items   []invoiceitem.Item
}

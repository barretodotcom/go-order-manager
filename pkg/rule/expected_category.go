package rule

import "github.com/barretodotcom/go-credit/internal/model"

type ExpectedCategory struct {
	Category string
}

func (e ExpectedCategory) Satisfy(order model.Order) bool {
	return order.Product.Category == e.Category
}

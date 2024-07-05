package rule

import "github.com/barretodotcom/go-credit/internal/model"

type MinValue struct {
	Value int
}

func (f MinValue) Satisfy(order model.Order) bool {
	return order.Payment.Value > f.Value
}

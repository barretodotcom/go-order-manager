package action

import "github.com/barretodotcom/go-credit/internal/model"

type Discount struct {
	Value int
}

func (a Discount) Execute(order model.Order) model.Order {
	order.Payment.Value = order.Payment.Value - (order.Payment.Value*a.Value)/100

	return order
}

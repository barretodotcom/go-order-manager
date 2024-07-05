package action

import "github.com/barretodotcom/go-credit/internal/model"

type SetLabel struct {
	Label string
}

func (s SetLabel) Execute(order model.Order) model.Order {
	order.ShippingLabels = append(order.ShippingLabels, s.Label)
	return order
}

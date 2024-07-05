package action

import "github.com/barretodotcom/go-credit/internal/model"

type Action interface {
	Execute(order model.Order) model.Order
}

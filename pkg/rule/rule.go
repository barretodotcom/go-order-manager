package rule

import "github.com/barretodotcom/go-credit/internal/model"

type Rule interface {
	Satisfy(order model.Order) bool
}

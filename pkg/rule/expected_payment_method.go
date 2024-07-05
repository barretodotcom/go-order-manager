package rule

import "github.com/barretodotcom/go-credit/internal/model"

type ExpectedPaymentMethod struct {
	PaymentMethod string
}

func (e ExpectedPaymentMethod) Satisfy(order model.Order) bool {
	return e.PaymentMethod == order.Payment.Method
}

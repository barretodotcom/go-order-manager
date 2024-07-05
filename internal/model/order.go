package model

type Order struct {
	Product        Product
	Payment        Payment
	ShippingLabels []string
}

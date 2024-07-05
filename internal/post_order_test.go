package internal

import (
	"testing"

	"github.com/barretodotcom/go-credit/internal/model"
	"github.com/barretodotcom/go-credit/pkg/action"
	"github.com/barretodotcom/go-credit/pkg/rule"
	"github.com/go-playground/assert/v2"
)

func TestFreeShipping(t *testing.T) {
	post := &PostOrder{}

	minValueRule := rule.MinValue{Value: 1000}

	r := []rule.Rule{minValueRule}
	a := action.SetLabel{Label: "free-shipping"}
	post.AddRules(r, a)

	product := model.Product{Description: "Product 1", Category: "T-Shirt"}

	payment := model.Payment{Method: model.MethodCreditCard, Value: 1001}

	order := model.Order{Product: product, Payment: payment, ShippingLabels: []string{}}
	newOrder := post.Execute(order)

	assert.Equal(t, order.ShippingLabels, []string{})
	assert.Equal(t, newOrder.ShippingLabels, []string{"free-shipping"})
}

func TestCreditCard(t *testing.T) {
	post := &PostOrder{}

	expectedPaymentMethodRule := rule.ExpectedPaymentMethod{PaymentMethod: model.MethodCreditCard}
	rules := []rule.Rule{expectedPaymentMethodRule}
	a := action.SetLabel{Label: "free-shipping"}
	post.AddRules(rules, a)

	product := model.Product{Description: "Product 2", Category: "T-Shirt"}

	payment := model.Payment{Method: model.MethodCreditCard}

	order := model.Order{Product: product, Payment: payment, ShippingLabels: []string{}}
	newOrder := post.Execute(order)

	assert.Equal(t, order.ShippingLabels, []string{})
	assert.Equal(t, newOrder.ShippingLabels, []string{"free-shipping"})
}

func TestCategory(t *testing.T) {
	post := &PostOrder{}

	expectedCategoryRule := rule.ExpectedCategory{Category: "T-Shirt"}
	rules := []rule.Rule{expectedCategoryRule}
	action := action.SetLabel{Label: "free-shipping"}
	post.AddRules(rules, action)

	product := model.Product{Description: "Product 3", Category: "T-Shirt"}

	payment := model.Payment{Method: model.MethodCreditCard}

	order := model.Order{Product: product, Payment: payment, ShippingLabels: []string{}}
	newOrder := post.Execute(order)

	assert.Equal(t, order.ShippingLabels, []string{})
	assert.Equal(t, newOrder.ShippingLabels, []string{"free-shipping"})
}

func TestDiscount(t *testing.T) {
	post := &PostOrder{}

	expectedCategoryRule := rule.ExpectedCategory{Category: "T-Shirt"}
	expectedPaymentMethod := rule.ExpectedPaymentMethod{PaymentMethod: model.MethodPix}

	rules := []rule.Rule{expectedCategoryRule, expectedPaymentMethod}
	action := action.Discount{Value: 20}
	post.AddRules(rules, action)

	product := model.Product{Description: "Product 3", Category: "T-Shirt"}

	payment := model.Payment{Method: model.MethodPix, Value: 1000}

	order := model.Order{Product: product, Payment: payment, ShippingLabels: []string{}}
	newOrder := post.Execute(order)

	assert.Equal(t, order.Payment.Value, 1000)
	assert.Equal(t, newOrder.Payment.Value, 800)
}

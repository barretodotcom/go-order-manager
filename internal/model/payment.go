package model

const MethodCreditCard = "credit"
const MethodDebitCard = "debit"
const MethodPix = "pix"

type Payment struct {
	Method string
	Value  int
}

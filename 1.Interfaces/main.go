package main

import (
	"study/payments"
	"study/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {

	payment := payments.NewPaymentModule(methods.NewCrypto())
	payment.Pay("Burger", 5)
	id := payment.Pay("iPhone 17", 900)
	payment.Cancel(id)

	allInfo := payment.AllInfo()
	pp.Print(allInfo)

}

package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Оплата криптой")
	fmt.Printf("Размер оплаты - %d\n", usd)
	return rand.Int()
}
func (c Crypto) Cancel(id int) {
	fmt.Println("Отмена крипто операции - ", id)

}

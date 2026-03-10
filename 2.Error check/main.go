package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type User struct {
	Name    string
	Balance int
}

func Pay(user *User, usd int) error {
	if user.Balance-usd < 0 {
		return errors.New("Недостаточно средств!")
	}
	user.Balance -= usd
	return nil
}

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	if c.Armor-10 <= 0 {
		return 0, errors.New("Мы не стали газовать, чтобы не сломаться!")
	}
	kmch := rand.Intn(150)
	c.Armor -= 10
	return kmch, nil
}

func main() {
	// user := User{
	// 	Name:    "Dastan",
	// 	Balance: 10,
	// }
	// pp.Println("User before: ", user)
	// err := Pay(&user, 22)
	// pp.Println("User after: ", user)

	// if err != nil {
	// 	fmt.Println("Оплата не произведена, причина: ", err.Error())
	// } else {
	// 	fmt.Println("Оплата произведена")
	// }
	// car := Car{
	// 	Armor: 25,
	// }
	// for {
	// 	pp.Println("Car before: ", car)
	// 	kmch, err := car.Gas()
	// 	if err != nil {
	// 		fmt.Println("Error:", err.Error())
	// 		break
	// 	}
	// 	fmt.Println("Speed:", kmch)
	// 	pp.Println("Car after: ", car)
	// }

	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("ПАНИКА!!!", p)
		}
	}()
	slice := []int{1, 2, 3}
	fmt.Println(slice[4])

}

package main

import (
	"fmt"
	"littlelight/orm"
)

type Carro struct {
	Nome   string `json:"nome"`
	Modelo string `json:"modelo"`
}

func main() {
	orm := orm.New[Carro]("revenda")

	myCar, err := orm.SelectByID("12f5c589-06d6-4052-aacb-43fdd17a153d")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v\n", myCar)
	}
}

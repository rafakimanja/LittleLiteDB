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
	myCar, err := orm.SelectByID[Carro]("12f5c589-06d6-4052-aacb-43fdd17a153d", "revenda")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v\n", myCar)
	}
}

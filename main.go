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

	myCar := Carro{Nome: "mustang", Modelo: "2025 5.0 v8"}

	lorm := orm.New()
	lorm.ConnectDB("Concessecionaria")
	lorm.Migrate(myCar)
	err := lorm.Insert(myCar)
	if err != nil {
		fmt.Println(err.Error())
	}
}

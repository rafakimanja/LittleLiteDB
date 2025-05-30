package main

import (
	"fmt"
	"littlelight/orm"
)

type Carro struct {
	Nome   string `json:"nome"`
	Modelo string `json:"modelo"`
}

var lorm *orm.ORM[Carro]

func printTable() {
	mdatas, err := lorm.Select(10, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, item := range mdatas {
		fmt.Printf("%d. %s - [%s | %s]\n", i, item.ID, item.Content.Nome, item.Content.Modelo)
	}
}

func main() {
	lorm = orm.New[Carro]("revenda")

	reformedCar := Carro{Nome: "Camaro", Modelo: "V8 5.0 2020"}

	lorm.Update("a105d31c-563f-4072-9fc3-5e5a17173a2b", reformedCar)
	printTable()
}

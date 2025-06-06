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
	mdatas, err := lorm.Select(10, 1, false)
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
	lorm.Delete("12f5c589-06d6-4052-aacb-43fdd17a153d", false)
}

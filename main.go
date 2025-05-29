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

	mdatas, err := orm.Select(10, 1)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(mdatas) >= 0 {
		for i, item := range(mdatas){
			fmt.Printf("%d. %s - [%s | %s]\n", i, item.ID, item.Content.Nome, item.Content.Modelo)
		}
	} else {
		fmt.Println("Nenhum dado armazenado")
	}
}

package main

import (
	"littlelight/db"
	"littlelight/table"
)

type Funcionario struct {
	nome  string
	cargo string
}

func main() {
	newFunc := Funcionario{nome: "Rafael", cargo: "Desenvolvedor"}
	mydb := db.ConectDB("Empresa")
	table.New(mydb, newFunc)
}
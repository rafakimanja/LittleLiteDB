package main

import "littlelight/orm"

type Funcionario struct {
	Nome  string `json:"nome"`
	Cargo string `json:"cargo"`
}

func main() {
	newFunc := Funcionario{Nome: "Pedro Henrique", Cargo: "Detran"}

	lorm := orm.New()

	lorm.ConectDB("Empresa")
	lorm.Migrate(newFunc)
}

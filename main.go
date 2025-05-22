package main

import (
	"fmt"
	"littlelight/orm"
)

type Funcionario struct {
	Nome  string `json:"nome"`
	Cargo string `json:"cargo"`
}

type Empresa struct {
	Nome string `json:"nome"`
	Cnpj string `json:"cnpj"`
}

func main() {
	lorm := orm.New()

	minhaEmpresa := Empresa{"Dev's S.A", "12.345.678/0001-95"}
	lorm.ConnectDB("Empresa")
	lorm.Migrate(minhaEmpresa)
	err := lorm.Insert(minhaEmpresa)
	if err != nil {
		fmt.Println(err.Error())
	}
}

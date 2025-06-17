package main

import (
	"fmt"

	"github.com/rafakimanja/LittleLiteDB/orm"
)

type Car struct {
	Name   string `json:"name"`
	Model string `json:"model"`
}

var lorm *orm.ORM[Car] = orm.New[Car]("database")

//function to migrate the table 
func MigrationCar(){
	lorm.MigrateTable(Car{})
}

//function for insert new car in database
func InsertCar(){
	newCar := Car{Name: "Mustang", Model: "V8 5.0 2024"}

	lorm.Insert(newCar)
}

//function for return cars in database
func SelectCar(){
	cars, err := lorm.Select(10, 1, false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, item := range cars {
		fmt.Printf("%d. %s - [%s | %s]\n", i+1, item.ID, item.Content.Name, item.Content.Model)
	}
}

//function for get car with ID
func GetCar(id string){
	fmt.Println(lorm.SelectByID(id, false))
}

//function to update car in database
func UpdateCar(id string){
	updateCar := Car{Name: "Ford Mustang", Model: "Eletric 2025"}
	lorm.Update(id, updateCar)
}

//function to delete car in database
func DeleteCar(id string){
	lorm.Delete(id, true)
}
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
	newCar := Car{Name: "Ford Mustang", Model: "5.0 V8 2024"}

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
		fmt.Printf("%d. %s - [%s | %s]\n", i, item.ID, item.Content.Name, item.Content.Model)
	}
}

//function to update car in database
func UpdateCar(id string){
	updateCar := Car{Name: "Ford Mustang", Model: "Eletric 2025"}
	lorm.Update(id, updateCar)
}

//function to delete car in database
func DeleteCar(id string){
	lorm.Delete(id, false)
}

func main(){
	DeleteCar("5bba53de-3e3f-494c-b797-8d3664aaabf9")
}
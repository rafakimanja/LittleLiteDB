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
func UpdateCar(){
	id := "6b0c727f-1fd8-47df-8e4d-0339475b14e1" //id example
	updateCar := Car{Name: "Ford Mustang", Model: "Eletric 2025"}

	lorm.Update(id, updateCar)
}

//function to delete car in database
func DeleteCar(){
	id := "6b0c727f-1fd8-47df-8e4d-0339475b14e1" //id example

	lorm.Delete(id, false)
}
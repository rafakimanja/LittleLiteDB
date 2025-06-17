package controller

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rafakimanja/LittleLiteDB/services"
	"github.com/rafakimanja/LittleLiteDB/types"
)

func (dbc *DBController) writeMetadata(pathFile string, metadata Metadata) error {	
	if services.FSvalidFile(pathFile) {

		_, err := os.ReadFile(pathFile)
		if err != nil {
			return err
		}

		data, err := json.MarshalIndent(metadata, "", "  ")
		if err != nil {
			return err
		}

		return os.WriteFile(pathFile, data, 0644)
	}
	
	return fmt.Errorf("filepath invalid")
}

func (dbc *DBController) readMetadata(filename string) (*Metadata, error){
	var metadata Metadata
	
	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dataBytes, &metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (dbc *DBController) selectTable(filename string) ([]types.Model, error){
	return loadTable(filename)
}

func (dbc *DBController) insertTable(filename string, newItem types.Model) error {
	mdatas, err := loadTable(filename)
	if err != nil {
		return err
	}
	mdatas = append(mdatas, newItem)
	return saveTable(filename, mdatas)
}

func (dbc *DBController) updateTable(filename string, updateItem types.Model) error {
	mdatas, err := loadTable(filename)
	if err != nil {
		return err
	}
	
	for i, data := range(mdatas){
		if data.ID == updateItem.ID {
			mdatas[i] = updateItem
			return saveTable(filename, mdatas)
		}
	}
	return fmt.Errorf("item not found")
}

func (dbc *DBController) deleteTable(filename string, idItem string) error {
	mdatas, err := loadTable(filename)
	if err != nil {
		return err
	}

	position := -1
	for i, item := range(mdatas) {
		if item.ID == idItem {
			position = i
			break
		}
	}

	if position == -1 {
		return fmt.Errorf("item with ID %s not found", idItem)
	}

	mdatas = append(mdatas[:position], mdatas[position+1:]...)
	return saveTable(filename, mdatas)
}

func loadTable(filename string) ([]types.Model, error){
	var mdatas []types.Model

	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dataBytes, &mdatas)
	return mdatas, err
}

func saveTable(filename string, data []types.Model) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, bytes, 0644)
}
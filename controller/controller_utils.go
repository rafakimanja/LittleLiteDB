package controller

import (
	"encoding/json"
	"fmt"
	"littlelite/types"
	"os"
)

func (dbc *DBController) saveMetadata(filename string, metadata Metadata) error {	
	_, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (dbc *DBController) readMetadata(filename string) (*Metadata, error){
	var metadata Metadata
	
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (dbc *DBController) selectTable(filename string) ([]types.Model, error){
	var mdatas []types.Model

	dataBytes, err := os.ReadFile(filename)
	if err != nil{
		return nil, err
	}

	err = json.Unmarshal(dataBytes, &mdatas)
	if err != nil {
		return nil, err
	}

	return mdatas, nil
}

func (dbc *DBController) insertTable(filename string, newItem types.Model) error {
	var mdatas []types.Model
	
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &mdatas)
	if err != nil {
		return err
	}

	mdatas = append(mdatas, newItem)
	ndata, err := json.MarshalIndent(mdatas, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, ndata, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (dbc *DBController) updateTable(filename string, updateItem types.Model) error {
	var mdatas []types.Model

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &mdatas)
	if err != nil {
		return err
	}

	for i, data := range(mdatas){
		if data.ID == updateItem.ID {
			mdatas[i] = updateItem
			
			ndata, err := json.MarshalIndent(mdatas, "", "  ")
			if err != nil {
				return err
			}

			err = os.WriteFile(filename, ndata, 0755)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

func (dbc *DBController) deleteTable(filename string, idItem string) error {
	var mdatas []types.Model

	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataBytes, &mdatas)
	if err != nil {
		return err
	}

	var position int = -1
	for i, item := range(mdatas) {
		fmt.Printf("i. elemento: %d, %v\n", i, item)
		if item.ID == idItem {
			position = i
			break
		}
	}

	if position == -1 {
		return fmt.Errorf("don't find element")
	}

	mdatas = append(mdatas[:position], mdatas[position+1:]...)
	ndata, err := json.MarshalIndent(mdatas, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, ndata, 0755)
	if err != nil {
		return err
	}

	return nil
}
package controller

import (
	"encoding/json"
	"fmt"
	"littlelight/types"
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

func (dbc *DBController) selectTable(filename string, id string) (*types.Model, error){
	var mdatas []types.Model

	dataBytes, err := os.ReadFile(filename)
	if err != nil{
		return nil, err
	}

	err = json.Unmarshal(dataBytes, &mdatas)
	if err != nil {
		return nil, err
	}

	for _, data := range(mdatas) {
		if data.ID == id {
			return &data, nil
		}
	}

	return nil, fmt.Errorf("object not found")
}
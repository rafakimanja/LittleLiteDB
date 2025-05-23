package orm

import (
	"encoding/json"
	"fmt"
	"littlelight/table"
	"os"
)

func (orm *ORM) saveORM(filename string, metadata OrmMeta) error {	
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

func (orm *ORM) insertTable(filename string, newItem table.Model) error {
	var mdatas []table.Model
	
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

func (orm *ORM) selectTable(filename string, id string) (*table.Model, error){
	var mdatas []table.Model

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
package orm

import (
	"encoding/json"
	"littlelight/table"
	"os"
)

func (l *Lorm) insertTable(filename string, newItem table.Model) error {
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
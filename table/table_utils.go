package table

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rafakimanja/LittleLiteDB/services"
)

func (t *Table) writeConfigFile(pathFile string, dados []TableConfig) error {
	if services.FSvalidFile(pathFile){

		_, err := os.ReadFile(pathFile)
		if err != nil {
			return err
		}

		bytes, err := json.MarshalIndent(dados, "", "  ")
		if err != nil {
			return err
		}

		return os.WriteFile(pathFile, bytes, 0644)
	}

	return fmt.Errorf("filepath invalid")
}
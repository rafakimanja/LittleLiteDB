package table

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (t *Table) createFile(name string) error {
	fullPath := filepath.Join(t.path, name)
	file, err := os.Create(fullPath+".json")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString("[]")
	return err
}

func (t *Table) createConfigFile(dados []TableConfig, name string) error {
	fullPath := filepath.Join(t.path, name)
	file, err := os.Create(fullPath+".config.json")
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(dados)
	if err != nil {
		return err
	}
	return nil
}
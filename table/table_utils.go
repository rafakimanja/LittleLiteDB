package table

import (
	"encoding/json"
	"os"
)

func (t *Table) createFile(name string) error {
	file, err := os.Create(t.path + "/" + name)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString("[]")
	return err
}

func (t *Table) createConfigFile(dados []TableConfig, name string) error {
	file, err := os.Create(t.path + "/" + name)
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
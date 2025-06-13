package table

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func (t *Table) createFile(name string) error {
	lowerName := strings.ToLower(name)
	fullPath := filepath.Join(t.path, lowerName)
	file, err := os.Create(fullPath+".json")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString("[]")
	return err
}

func (t *Table) createConfigFile(dados []TableConfig, name string) error {
	lowerName := strings.ToLower(name)
	fullPath := filepath.Join(t.path, lowerName)
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
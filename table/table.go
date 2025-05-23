package table

import (
	"fmt"
	"littlelight/db"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type Table struct {
	path      string
	dbPath    string
	nameTable string
}

type TableConfig struct {
	Field string `json:"field"`
	Types string `json:"type"`
}

// criar o Obj Table (criar o path, a pasta e o objeto)
func New(db *db.Database, table any) (*Table, error) {
	types := reflect.TypeOf(table)
	tbl := &Table{
		path:       buildPath(db.GetPath(), types.Name()),
		dbPath:    db.GetPath(),
		nameTable: types.Name(),
	}
	if !tbl.searchTable() {
		
		err := tbl.buildTable()
		if err != nil {
			return nil, err
		}

		err = tbl.create(table)
		if err != nil {
			return nil, err
		}
	}
	return tbl, nil
}

// procura e valida a tabela em questao
func (t *Table) searchTable() bool {
	//valida se tem um diretorio
	if validPath(t.path) {
		//busca por arquivos
		entries, err := os.ReadDir(t.path)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, entry := range entries {
				//ve se e um arquivo, e se tem .json no nome
				if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
					return true
				}
			}
		}
	}
	return false
}

// criar o arquivo da tabela e da config da tabela em json
func (t *Table) create(table any) error {
	if !validPath(t.dbPath) {
		return fmt.Errorf("database path does not exist")
	}

	if !validPath(t.path) {
		return fmt.Errorf("table path does not exist")
	}

	err := t.createFile(t.nameTable)
	if err != nil {
		return err
	}

	if strings.ToLower(t.nameTable) != "ormmeta" {
		tableModel, err := Init(table)
		if err != nil {
			return err
		}

		configs := t.extractConfig(tableModel.GetContent())
		err = t.createConfigFile(configs, t.nameTable)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func (t *Table) buildTable() error {
	return os.MkdirAll(t.path, 0755)
}

func (t *Table) extractConfig(fields any) []TableConfig {
	configs := []TableConfig{
		{Field: "id", Types: "string"},
		{Field: "created_at", Types: "Time"},
		{Field: "updated_at", Types: "Time"},
		{Field: "deleted_at", Types: "Time"},
	}
	
	typ := reflect.TypeOf(fields)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		configs = append(configs, TableConfig{Field: field.Name, Types: field.Type.Name()})
	}

	return configs
}

func (t *Table) GetPath() string {
	return t.path
}

func (t *Table) GetNameTable() string {
	return t.nameTable
}

func refactorName(name string) string {
	return strings.TrimSpace(strings.ToLower(name))
}

func buildPath(db_path string, name string) string {
	tableName := refactorName(name)
	return filepath.Join(db_path, tableName)
}

func validPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

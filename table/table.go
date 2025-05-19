package table

import (
	"fmt"
	"littlelight/db"
	"os"
	"reflect"
	"strings"
)

type Table struct {
	path       string
	db_path    string
	name_table string
}

type TableConfig struct {
	Field string `json:"field"`
	Types string `json:"type"`
}

// criar o Obj Table (criar o path, a pasta e o objeto)
func New(db *db.Database, table any) *Table {
	types := reflect.TypeOf(table)
	newTable := Table{
		path:    buildPath(db.GetPath(), types.Name()),
		db_path: db.GetPath(),
		name_table: types.Name(),
	}
	newTable.buildTable()
	newTable.create(table)
	return &newTable
}

// criar o arquivo da tabela e da config da tabela em json
func (t *Table) create(table any) {
	if !t.validDB() {
		fmt.Println("Database don't exist!")
		return
	}

	if !t.validPath() {
		fmt.Println("Table don't exists")
		return
	}

	tableModel, err := Init(table)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.createFile(t.name_table+".json")
	configs := t.extractConfig(tableModel.GetContent())
	t.createConfigFile(configs, t.name_table+".config.json")
	fmt.Println("Table created succefull!")
}

func (t *Table) validPath() bool {
	_, err := os.Stat(t.path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (t *Table) validDB() bool {
	_, err := os.Stat(t.db_path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (t *Table) buildTable() bool {
	err := os.MkdirAll(t.path, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err == nil
}

func (t *Table) extractConfig(fields any) []TableConfig {
	var configs []TableConfig
	typesFields := reflect.TypeOf(fields)

	if typesFields.Kind() == reflect.Ptr {
		typesFields = typesFields.Elem()
	}

	for i := 0; i < typesFields.NumField(); i++ {
		field := typesFields.Field(i)
		configs = append(configs, TableConfig{Field: field.Name, Types: field.Type.Name()})
	}
	return configs
}

func refactorName(name string) string {
	name_lower := strings.ToLower(name)
	return strings.TrimSpace(name_lower)
}

func buildPath(db_path string, name string) string {
	tableName := refactorName(name)
	return db_path + "/" + tableName
}

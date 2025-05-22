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
		path:       buildPath(db.GetPath(), types.Name()),
		db_path:    db.GetPath(),
		name_table: types.Name(),
	}
	if !newTable.searchTable() {
		newTable.buildTable()
		newTable.create(table)
	}
	return &newTable
}

// procura e valida a tabela em questao, caso nao exista, cria uma nova
func (t *Table) searchTable() bool {
	//valida se tem um diretorio
	if validPath(t.path){
		//busca por arquivos
		entries, err := os.ReadDir(t.path)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, entry := range entries {
				//ve se e um arquivo, e se tem .json no nome
				if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json"){
					return true
				}
			}
		}
	}
	return false
}

// criar o arquivo da tabela e da config da tabela em json
func (t *Table) create(table any) bool {
	if !validPath(t.db_path) {
		fmt.Println("Database don't exist!")
		return false
	}

	if !validPath(t.path) {
		fmt.Println("Table don't exists")
		return false
	}

	tableModel, err := Init(table)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	t.createFile(t.name_table + ".json")
	configs := t.extractConfig(tableModel.GetContent())
	t.createConfigFile(configs, t.name_table+".config.json")
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
	var configs []TableConfig = []TableConfig{
		{Field: "id", Types: "string"},
		{Field: "created_at", Types: "Time"},
		{Field: "updated_at", Types: "Time"},
		{Field: "deleted_at", Types: "Time"},
	}
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

func (t *Table) GetPath() string {
	return t.path
}

func (t *Table) GetNameTable() string {
	return t.name_table
}

func refactorName(name string) string {
	name_lower := strings.ToLower(name)
	return strings.TrimSpace(name_lower)
}

func buildPath(db_path string, name string) string {
	tableName := refactorName(name)
	return db_path + "/" + tableName
}

func validPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

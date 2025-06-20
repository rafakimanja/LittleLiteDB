package table

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strings"

	"github.com/rafakimanja/LittleLiteDB/db"
	"github.com/rafakimanja/LittleLiteDB/services"
	"github.com/rafakimanja/LittleLiteDB/types"
)

var (
	logger *slog.Logger
)

type Table struct {
	path   string
	dbPath string
	name   string
}

// criar o Obj Table (criar o path, a pasta e o objeto)
func New(db *db.Database, table any) (*Table, error) {
	logger = slog.Default()
	tTypes := reflect.TypeOf(table)
	tbl := &Table{
		path:   services.FSbuildPath(db.GetPath(), tTypes.Name()),
		dbPath: db.GetPath(),
		name:   tTypes.Name(),
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
	if services.FSvalidPath(t.path) {
		//busca por arquivos
		entries, err := os.ReadDir(t.path)
		if err != nil {
			logger.Error(err.Error())
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
	if !services.FSvalidPath(t.dbPath) {
		return fmt.Errorf("database path does not exist")
	}

	if !services.FSvalidPath(t.path) {
		return fmt.Errorf("table path does not exist")
	}

	_, err := services.FSbuildJSONFile(t.path, t.name, ".json")
	if err != nil {
		return err
	}

	if strings.ToLower(t.name) != "metadata" {
		tableModel, err := types.Init(table)
		if err != nil {
			return err
		}

		fp, err := services.FSbuildJSONFile(t.path, t.name, ".config.json")
		if err != nil {
			return err
		}

		configs := t.extractConfig(tableModel.GetContent())
		err = t.writeConfigFile(fp, configs)
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

func (t *Table) GetName() string {
	return t.name
}

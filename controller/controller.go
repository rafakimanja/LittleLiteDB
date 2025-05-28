package controller

import (
	"fmt"
	"littlelight/db"
	"littlelight/services"
	"littlelight/table"
	"littlelight/types"
	"path/filepath"

	"github.com/google/uuid"
)

type DBController struct {
	dbRef    ModelDB
	tableRef types.TableMeta
}

func New() *DBController {
	return &DBController{}
}


func (dbc *DBController) ConnectDB(dbname string) {
	database := db.Connect(dbname)
	dbc.dbRef.db = database
	dbc.dbRef.dbPath = database.GetPath()
}

func (dbc *DBController) Migrate(tableAny any) {
	tb, err := table.New(dbc.dbRef.db, tableAny)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dbc.tableRef.Path = tb.GetPath()
	dbc.tableRef.Name = tb.GetNameTable()
	dbc.tableRef.DataFile = filepath.Join(tb.GetPath(), tb.GetNameTable()+".json")
	dbc.tableRef.ConfigFile = filepath.Join(tb.GetPath(), tb.GetNameTable()+".config.json")

	err = dbc.save()
	if err != nil {
		fmt.Println(err)
	}
}

func (dbc *DBController) Insert(model any) error {
	if dbc.tableRef.Path == "" {
		err := dbc.load()
		if err != nil {
			return err
		}
	}

	dataTable, err := services.ToModel(model)
	if err != nil {
		return err
	}

	err = dbc.insertTable(dbc.tableRef.DataFile, *dataTable)
	if err != nil {
		return err
	}
	return nil
}

func (dbc *DBController) Select(id string) (*types.Model, error) {
	if dbc.tableRef.Path == "" {
		err := dbc.load()
		if err != nil {
			return nil, err
		}
	}

	if id == "" || !dbc.isUUID(id) {
		return nil, fmt.Errorf("this id is invalid")
	}

	data, err := dbc.selectTable(dbc.tableRef.DataFile, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// pega a referencia do DB e busca a o arquivo dbcMeta.json
func (dbc *DBController) load() error {
	tbPath := filepath.Join(dbc.dbRef.dbPath, "metadata")
	pathFile := filepath.Join(tbPath, "metadata.json")
	configs, err := dbc.readMetadata(pathFile)
	if err != nil {
		return err
	}
	dbc.tableRef = configs.Table
	return nil
}

func (dbc *DBController) save() error {

	metadados := Metadata{
		DbPath: dbc.dbRef.dbPath,
		Table: types.TableMeta{
			Name:       dbc.tableRef.Name,
			Path:       dbc.tableRef.Path,
			DataFile:   dbc.tableRef.DataFile,
			ConfigFile: dbc.tableRef.ConfigFile,
		},
		Version: "1.0",
	}

	db := db.Connect(dbc.dbRef.db.GetName())
	tbl, err := table.New(db, metadados)
	if err != nil {
		return fmt.Errorf("error create table metadata")
	}

	err = dbc.saveMetadata(filepath.Join(tbl.GetPath(), tbl.GetNameTable()+".json"), metadados)
	if err != nil {
		return fmt.Errorf("error save metadata to controller")
	}

	return nil
}

func (dbc *DBController) isUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
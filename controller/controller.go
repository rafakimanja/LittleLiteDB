package controller

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/rafakimanja/LittleLiteDB/db"
	"github.com/rafakimanja/LittleLiteDB/internal"
	"github.com/rafakimanja/LittleLiteDB/services"
	"github.com/rafakimanja/LittleLiteDB/table"
	"github.com/rafakimanja/LittleLiteDB/types"

	"github.com/google/uuid"
)

type DBController struct {
	dbRef    ModelDB
	tableRef types.TableMeta
}

func New() *DBController {
	return &DBController{}
}

func (dbc *DBController) ConnectDB(dbname string) error {
	database, err := db.Connect(dbname)
	if err != nil {
		return err
	}
	dbc.dbRef.db = database
	dbc.dbRef.dbPath = database.GetPath()
	return nil
}

func (dbc *DBController) Migrate(tableAny any) error {
	tb, err := table.New(dbc.dbRef.db, tableAny)
	if err != nil {
		return err
	}
	dbc.tableRef.Path = tb.GetPath()
	dbc.tableRef.Name = tb.GetName()
	dbc.tableRef.DataFile = filepath.Join(tb.GetPath(), tb.GetName()+".json")
	dbc.tableRef.ConfigFile = filepath.Join(tb.GetPath(), tb.GetName()+".config.json")

	err = dbc.save()
	if err != nil {
		return err
	}
	return nil
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

func (dbc *DBController) SelectById(id string, flag bool) (*types.Model, error) {
	if dbc.tableRef.Path == "" {
		err := dbc.load()
		if err != nil {
			return nil, err
		}
	}

	if id == "" || !dbc.isUUID(id) {
		return nil, fmt.Errorf("this id is invalid")
	}

	mdata, err := dbc.selectTable(dbc.tableRef.DataFile)
	if err != nil {
		return nil, err
	}

	if !flag {
		for _, data := range(mdata){
			if data.ID == id && data.Deleted_At == nil {
				return &data, nil
			}
		}
	} else {
		for _, data := range(mdata){
			if data.ID == id {
				return &data, nil
			}
		}
	}

	return nil, fmt.Errorf("data not found")
}

func (dbc *DBController) Select(limit int, offset int, flag bool)([]types.Model, error){
	offset = offset-1 //because, inital position is 0 and don't 1
	if dbc.tableRef.Path == "" {
		err := dbc.load()
		if err != nil {
			return nil, err
		}
	}

	mdata, err := dbc.selectTable(dbc.tableRef.DataFile)
	if err != nil {
		return nil, err
	}

	if offset >= len(mdata) {
		return []types.Model{}, nil
	}

	var result []types.Model
	skipped := 0

	if !flag {
		for _, item := range mdata {
			if item.Deleted_At == nil {
				result = append(result, item)
			}
			// Aplica o offset
			if skipped < offset {
				skipped++
				continue
			}
			// Aplica o limit
			if len(result) >= limit {
				break
			}
		}
	} else {
		for _, item := range mdata {
			result = append(result, item)
			if skipped < offset {
				skipped++
				continue
			}
			if len(result) >= limit {
				break
			}
		}
	}

	return result, nil
}

func (dbc *DBController) Update(id string, model any) error {
	if dbc.tableRef.Path == "" {
		err := dbc.load()
		if err != nil {
			return err
		}
	}

	newModel, err := services.ToModel(model)
	if err != nil {
		return err
	}

	oldModel, err := dbc.SelectById(id, false)
	if err != nil {
		return err
	}

	oldModel.Content = newModel.Content
	oldModel.Updated_At = time.Now()

	err = dbc.updateTable(dbc.tableRef.DataFile, *oldModel)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (dbc *DBController) Delete(id string, flag bool) error {
	if dbc.tableRef.Path == "" {
		err := dbc.load()
		if err != nil {
			return err
		}
	}

	if !flag {
		element, err := dbc.SelectById(id, false)
		if err != nil {
			return err
		}

		now := time.Now()
		element.Deleted_At = &now

		err = dbc.updateTable(dbc.tableRef.DataFile, *element)
		return err
	} else {
		err := dbc.deleteTable(dbc.tableRef.DataFile, id)
		return err
	}
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
		Version: internal.VERSION,
	}

	db, err := db.Connect(dbc.dbRef.db.GetName())
	if err != nil {
		return err
	}
	tbl, err := table.New(db, metadados)
	if err != nil {
		return fmt.Errorf("error create table metadata")
	}

	fp, err := services.FSbuildJSONFile(tbl.GetPath(), tbl.GetName(), ".json")
	if err != nil {
		return fmt.Errorf("error create metadata file")
	}

	err = dbc.writeMetadata(fp, metadados)
	if err != nil {
		return fmt.Errorf("error save metadata")
	}

	return nil
}

func (dbc *DBController) isUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
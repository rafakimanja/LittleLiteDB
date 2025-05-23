package orm

import (
	"fmt"
	"littlelight/db"
	"littlelight/table"
	"path/filepath"

	"github.com/google/uuid"
)

type ORM struct {
	dbRef          ModelDB
	tableRef       ModelTable
	pathFile       string
	pathConfigFile string
}

func New() *ORM {
	return &ORM{}
}

func (orm *ORM) ConnectDB(dbname string) {
	database := db.Connect(dbname)
	orm.dbRef.db = database
	orm.dbRef.dbPath = database.GetPath()
}

func (orm *ORM) Migrate(tableAny any) {
	tb, err := table.New(orm.dbRef.db, tableAny)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	orm.tableRef.table = tb
	orm.tableRef.tablePath = tb.GetPath()

	orm.pathFile = filepath.Join(orm.tableRef.tablePath, orm.tableRef.table.GetNameTable()+".json")
	orm.pathConfigFile = filepath.Join(orm.tableRef.tablePath, orm.tableRef.table.GetNameTable()+".config.json")

	err = orm.save()
	if err != nil {
		fmt.Println(err)
	}
}

func (orm *ORM) Insert(model any) error {
	if orm.tableRef.table == nil {
		return fmt.Errorf("table is not initializede")
	}

	dataTable, err := orm.convertData(model)
	if err != nil {
		return err
	}

	err = orm.insertTable(orm.pathFile, *dataTable)
	if err != nil {
		return err
	}
	return nil
}

func (orm *ORM) Select(id string) (*table.Model, error) {

	if orm.tableRef.table == nil {
		return nil, fmt.Errorf("table is not initializede")
	}

	if id == "" || !orm.isUUID(id) {
		return nil, fmt.Errorf("this id is invalid")
	}

	data, err := orm.selectTable(orm.pathFile, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (orm *ORM) save() error {

	metadados := OrmMeta{
		DbPath: orm.dbRef.dbPath,
		Table: TableMeta{
			Name: orm.tableRef.table.GetNameTable(),
			Path: orm.tableRef.table.GetPath(),
			DataFile: orm.pathFile,
			ConfigFile: orm.pathConfigFile,
		},
		Version: "1.0",
	}

	dbOrm := db.Connect(orm.dbRef.db.GetName())
	tblOrm, err := table.New(dbOrm, metadados)
	if err != nil {
		return fmt.Errorf("error create table metadata orm")
	}

	err = orm.saveORM(filepath.Join(tblOrm.GetPath(), tblOrm.GetNameTable()+".json"), metadados)
	if err != nil {
		return fmt.Errorf("error save metadata orm")
	}

	return nil
}

func (orm *ORM) convertData(data any) (*table.Model, error) {
	return table.Init(data)
}

func (orm *ORM) isUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

package orm

import (
	"fmt"
	"littlelight/db"
	"littlelight/table"
)

type ORM struct {
	dbRef    ModelDB
	tableRef ModelTable
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
}

func (orm *ORM) Insert(model any) error {
	if orm.tableRef.table == nil {
		return fmt.Errorf("table is not initializede")
	}

	pathFile := orm.tableRef.tablePath + "/" + orm.tableRef.table.GetNameTable() + ".json"
	
	dataTable, err := orm.convertData(model)
	if err != nil {
		return err
	}

	err = orm.insertTable(pathFile, *dataTable)
	if err != nil {
		return err
	}
	return nil
}

func (orm *ORM) convertData(data any) (*table.Model, error) {
	return table.Init(data)
}

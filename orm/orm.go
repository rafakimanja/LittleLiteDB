package orm

import (
	"fmt"
	"littlelight/db"
	"littlelight/table"
)


type Lorm struct {
	db_ref     ModelDB
	tables_ref ModelTable
}

func New() *Lorm {
	return &Lorm{}
}

func (l *Lorm) ConectDB(dbname string) {
	db := db.ConectDB(dbname)
	l.db_ref.db = db
	l.db_ref.db_path = db.GetPath()
}

func (l *Lorm) Migrate(tableAny any) {
	tb := table.New(l.db_ref.db, tableAny)
	l.tables_ref.table = tb
	l.tables_ref.table_path = tb.GetPath()

	if l.Insert(tableAny){
		fmt.Println("Insert data succefull")
	} else {
		fmt.Println("Error insert data in table")
	}
}

func (l *Lorm) Insert(content any) bool {
	pathFile := l.tables_ref.table_path + "/" + l.tables_ref.table.GetNameTable() + ".json"
	dataTable, err := l.convertData(content)
	if err != nil {
		return false
	}
	fmt.Println(dataTable)
	err = l.insertTable(pathFile, *dataTable)
	return err == nil
}

func (l *Lorm) convertData(data any) (*table.Model, error) {
	return table.Init(data)
}
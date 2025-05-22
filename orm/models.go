package orm

import (
	"littlelight/db"
	"littlelight/table"
)

type ModelDB struct {
	db      *db.Database
	db_path string
}

type ModelTable struct {
	table      *table.Table
	table_path string
}

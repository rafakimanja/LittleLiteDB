package orm

import (
	"littlelight/db"
	"littlelight/table"
)

type ModelDB struct {
	db     *db.Database
	dbPath string
}

type ModelTable struct {
	table     *table.Table
	tablePath string
}

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

type OrmMeta struct {
	DbPath  string    `json:"db_path"`
	Table   TableMeta `json:"table"`
	Version string    `json:"version"`
}

type TableMeta struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	DataFile   string `json:"data_file"`
	ConfigFile string `json:"config_file"`
}
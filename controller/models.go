package controller

import (
	"littlelite/db"
	"littlelite/types"
)

type ModelDB struct {
	db     *db.Database
	dbPath string
}

type Metadata struct {
	DbPath  string          `json:"db_path"`
	Table   types.TableMeta `json:"table"`
	Version string          `json:"version"`
}
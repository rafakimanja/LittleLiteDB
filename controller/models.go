package controller

import (
	"littlelight/db"
	"littlelight/types"
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
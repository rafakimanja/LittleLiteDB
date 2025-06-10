package controller

import (
	"github.com/rafakimanja/LittleLiteDB/db"
	"github.com/rafakimanja/LittleLiteDB/types"
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
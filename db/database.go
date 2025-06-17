package db

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/rafakimanja/LittleLiteDB/services"
)

var (
	logger *slog.Logger
)

type Database struct {
	path string
	name string
}

func Connect(name string) (*Database, error) {
	logger = slog.Default()
	conector := Database{path: buildPath(name), name: strings.ToLower(name)}
	if err := conector.searchDB(); err != nil {
		return nil, err
	}
	return &conector, nil
}

func (d *Database) searchDB() error {
	if !services.FSvalidPath(d.path){
		if flag := d.buildDB(); !flag {
			return fmt.Errorf("failed to create database directory")
		}
		logger.Info("database created successfully")
	}
	return nil
}

func (d *Database) buildDB() bool {
	err := os.MkdirAll(d.path, 0755)
	if err != nil {
		logger.Error(err.Error())
	}
	return err == nil
}

func (d *Database) GetPath() string {
	return d.path
}

func (d *Database) GetName() string {
	return d.name
}

func buildPath(name string) string {
	pt, err := services.FSgetRootPath()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	dbName := services.FSrefactorName(name)
	return filepath.Join(pt, "lldb", dbName)
}

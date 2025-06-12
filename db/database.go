package db

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
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
	if !d.validPath() {
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

func (d *Database) validPath() bool {
	_, err := os.Stat(d.path)
	return !os.IsNotExist(err)
}

func (d *Database) GetPath() string {
	return d.path
}

func (d *Database) GetName() string {
	return d.name
}

func refactorName(name string) string {
	nameLower := strings.ToLower(name)
	return strings.TrimSpace(nameLower)
}

func buildPath(name string) string {
	dbName := refactorName(name)
	return "./LLDB/" + dbName
}

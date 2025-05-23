package db

import (
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

func Connect(name string) *Database {
	logger = slog.Default()
	conector := Database{path: buildPath(name), name: strings.ToLower(name)}
	conector.searchDB()
	return &conector
}

func (d *Database) searchDB() {
	if !d.validPath() {
		if d.buildDB() {
			logger.Info("database make's succeful!")
		} else {
			logger.Error("erro in build database")
			os.Exit(1)
		}
	}
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
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (d *Database) GetPath() string {
	return d.path
}

func (d *Database) GetName() string {
	return d.name
}

func refactorName(name string) string {
	name_lower := strings.ToLower(name)
	return strings.TrimSpace(name_lower)
}

func buildPath(name string) string {
	dbName := refactorName(name)
	return "./LLDB/" + dbName
}

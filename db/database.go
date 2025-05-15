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
}

func ConectDB(name string) *Database {
	logger = slog.Default()
	conector := Database{path: buildPath(name)}
	conector.searchDB()
	return &conector
}

func (d* Database) searchDB(){
	if !d.validPath() {
		if d.buildDB() {
			logger.Info("database make's succeful!")
		} else {
			logger.Error("erro in build database")
			os.Exit(1)
		}
	} else {
		logger.Info("database it's exist")
	}	
}

func (d *Database) buildDB() bool {
	err := os.MkdirAll(d.path, 0755)
	if err != nil {
		logger.Error(err.Error())
	}
	return err == nil
}

func (d* Database) validPath() bool {
	_, err := os.Stat(d.path)
    if os.IsNotExist(err) {
        return false
    }
	return true
}

func refactorNameDB(name string) string {
	nameDB := strings.ToLower(name)
	return strings.TrimSpace(nameDB)
}

func buildPath(name string) string {
	dbName := refactorNameDB(name)
	return "./littlelight/"+dbName
}
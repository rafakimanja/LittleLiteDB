package services

import (
	"os"
	"path/filepath"
	"strings"
)

func FSrefactorName(name string) string {
	return strings.TrimSpace(strings.ToLower(name))
}

func FSbuildPath(db_path string, name string) string {
	tableName := FSrefactorName(name)
	return filepath.Join(db_path, tableName)
}

func FSvalidPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
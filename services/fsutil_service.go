package services

import (
	"fmt"
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

func FSgetRootPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil // achou a raiz
		}

		// sobe um nível
		pai := filepath.Dir(dir)
		if pai == dir {
			break // chegou no topo
		}
		dir = pai
	}
	return "", os.ErrNotExist
}

func FSvalidPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func FSvalidFile(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func FSbuildJSONFile(dirPath string, name string, suffix string) (string, error) {
	lowerName := strings.ToLower(name)
	fullPath := filepath.Join(dirPath, lowerName+suffix)
	if FSvalidPath(dirPath) {	
		file, err := os.Create(fullPath)
		if err != nil {
			return "", err
		}

		defer file.Close()

		_, err = file.WriteString("[]")
		if err != nil {
			return "", err
		} else {
			return fullPath, nil
		}
	} else {
		return "", fmt.Errorf("dirPath invalid")
	}
}
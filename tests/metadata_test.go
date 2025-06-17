package tests

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"unicode"

	"github.com/rafakimanja/LittleLiteDB/orm"
)

var (
	lorm *orm.ORM[Tarefas]
	dbName = "database"
	metadataPath = "./LLDB/" + dbName + "/metadata"
)

func TestMain(m *testing.M) {
	fmt.Println("preparando o ambiente...")
	lorm = orm.New[Tarefas](dbName)
	lorm.MigrateTable(Tarefas{})

	exitcode := m.Run()

	fmt.Println("finalizando o ambiente...")
	os.RemoveAll("./LLDB")
	os.Exit(exitcode)
}


//verifica se o arquivo metadata foi criado com o nome em minusculo
func TestMetadataFile(t *testing.T) {
	entries, err := os.ReadDir(metadataPath)
	if err != nil {
		t.Fatalf("Erro ao ler diretório %s: %v", metadataPath, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // ignora subdiretórios
		}

		nomeArquivo := entry.Name()

		if strings.IndexFunc(nomeArquivo, unicode.IsUpper) != -1 {
			t.Errorf("Arquivo '%s' contém letras maiúsculas", nomeArquivo)
		}
	}
}

//verifica se existe mais de um arquivo dentro da pasta metadata
func TestMetadataFiles(t *testing.T) {
	entries, err := os.ReadDir(metadataPath)
	if err != nil {
		t.Fatalf("Erro ao abrir o diretório %s: %v", metadataPath, err)
	}

	contadorArquivos := 0

	for _, entry := range entries {
		if !entry.IsDir() {
			contadorArquivos++
		}
	}

	if contadorArquivos > 1 {
		t.Errorf("Esperava no máximo 1 arquivo no diretório '%s', mas encontrou %d", metadataPath, contadorArquivos)
	}
}
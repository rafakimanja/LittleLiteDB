package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/rafakimanja/LittleLiteDB/orm"
)

type Tarefas struct {
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Status    bool   `json:"status"`
}

var (
	lorm   *orm.ORM[Tarefas]
	dbName = "database"
	path = "./LLDB/" + dbName + "/tarefas/metadata.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Preparando o ambiente...")
	lorm = orm.New[Tarefas](dbName)
	lorm.MigrateTable(Tarefas{})

	exitcode := m.Run()

	fmt.Println("Finalizando o ambiente...")
	//os.RemoveAll("./LLDB")
	os.Exit(exitcode)
}

func TestArqMetadata(t *testing.T) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		t.Fatalf("metadata.json não foi criado em LLDB/%s/tarefas", dbName)
	} else if err != nil {
		t.Fatalf("erro ao acessar %s: %v", path, err)
	}
}

func TestConteudoMetadataJson(t *testing.T) {
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("erro ao ler o arquivo %s: %v", path, err)
	}

	if len(data) == 0 {
		t.Errorf("o arquivo %s está vazio", path)
	}
}

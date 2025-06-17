package tests

import (
	"fmt"
	"testing"

	"github.com/rafakimanja/LittleLiteDB/services"
)

func TestWorkPath(t *testing.T) {
	dir, err := services.FSgetRootPath()
	if err != nil {
		t.Fatal(err.Error())
	}
	
	fmt.Println("Diretorio retornado: ", dir)

	if dir != `c:\Users\Rafae\OneDrive\Documentos\codigos\Go\LittleLiteDB` {
		t.Error("Diretorio esta errado")
	}
}
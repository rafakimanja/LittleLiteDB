package tests

type Tarefas struct {
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Status    bool   `json:"status"`
}
# LittleLiteDB

O LittleLiteDB é um banco de dados não relacional desenvolvido em Golang. Ele tem o propósito de ser leve e prático, seu caso de uso, é o de testes ou pequenos apps,
ele salva os dados em arquivos ```json```, portanto, não é recomendado para dados sensíveis ou produção, pois, não existe uma camada de segurança dos dados.

## 01. Instalação

Para instalar basta realizar:
```bash

go get github.com/rafakimanja/LittleLiteDB
```

***

## 02. Funcionalidades

O LLDB vem com algumas funcionalidades que ajudam o uso e a implementação deste banco de dados no seu sistema, ele funciona com base em tabelas que usam structs como modelo para serem criada. 
Qualquer dado salvo, é salvo em um modelo específico que contém propriedades úteis na hora de implementar funcionalidades:
```go
type Model struct {
	ID         string     `json:"id"` //e usado UUID para facilitar a manipulação dos dados
	Content    any        `json:"content"` //a struct usada como modelo
	Created_At time.Time  `json:"created_at"` //campo de tempo para saber quando o registro foi criado
	Updated_At time.Time  `json:"updated_at"` //campo de tempo para saber quando o registro foi atualizado
	Deleted_At *time.Time `json:"deleted_at"` //campo de tempo para definir uma exclusao logica do registro
}
```

### Soft Delete

O LLDB possui soft delete, ou seja, um campo que indica a data de quando o usuario optou por apagar aquele registro, sem apaga-lo de fato. Isso é muito útil quando se precisa manter uma consistência de dados e evitar erros.

### ORM
Por padrão, o LLDB vem com seu próprio ORM, o que facilita na criação e implementação do banco de dados no seu sistema. O ORM vem com 5 funcionalidades básicas para a manipulação do banco:

| Função | Descrição |
:-------|:---------
`Select(limit int, offset int, delete bool)`| Retorna uma quantidade de elementos dentro do _limit_ e _offset_ definidos
`SelectByID(id string, delete bool)` | Retorna um elemento com base no ID
`Insert(data any)` | Insere um novo registro na tabela
`Update(id string, data any)` | Atualiza um registro na tabela de acordo com o ID
`Delete(id string, delete bool)` | Deleta um registro no BD de acordo com ID

Campos que contém a opção de variáveis booleanas `delete`, refere-se a dados que foram apagados por _soft delete_, passando o parâmetro como `false` estes dados não serão exibidos ou alterados, caso queira vizualiza-los ou apaga-los de fato, basta passar `true` como parâmetro. 

***

## 03. Uso

Para usar o LLDB, muitas coisas foram abstraídas para tornar o uso o mais simples possível, com apenas um método é possível criar um novo banco de dados e tabelas, o ORM pussí um método `New[T any](db string)` que retorna a instância de um novo ORM e recebe por parâmetros uma struct `T` e o banco de dados `(db string)`. Por baixo dos panos o LLDB se conecta com o bd e a tabela correspondentes, caso elas não existam, elas são criadas na hora, por tanto, não é necessário criar funções de conexões com banco de dados, criação de tabelas ou algo do tipo.

```go
func New[T any](db string) *ORM[T] {}
```

* Exemplo de uso criando um ORM para uma tabela `Car` e um banco de dados com nome `database`:
```go
type Car struct {
	Name   string `json:"name"`
	Model string `json:"model"`
}

var lorm *orm.ORM[Car] = orm.New[Car]("database")
```

### Observações

Existem alguns pontos importântes na hora de criar e manipular suas tabelas que são essenciais para funcionar corretamente.

* O LLBD suporta diversas tabelas no mesmo banco de dados, no entanto, cada instância de um ORM lida apenas com 1 tabela, então, cada tabela deve ter seu ORM.
* Como os dados são salvos em `JSON`, é obrigatório passar as tags json nas structs, exemplo: `json:"name"`.
* É necessário deixar os atributos da struct como públicos, ou seja, primeira letra maiúscula, exemplo: `Name string`.


Temos disponibilizado um exemplo completo [aqui](/examples/car_example.go)

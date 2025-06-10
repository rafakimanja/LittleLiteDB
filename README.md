# LittleLiteDB

O LittleLiteDB é um banco de dados não relacional desenvolvido em Golang. Ele tem o propósito de ser leve e prático, ser usado em casos de testes ou pequenos apps,
salva os dados em arquivos ```json```, portanto, não é recomendado para dados sensíveis ou produção, pois, não existe uma camada de segurança dos dados.

## 01. Instalação

Para instalar o LLDB basta realizar:
```bash

go get github.com/rafakimanja/LittleLiteDB
```

***

## 02. Funcionalidades

O LLDB vem com algumas funcionalidades que ajudam o uso e a implementação deste banco de dados no seu sistema, este banco de dados funciona com base em tabelas, que usam **structs** como modelo para serem criada. 
Todo dado salvo no LLDB, é salvo em um modelo específico que contém propriedades úteis na hora de implementar funcionalidades:
```
type Model struct {
	ID         string     `json:"id"` //e usado UUID para facilitar a manipulação dos dados
	Content    any        `json:"content"` //a struct usada como modelo
	Created_At time.Time  `json:"created_at"` //campo de tempo para saber quando o registro foi criado
	Updated_At time.Time  `json:"updated_at"` //campo de tempo para saber quando o registro foi atualizado
	Deleted_At *time.Time `json:"deleted_at"` //campo de tempo para definir uma exclusao logica do registro
}
```

### Soft Delete

O LLDB possui soft delete, ou seja, um campo que indica a data de quando o usuario optou por apagar aquele registro, sem apaga-lo de fato. Isso é muito útil quando se precisa manter uma consistência de dados e efitar erros.

### ORM
Por padrão, o LLDB vem com seu próprio ORM o que facilita na criação na implementação do banco de dados no seu sistema. O ORM vem com 5 funcionalidades básicas para a manipulação do banco:

| Função | Descrição |
:-------|:---------
`Select(limit int, offset int, delete bool)`| Retorna os elementos da tabela dentro do _limit_ e _offset_ definidos
`SelectByID(id string, delete bool)` | Retorna um elemento com base no ID
`Insert(data any)` | Insere um novo dado na tabela
`Update(id string, data any)` | Atualiza um dado na tabela de acordo com o ID
`Delete(id string, delete bool)` | Deleta um dado no BD de acordo com ID

Campos que contém a opção de variáveis booleanas `delete`, refere-se a dados que foram apagados por _soft delete_, passando o parâmetro como `false` estes dados não serão exibidos ou alterados, caso queira vizualiza-los ou apaga-los de fato, basta passar `true` como parâmetro. 

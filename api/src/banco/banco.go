package banco

// pacote banco é responsável por estabelecer a conexão com o banco de dados e fornecer uma função para conectar-se a ele. A função Conectar() retorna um ponteiro para sql.DB, que é a estrutura usada para interagir com o banco de dados, e um erro caso ocorra algum problema durante a conexão. Essa função será utilizada em outras partes da aplicação para realizar operações de leitura e escrita no banco de dados.

import (
	"api/src/config"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db,nil
}

// nesse arquivo criamos a função Conectar() que é responsável por estabelecer a conexão com o banco de dados 
// e fornecer uma função para conectar-se a ele. A função Conectar() retorna um ponteiro para sql.DB, 
// que é a estrutura usada para interagir com o banco de dados, e um erro caso ocorra algum problema durante 
// a conexão. Essa função será utilizada em outras partes da aplicação para realizar operações de leitura e escrita no banco de dados.
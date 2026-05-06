package repositorios

// o pacote repositorios é responsável por implementar as funções que serão executadas para cada rota da API

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados e retorna o ID do usuário criado ou um erro caso ocorra algum problema durante a inserção.
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}

// nesse arquivo criamos a estrutura usuarios que possui um campo db do tipo *sql.DB,
// que é a estrutura usada para interagir com o banco de dados.
// A função NovoRepositorioDeUsuarios recebe um ponteiro para sql.DB e retorna um ponteiro para a
// estrutura usuarios, permitindo que as funções de criação, leitura, atualização e exclusão de usuários
// possam acessar o banco de dados através do campo db.

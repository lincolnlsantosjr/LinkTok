package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// controllers é o pacote responsável por implementar as funções que serão executadas para cada rota da API,
// como criar, listar, atualizar e deletar usuários.

// CriarUsuario cria um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// ler o corpo da requisição e converter para a struct usuario
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	// criar uma variável do tipo usuario para armazenar os dados do usuário que serão criados
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}
	// conectar ao banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	// criar um repositório de usuários para acessar as funções de criação, leitura, atualização e exclusão de usuários
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))
}

// BuscarUsuarios busca todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

// BuscarUsuario busca um usuário específico no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário"))
}

// AtualizarUsuario atualiza um usuário específico no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// DeletarUsuario deleta um usuário específico no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}

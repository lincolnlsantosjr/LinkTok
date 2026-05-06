package rotas

// rotas é o pacote responsável por definir as rotas da API, ou seja,
// os endpoints que serão acessados pelos clientes para realizar operações relacionadas aos usuários,
// como criar, listar, atualizar e deletar usuários. Cada rota é definida com seu URI, método HTTP, função a
//  ser executada e se requer autenticação ou não.

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuario/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}

// essa variável rotasUsuarios é uma slice de Rota que contém as rotas relacionadas aos usuários da API.
// Cada rota é definida com seu URI, método HTTP, função a ser executada e se requer autenticação ou não.
// o pacote controllers será responsável por implementar as funções que serão executadas para cada rota,
//  como criar, listar, atualizar e deletar usuários.

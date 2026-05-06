package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux" // Importa o pacote mux para criar as rotas da API
)

// Gerar é responsável por criar as rotas da API
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
	// a função Gerar é responsável por criar um novo router mux 
	// e chamar a função Configurar para configurar as rotas da API.
}

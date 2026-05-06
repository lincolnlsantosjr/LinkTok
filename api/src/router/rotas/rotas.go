package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// a func Configurar é responsável por configurar as rotas da API, ou seja,
// registrar cada rota definida na variável rotasUsuarios no router mux.
// Ela recebe um router mux como parâmetro e retorna o mesmo router com as rotas configuradas.
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
// nessa função nós iteramos sobre a slice de rotasUsuarios e para cada rota, 
// registramos a função associada à rota no router mux usando o método HandleFunc.
}


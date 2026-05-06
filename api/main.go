package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	r := router.Gerar() // Gerar as rotas da API

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}

// 1º - rotas/rotas.go (struct Rota e função Configurar)
// 2º - rotas/usuarios.go (definição das rotas relacionadas aos usuários com base na struct Rota (receberá uma função do controller, método HTTP, URI e se requer autenticação ou não))
// 3º - controllers/usuarios.go (funções que serão executadas para cada rota relacionada aos usuários)
// 4º - router/router.go (função Gerar que chama a função Configurar para configurar as rotas da API)
// 5º - main.go (função main que chama a função Gerar para gerar as rotas da API e iniciar o servidor HTTP)
// 6º - modelos/usuario.go (definição da struct Usuario e funções(métodos) relacionadas à manipulação de dados de usuários, como criação, leitura, atualização e exclusão)
// 7 - criado o banco de dados no dbeaver com a estrutura que ta na pasta /sql/sql.sql (tabela usuarios com os campos id, nome, nick, email, senha e criado_em)
// 8 - config/config.go (criada a função Carregar() que é responsável por carregar as variáveis de ambiente do sistema, como a string de conexão com o banco de dados, e outras configurações necessárias para o funcionamento da API. A função Carregar() é chamada no início da função main() para garantir que as variáveis de ambiente sejam carregadas antes de qualquer outra parte da aplicação tentar acessá-las.)
// 9 - criado o arquivo .env na raiz do projeto para armazenar as variáveis de ambiente, como a string de conexão com o banco de dados e a porta em que a API irá rodar. 
// 10 - banco/banco.go (criada a função Conectar() que é responsável por estabelecer a conexão com o banco de dados e fornecer uma função para conectar-se a ele. A função Conectar() retorna um ponteiro para sql.DB, que é a estrutura usada para interagir com o banco de dados, e um erro caso ocorra algum problema durante a conexão. Essa função será utilizada em outras partes da aplicação para realizar operações de leitura e escrita no banco de dados.)
// 11 - repositorios/usuarios.go (criada a estrutura Usuarios que possui um campo db do tipo *sql.DB, que é a estrutura usada para interagir com o banco de dados. A função NovoRepositorioDeUsuarios recebe um ponteiro para sql.DB e retorna um ponteiro para a estrutura Usuarios, permitindo que as funções de criação, leitura, atualização e exclusão de usuários possam acessar o banco de dados através do campo db.)
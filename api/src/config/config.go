package config

// definir configuração das variaveis de ambiente, como a string de conexão com o banco de dados

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// var são variáveis globais que podem ser acessadas em todo o pacote config. Neste caso, StringConexaoBancoDeDados é uma variável que armazenará a string de conexão com o banco de dados, permitindo que outras partes da aplicação acessem essa informação para se conectar ao banco de dados.
var (
	// StringConexaoBancoDeDados é a string de conexão com o banco de dados, que será carregada a partir das variáveis de ambiente ou de um arquivo de configuração.
	StringConexaoBanco = ""

	// Porta é a porta em que a API irá rodar, que pode ser configurada a partir das variáveis de ambiente ou de um arquivo de configuração. O valor padrão é 0, o que significa que a API irá escolher uma porta disponível automaticamente.
	Porta = 0
)

// Carregar é responsável por carregar as variáveis de ambiente do sistema, como a string de conexão com o banco de dados, e outras configurações necessárias para o funcionamento da API.
func Carregar() {
	var erro error

	// godotenv.Load() é uma função da biblioteca godotenv que carrega as variáveis de ambiente a partir de um arquivo .env. Se ocorrer um erro ao carregar as variáveis de ambiente, o log.Fatal() é chamado para registrar o erro e encerrar a aplicação.
	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente: ", erro)
	}

	// abaixo temos a atribuição das variáveis de ambiente às variáveis globais do pacote config. A função os.Getenv() é usada para acessar as variáveis de ambiente, e a função strconv.Atoi() é usada para converter a string da porta em um inteiro.
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	// StringConexaoBanco é formatada usando fmt.Sprintf() para criar a string de conexão com o banco de dados, utilizando as variáveis de ambiente para obter o usuário, senha e nome do banco de dados. A string de conexão é construída no formato esperado pelo driver do banco de dados, incluindo parâmetros adicionais como charset, parseTime e loc.
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=uft8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"))
}
// nesse arquivo criamos a função Carregar() que é responsável por carregar as variáveis de ambiente do sistema, 
// como a string de conexão com o banco de dados, e outras configurações necessárias para o funcionamento da API. 
// A função Carregar() é chamada no início da função main() para garantir que as variáveis de ambiente sejam 
// carregadas antes de qualquer outra parte da aplicação tentar acessá-las.



// criamos o arquivo .env na raiz do projeto para armazenar as variáveis de ambiente, como a string de conexão
// com o banco de dados e a porta em que a API irá rodar. O conteúdo do arquivo .env pode ser algo como:
// STRING_CONEXAO_BANCO=usuario:senha@tcp(localhost:3306)/nome_do_banco
// PORTA=5000

// após isso, usando o pacote os, podemos acessar essas variáveis de ambiente em nosso código Go
// usando os.Getenv("NOME_DA_VARIAVEL"). Por exemplo, para acessar a string de conexão com o banco de dados,
// podemos usar os.Getenv("STRING_CONEXAO_BANCO") e atribuir esse valor à variável StringConexaoBanco.
// Da mesma forma, para acessar a porta em que a API irá rodar, podemos usar os.Getenv("PORTA") e converter esse
// valor para um inteiro antes de atribuí-lo à variável Porta.

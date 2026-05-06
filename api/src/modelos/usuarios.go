package modelos
// define a struct Usuario e funções(métodos) relacionadas à manipulação de dados de usuários, como criação, leitura, atualização e exclusão

import "time"
// struct Usuario representa um usuário do sistema
type Usuario struct{
	ID       uint64 `json:"id,omitempty"` // ID do usuário, omitido na resposta JSON se for zero (se vier zerado, não passa o campo na resposta)
	Nome     string `json:"nome,omitempty"` 
	Nick	 string `json:"nick,omitempty"` 
	Email    string `json:"email,omitempty"` 
	Senha    string `json:"senha,omitempty"` 
	CriadoEm time.Time `json:"criadoEm,omitempty"` 
}
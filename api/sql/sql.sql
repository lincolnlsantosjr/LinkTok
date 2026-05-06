-- aqui ficarão alguns comando para facilitar a vida no desenvolvimento da API, como criação de tabelas, inserção de dados, etc.

CREATE DATABASE IF NOT EXISTS linktok;
USE linktok; --antes de chamar funções no DB, sempre usar esse comando pra referenciar de onde que esta vindo

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(50) not null,
    nickname varchar(50) not null unique, --VARCHAR(50) é um tipo de dado que armazena uma string de até 50 caracteres. NOT NULL significa que esse campo é obrigatório e não pode ser deixado em branco. UNIQUE garante que cada valor nesse campo seja único, ou seja, não pode haver dois usuários com o mesmo nickname.
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    criado_em timestamp default current_timestamp -- timestamp é um tipo de dado que armazena data e hora. DEFAULT CURRENT_TIMESTAMP define o valor padrão para esse campo como a data e hora atual no momento da inserção do registro.
)ENGINE=INNODB; -- InnoDB é um mecanismo de armazenamento para o MySQL que oferece suporte a transações, chaves estrangeiras e integridade referencial. Ele é recomendado para a maioria dos casos de uso devido à sua confiabilidade e desempenho.     
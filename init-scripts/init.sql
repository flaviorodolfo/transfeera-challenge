--CREATE DATABASE transfeera_db ;

CREATE SCHEMA IF NOT EXISTS pagamento;

CREATE TYPE pagamento.tipo_chave_pix_enum AS ENUM ('CPF', 'CNPJ', 'EMAIL', 'TELEFONE', 'CHAVE_ALEATORIA');

CREATE TYPE pagamento.status_recebedor_enum AS ENUM('RASCUNHO','VALIDADO');

CREATE TABLE recebedores (
	recebedor_id SERIAL PRIMARY KEY,
	cpf_cnpj VARCHAR(20) NOT NULL,
	nome VARCHAR(100) NOT NULL,
	tipo_chave_pix pagamento.tipo_chave_pix_enum NOT NULL,
	chave_pix VARCHAR(140) NOT NULL,
	status_recebedor pagamento.status_recebedor_enum NOT NULL,
	email VARCHAR(250) DEFAULT NULL
	
	
)
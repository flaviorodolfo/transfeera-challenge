--CREATE DATABASE transfeera_db ;

CREATE SCHEMA IF NOT EXISTS pagamento;

CREATE TYPE pagamento.tipo_chave_pix_enum AS ENUM ('CPF', 'CNPJ', 'EMAIL', 'TELEFONE', 'CHAVE_ALEATORIA');

CREATE TABLE pagamento.recebedores (
	recebedor_id SERIAL PRIMARY KEY,
	cpf_cnpj VARCHAR(20) NOT NULL,
	nome VARCHAR(100) NOT NULL,
	tipo_chave_pix pagamento.tipo_chave_pix_enum NOT NULL,
	chave_pix VARCHAR(140) NOT NULL,
	status_recebedor VARCHAR(15) DEFAULT 'Rascunho',
	email VARCHAR(250) DEFAULT NULL
	
	
);


INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('783.852.830-56', 'flavio rodolfo', 'CHAVE_ALEATORIA', '0c75c5e2-098b-4843-8cc2-ffa5e291e8b0', 'flaviorodolfo@transfeera.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('908.416.320-65', 'joão dos samtps', 'CHAVE_ALEATORIA', '46892703-d647-4a2c-a6be-a6e0f1488da7', 'joao@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('908.416.320-65', 'joão da silva', 'CPF', '853.464.050-54', 'joao@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('780.015.820-94', 'maria oliveira', 'CPF', '994.405.470-49', 'maria@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('360.657.190-99', 'pedro souza', 'CPF', '228.167.480-06', 'pedro@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('516.785.430-04', 'ana santos', 'CPF', '388.361.480-77', 'ana@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('952.834.790-80', 'lucas pereira', 'CPF', '060.346.360-60', 'lucas@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('230.158.830-03', 'josé silva', 'CPF', '683.284.840-48', 'jose@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('230.158.830-03', 'aline oliveira', 'CPF', '845.108.060-00', 'aline@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('789.654.123-01', 'rafael souza', 'CPF', '78965412301', 'rafael@example.com', 'Validado');


INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('28.802.905/0001-02', 'maria da silva', 'TELEFONE', '11987654321', 'maria@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('28.802.905/0001-02', 'josé oliveira', 'CNPJ', '28.802.905/0001-02', 'jose@example.com', 'Validado');


INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('65.157.117/0001-29', 'ana souza', 'TELEFONE', '33987654321', 'ana@example.com', 'Validado');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('65.157.117/0001-29', 'pedro santos', 'CNPJ', '86.884.624/0001-34', 'pedro@example.com', 'Validado');


INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
VALUES ('01.963.173/0001-78', 'carla oliveira', 'TELEFONE', '55987654321', 'carla@example.com', 'Validado');


INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('01.963.173/0001-78', 'lucas silva', 'CNPJ', '13.127.120/0001-04', 'lucas@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('79.329.147/0001-80', 'fernanda lima', 'TELEFONE', '77987654321', 'fernanda@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('79.329.147/0001-80', 'rafael martins', 'CNPJ', '69.184.715/0001-48', 'rafael@example.com');


INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('10.371.522/0001-53', 'juliana pereira', 'TELEFONE', '99987654321', 'juliana@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('10.371.522/0001-53', 'felipe oliveira', 'CNPJ', '58.341.038/0001-08', 'felipe@example.com');


INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('80.560.231/0001-99', 'camila rodrigues', 'TELEFONE', '12987654321', 'camila@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('80.560.231/0001-99', 'gabriel almeida', 'CNPJ', '73.022.923/0001-18', 'gabriel@example.com');


INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('83.288.301/0001-90', 'mariana costa', 'TELEFONE', '14987654321', 'mariana@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('83.288.301/0001-90', 'carlos santos', 'CNPJ', '60.498.250/0001-25', 'carlos@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('67.904.100/0001-13', 'amanda oliveira', 'TELEFONE', '16987654321', 'amanda@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('67.904.100/0001-13', 'bruno silva', 'CNPJ', '10.923.181/0001-81', 'bruno@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('39.110.459/0001-83', 'carolina alves', 'TELEFONE', '18987654321', 'carolina@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('39.110.459/0001-83', 'lucas oliveira', 'CNPJ', '44.664.436/0001-50', 'lucas@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('28.937.784/0001-06', 'mariana ferreira', 'TELEFONE', '20987654321', 'mariana@example.com');

INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email)
VALUES ('28.937.784/0001-06', 'gustavo santos', 'CNPJ', '75.032.552/0001-80', 'gustavo@example.com');



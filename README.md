# Transfeera-Challenge

Este projeto foi desenvolvido como parte do desafio proposto pela Transfeera.

## Instruções de Execução

Para executar este projeto localmente certifique-se de ter o Docker e siga os passos abaixo:

1. Clone este repositório para o seu ambiente local.
2. Navegue até o diretório do projeto.
3. Execute o seguinte comando para iniciar os contêineres definidos no Docker Compose:
```
docker compose up
```

## Instruçoes de teste

Para testar este projeto, siga estas instruções:
1. Certifique-se de que o Go está instalado em sua máquina.
2. Navegue até o diretório do projeto.
3. Execute os testes utilizando o comando
```
go test ./....
```

## Utilizando a api

Para utilizar a API Transfeera Challenge, siga estas instruções:

1. Importe a coleção de requisições que está na pasta "collection" deste repositório para o seu cliente de API (por exemplo, Postman, Insomnia, etc.).
2. Certifique-se de que a API esteja sendo executada localmente conforme descrito nas instruções de execução.
2. Agora você pode utilizar os endpoints disponíveis na coleção para interagir com a API

### Endpoints da aplicação
Após a inicialização do docker via docker compose a API estará disponível LOCALMENTE em:
```
 http://localhost:8080/api/v1/recebedores
```
- **GET /api/v1/recebedores/id/:id**: Retorna um recebedor com o ID especificado.
- **GET /api/v1/recebedores/nome/:nome**: Retorna os recebedores com o nome especificado.
- **GET /api/v1/recebedores/status/:status**: Retorna os recebedores com o status especificado.
- **GET /api/v1/recebedores/chave?chave={$chave}&pagina={$pagina}**: Retorna os recebedores com a chave especificada.
- **GET /api/v1/recebedores/tipoChave/:tipoChave**: Retorna os recebedores com o tipo de chave especificado.
- **POST /api/v1/recebedores**: Cria um novo recebedor.
- **PATCH /api/v1/recebedores**: Edita um recebedor existente.
- **PATCH /api/v1/recebedores/:id**: Edita o e-mail de um recebedor com o ID especificado(o email deve ser informado em formato JSON no BODY da requisicão)
- **DELETE /api/v1/recebedores/:id**: Deleta um recebedor com o ID especificado.
- **DELETE /api/v1/recebedores/deletar**: Deleta todos os recebedores (os IDS devem  ser informados no BODY da requisição).


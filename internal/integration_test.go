package internal

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/flaviorodolfo/transfeera-challenge/internal/app"
	"github.com/flaviorodolfo/transfeera-challenge/internal/infra/database"
	httpAdp "github.com/flaviorodolfo/transfeera-challenge/internal/infra/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"go.uber.org/zap"
	"gotest.tools/v3/assert"
)

var db *sql.DB
var router *gin.Engine

func startRouter() {
	logger := zap.NewNop()
	repo := database.NewPostgresRecebedorRepository(db)
	service := app.NewRecebedorService(repo, zap.NewNop())
	router = httpAdp.NewRouter(service, logger)
	gin.SetMode(gin.ReleaseMode)

}
func createTestTables() error {
	_, err := db.Exec(`
        CREATE SCHEMA pagamento;

        CREATE TYPE pagamento.tipo_chave_pix_enum AS ENUM ('CPF', 'CNPJ', 'EMAIL', 'TELEFONE', 'CHAVE_ALEATORIA');

        CREATE TABLE IF NOT EXISTS pagamento.recebedores (
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



		INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix, chave_pix, email, status_recebedor)
		VALUES ('783.852.830-56', 'flavio rodolfo', 'CHAVE_ALEATORIA', '0c75c5e2-098b-4843-8cc2-ffa5e291e8b0', 'flaviorodolfo@transfeera.com', 'Validado');
    `)
	if err != nil {
		return fmt.Errorf("erro criando tabela: %v", err)
	}

	return nil
}

func TestMain(m *testing.M) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("postgres", "16.3-alpine", []string{"POSTGRES_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	port := resource.GetPort("5432/tcp")

	connStr := fmt.Sprintf("postgres://postgres:secret@localhost:%s/postgres?sslmode=disable", port)
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	if err := createTestTables(); err != nil {
		log.Fatalf("Could not create test tables: %s", err)
	}
	startRouter()
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestCriarRecebedor(t *testing.T) {
	t.Run("recebedor válido", func(t *testing.T) {
		jsonData := map[string]interface{}{

			"cpf_cnpj":       "515.762.030-69",
			"nome":           "João da Silva",
			"tipo_chave_pix": "CPF",
			"chave_pix":      "515.762.030-69",
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/recebedores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusCreated, resp.Code)
	})
	t.Run("chave pix duplicada", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"cpf_cnpj":       "515.762.030-69",
			"nome":           "João da Silva",
			"tipo_chave_pix": "CPF",
			"chave_pix":      "515.762.030-69",
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/recebedores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
	t.Run("formato do body incorreto", func(t *testing.T) {
		jsonData := map[string]interface{}{}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/recebedores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
	t.Run("campo invalido", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"cpf_cnpj":       "515.762.030-68", //campo inválido
			"nome":           "João da Silva",
			"tipo_chave_pix": "CPF",
			"chave_pix":      "515.762.030-69",
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/recebedores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
	t.Run("buscar recebedor por id", func(t *testing.T) {

		//body, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/recebedores/id/1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
}

func TestBuscarRecebedorPorId(t *testing.T) {

	t.Run("buscar recebedor  por id sucesso", func(t *testing.T) {

		//body, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/recebedores/id/1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
	t.Run("buscar recebedor  por id falha", func(t *testing.T) {

		//body, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/recebedores/id/22", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}
func TestBuscarRecebedorPorNome(t *testing.T) {

	t.Run("buscar recebedor  por nome", func(t *testing.T) {

		//body, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/recebedores/nome/João da Silva", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
}

func TestBuscarRecebedorPorChave(t *testing.T) {

	t.Run("buscar recebedor por chave válida", func(t *testing.T) {

		//body, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/recebedores/chave?chave=28802905000102&pagina=1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
	t.Run("buscar recebedor por chave inválida", func(t *testing.T) {

		//body, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/recebedores/chave?chave=0000020203&pagina=1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestBuscarRecebedorPorStatus(t *testing.T) {

	t.Run("buscar recebedor por status", func(t *testing.T) {

		//body, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/recebedores/status/Rascunho?pagina=1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

}

func TestBuscarRecebedorPorTipoChave(t *testing.T) {

	t.Run("buscar recebedor por tipo chave válida", func(t *testing.T) {

		//body, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/recebedores/tipoChave/CPF?pagina=1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
	t.Run("buscar recebedor por tipo chave inválida", func(t *testing.T) {

		//body, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/recebedores/tipoChave/CPFSO?pagina=1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

}

func TestEditarRecebedor(t *testing.T) {
	t.Run("editar recebedor válido", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"id":             2,
			"cpf_cnpj":       "515.762.030-69",
			"nome":           "João da Silvveiroa",
			"tipo_chave_pix": "EMAIL",
			"chave_pix":      "nova_chave@pix.com",
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPatch, "/api/v1/recebedores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusCreated, resp.Code)
	})
	t.Run("editar recebedor campo inválido", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"id":             2,
			"cpf_cnpj":       "51576203069",
			"nome":           "João da Silvveiroa",
			"tipo_chave_pix": "CPF", //não é possível alterar o tipo de chave sem alterar a chave
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPatch, "/api/v1/recebedores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("editar recebedor com status Validado", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"id":             1,
			"cpf_cnpj":       "51576203069",
			"nome":           "João da Silvveiroa",
			"tipo_chave_pix": "EMAIL",
			"chave_pix":      "suachave@chave.com",
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPatch, "/api/v1/recebedores", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusConflict, resp.Code)
	})

}

func EditarEmailRecebedor(t *testing.T) {
	t.Run("editar email recebedor status Validado", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"email": "flavio@transfeera.com.br",
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPatch, "/api/v1/recebedores/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("editar email recebedor inexistente", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"email": "flavio@transfeera.com.br",
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPatch, "/api/v1/recebedores/99", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
	t.Run("editar email recebedor Rascunho", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"email": "flavio@transfeera.com.br",
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodPatch, "/api/v1/recebedores/99", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})

}

func TestDeletarRecebedor(t *testing.T) {
	t.Run("deletar recebedor por id", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/recebedores/1", bytes.NewBuffer(nil))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("deletar recebedor inexistente por id", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/recebedores/991", bytes.NewBuffer(nil))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})

}

func TestDeletarRecebedores(t *testing.T) {
	t.Run("teste deletar todos recebedores", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"ids": []uint{4, 5, 6},
		}
		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/recebedores/deletar", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("editar email recebedor inexistente", func(t *testing.T) {
		jsonData := map[string]interface{}{
			"ids": []uint{4, 5, 6},
		}

		body, _ := json.Marshal(jsonData)
		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/recebedores/deletar", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusMultiStatus, resp.Code)
	})

}

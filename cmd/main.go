package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/flaviorodolfo/transfeera-challenge/internal/app"
	"github.com/flaviorodolfo/transfeera-challenge/internal/infra/database"
	"github.com/flaviorodolfo/transfeera-challenge/internal/infra/http"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func construirDatabaseUri() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASS"), os.Getenv("DATABASE_NAME"))
}

func initializeDatabase(logger *zap.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", construirDatabaseUri())
	if err != nil {
		logger.Error("Erro ao abrir a conexão com o banco de dados", zap.Error(err))
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logger.Error("Erro ao pingar o banco de dados", zap.Error(err))
		return nil, err
	}

	return db, nil
}
func inicializarLog() *zap.Logger {
	var logger *zap.Logger
	var err error
	if os.Getenv("ENV") == "DEV" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatalf("Erro ao inicializar o logger: %v", err)
	}
	return logger
}
func run() error {
	logger := inicializarLog()
	db, err := initializeDatabase(logger)
	if err != nil {
		logger.Error("openning db conection", zap.Error(err))
		return err
	}
	userRepo := database.NewPostgresRecebedorRepository(db)
	recebedorService := app.NewRecebedorService(userRepo, logger)
	server := http.NewRouter(recebedorService, logger)
	server.Run(":8080")
	return nil
}

func main() {

	if err := run(); err != nil {
		log.Fatalf("iniciando servidor" + err.Error())
	}
	os.Exit(1)
}

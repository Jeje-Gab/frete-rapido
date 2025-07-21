package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"frete-rapido/internal/quote/delivery/http" // Renomeie para o nome real do seu pacote de handlers!
	"frete-rapido/internal/quote/freterapido"
	"frete-rapido/internal/quote/repository"
	"frete-rapido/internal/quote/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	missingVars := []string{}
	if dbUser == "" {
		missingVars = append(missingVars, "DB_USER")
	}
	if dbPass == "" {
		missingVars = append(missingVars, "DB_PASS")
	}
	if dbHost == "" {
		missingVars = append(missingVars, "DB_HOST")
	}
	if dbPort == "" {
		missingVars = append(missingVars, "DB_PORT")
	}
	if dbName == "" {
		missingVars = append(missingVars, "DB_NAME")
	}

	// variaveis vindas de .env
	frToken := os.Getenv("FR_TOKEN")
	frEndpoint := os.Getenv("FR_ENDPOINT")
	frCNPJ := os.Getenv("FR_CNPJ")
	frPlatformCode := os.Getenv("FR_PLATFORM_CODE")
	frDispatcherZip := os.Getenv("FR_DISPATCHER_ZIP")
	if frToken == "" {
		missingVars = append(missingVars, "FR_TOKEN")
	}
	if frEndpoint == "" {
		missingVars = append(missingVars, "FR_ENDPOINT")
	}
	if frCNPJ == "" {
		missingVars = append(missingVars, "FR_CNPJ")
	}
	if frPlatformCode == "" {
		missingVars = append(missingVars, "FR_PLATFORM_CODE")
	}
	if frDispatcherZip == "" {
		missingVars = append(missingVars, "FR_DISPATCHER_ZIP")
	}

	if len(missingVars) > 0 {
		log.Fatalf("Variáveis de ambiente obrigatórias não definidas: %v", missingVars)
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	// Testa a conexão
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Erro ao pingar o banco: %v", err)
	}

	// Instancia o client com todos os parâmetros
	frClient := freterapido.NewClient(
		frToken,
		frEndpoint,
		frCNPJ,
		frPlatformCode,
		frDispatcherZip,
	)

	e := echo.New()
	api := e.Group("/api")

	repo := repository.NewRepository(db)
	quoteUC := usecase.NewUseCase(repo, frClient)
	http.RegisterRoutes(api, quoteUC)

	// Inicia o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Servidor rodando em http://localhost:%s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

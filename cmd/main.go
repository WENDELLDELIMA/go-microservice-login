package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/WENDELLDELIMA/go-microservice-login/config"
	"github.com/WENDELLDELIMA/go-microservice-login/internal/db"
	"github.com/WENDELLDELIMA/go-microservice-login/internal/handler"
	"github.com/WENDELLDELIMA/go-microservice-login/internal/routes"
	"github.com/WENDELLDELIMA/go-microservice-login/internal/service"

	_ "github.com/jackc/pgx/v5/stdlib" // Importa o driver SQL do pgx
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()

	// Usar o driver do pgx com database/sql
	dbConn, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer dbConn.Close()

	// Inicializar Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Inicializar serviços e handlers
	authService := service.NewAuthService(cfg.JWTSecret)
	queries := db.New(dbConn) // Aqui dbConn implementa DBTX
	authHandler := handler.NewAuthHandler(authService, queries)

	// Registrar rotas
	routes.RegisterRoutes(e, authHandler)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

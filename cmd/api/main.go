package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JohnArllon/Meu-SaaS-Go/internal/plataform/supabase"
	"github.com/JohnArllon/Meu-SaaS-Go/internal/services/technicians"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Carrega as variáveis de ambiente (.env)
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de sistema.")
	}

	// 2. Estabelece conexão com o banco de dados (Supabase)
	conn, err := supabase.Connect()
	if err != nil {
		log.Fatalf("Erro crítico: falha na conexão com o banco: %v", err)
	}
	defer conn.Close(context.Background())

	// 3. Inicializa as camadas de serviço (Dependency Injection)
	repo := technicians.NewRepository(conn)
	handler := technicians.NewHandler(repo)

	// 4. Configura o Servidor HTTP (Gin)
	r := gin.Default()

	// 5. Definição das Rotas
	api := r.Group("/api/v1") // Boa prática: versionamento de API
	{
		techniciansGroup := api.Group("/technicians")
		{
			techniciansGroup.POST("", handler.CreateTechnician)
			techniciansGroup.GET("", handler.ListTechnicians)
		}
	}

	// 6. Inicia o servidor
	fmt.Println("🚀 Servidor rodando em http://localhost:8080/api/v1/technicians")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

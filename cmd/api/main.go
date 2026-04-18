package main

import (
	"context"
	"fmt"
	"log"

	//"github.com/JohnArllon/Meu-SaaS-Go/internal/domain"
	"github.com/JohnArllon/Meu-SaaS-Go/internal/plataform/supabase"
	"github.com/JohnArllon/Meu-SaaS-Go/internal/services/technicians"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//1.Carrega o .env
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Erro ao carregar o arquivo .env")
	//
	//}

	//2. Tenta conectar
	godotenv.Load()

	conn, err := supabase.Connect()
	if err != nil {
		log.Fatal("Falha na conexão: %v", err)

	}
	defer conn.Close(context.Background())

	// Iniciando as camadas
	repo := technicians.NewRepository(conn)
	handler := technicians.NewHandler(repo)

	// Configura o servidor Gin
	r := gin.Default()

	// APLICAÇÃO DAS ROTAS

	// 1. Rota para Criar (POST)
	r.POST("/technicians", handler.CreateTechnician)

	// 2. Rota para Listar (GET)
	r.GET("/technicians", handler.ListTechnicians)

	fmt.Println("🚀 Servidor rodando em http://localhost:8080")
	r.Run(":8080") // O servidor "trava" aqui e fica ouvindo

}

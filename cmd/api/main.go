package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JohnArllon/Meu-SaaS-Go/internal/domain"
	"github.com/JohnArllon/Meu-SaaS-Go/internal/plataform/supabase"
	"github.com/JohnArllon/Meu-SaaS-Go/internal/services/technicians"
	"github.com/joho/godotenv"
)

func main() {
	//1.Carrega o .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")

	}

	//2. Tenta conectar
	conn, err := supabase.Connect()
	if err != nil {
		log.Fatal("Falha na conexão: %v", err)

	}
	defer conn.Close(context.Background())

	fmt.Println("🚀 Conexão com o Supabase estabelecida com sucesso!")

	//3. Teste de Engenharia: Usando o Repository para inserir um dado

	// Instaciamos o repositório passando a conexão
	repo := technicians.NewRepository(conn)

	// Criamos um objeto (struct) do tipo Technician como dados de teste
	novoTecnico := domain.Technician{
		Name:  "John Arllon Teste",
		Email: "john.teste@email.com",
		Phone: "11988888877",
	}

	// Chamamos a função Create do repositório
	err = repo.Create(context.Background(), novoTecnico)
	if err != nil {
		fmt.Println("❌ Erro ao salvar técnico: %v\n", err)
	} else {
		fmt.Println("✅ Técnico de teste salvo com sucesso no Supabase!")
	}

	//4. Teste rápido de leitura (opcional, mantendo o que já tinha)
	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Banco de dados ativo!'").Scan(&greeting)
	if err != nil {
		fmt.Println("Erro ao executar query: %v\n", err)
		return
	}
	fmt.Println("Mensagem do banco: %s\n", greeting)
}

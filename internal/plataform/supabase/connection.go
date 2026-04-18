package supabase

import (
	"context"
	"fmt"
	"os"

	// Para evitar o 'apagamento' sempre deixar explicito!
	pgx "github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) { // O pgx aqui refere-se ao nome do pacote importado acima
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		return nil, fmt.Errorf("DATABASE_URL não configurada")
	}

	// Agora deve reconhecer o pgx.Connect
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar: %w", err)
	}

	return conn, nil
}

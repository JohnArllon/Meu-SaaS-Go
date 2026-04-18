package technicians

import (
	"context"
	"fmt"

	"github.com/JohnArllon/Meu-SaaS-Go/internal/domain"
	"github.com/jackc/pgx/v5"
)

// Repository lida com a comunicação direta com o banco
type Repository struct {
	conn *pgx.Conn
}

// NewRepository cria uma nova instancia do repositorio
func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{conn}
}

// Create insere um novo técnico no Supabase
func (r *Repository) Create(ctx context.Context, t domain.Technician) error {
	query := `INSERT INTO technicians (name, email, phone) VALUES ($1, $2, $3)`

	_, err := r.conn.Exec(ctx, query, t.Name, t.Email, t.Phone)
	if err != nil {
		return fmt.Errorf("erro ao inserir técnico: %w", err)
	}
	return nil
}

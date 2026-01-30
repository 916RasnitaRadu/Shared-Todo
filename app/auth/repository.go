package auth

import (
	"context"
	"shared-todo/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository interface {
	GetUser(ctx context.Context, username string) (db.User, error)
	//GetUserByEmail(ctx context.Context, email string) (db.User, error)
}

// DB Repository

type DBRepository struct {
	queries *db.Queries
}

func NewDBRepository(q *db.Queries) Repository {
	return &DBRepository{q}
}

func (repo *DBRepository) GetUser(ctx context.Context, username string) (db.User, error) {
	return repo.queries.GetUser(ctx, pgtype.Text{String: username, Valid: true})
}

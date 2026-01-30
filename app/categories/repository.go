package categories

import (
	"context"
	"fmt"
	"shared-todo/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository interface {
	GetCategories(ctx context.Context, userid int32) ([]db.CategoryDTO, error)
	GetUserID(ctx context.Context, username string) (int32, error)
	AddCategory(ctx context.Context, user_id int32, name string, priority int32) (int32, error)
	DeleteCategory(ctx context.Context, id int32) error
	GetCategoryByNameAndId(ctx context.Context, userid int32, name string) (db.Category, error)
}

type DBRepository struct {
	queries *db.Queries
}

func NewDBRepository(q *db.Queries) Repository {
	return &DBRepository{q}
}

func (r *DBRepository) GetCategories(ctx context.Context, userid int32) ([]db.CategoryDTO, error) {
	categoriesDB, err := r.queries.GetCategories(ctx, pgtype.Int4{Int32: userid, Valid: true})

	if err != nil {
		return nil, fmt.Errorf("categories error: %w", err)
	}

	categoriesDTO := make([]db.CategoryDTO, len(categoriesDB))

	for i, categ := range categoriesDB {
		categoriesDTO[i] = mapperCategoryDTO(categ)
	}
	return categoriesDTO, nil
}

func (r *DBRepository) GetUserID(ctx context.Context, username string) (int32, error) {
	return r.queries.GetUserId(ctx, pgtype.Text{String: username, Valid: true})
}

func (r *DBRepository) AddCategory(ctx context.Context, user_id int32, name string, priority int32) (int32, error) {
	return r.queries.CreateCategory(ctx, db.CreateCategoryParams{
		UserID:   pgtype.Int4{Int32: user_id, Valid: true},
		Name:     pgtype.Text{String: name, Valid: true},
		Priority: pgtype.Int4{Int32: priority, Valid: true},
	})
}

func (r *DBRepository) DeleteCategory(ctx context.Context, id int32) error {
	return r.queries.DeleteCategory(ctx, id)
}

func (r *DBRepository) GetCategoryByNameAndId(ctx context.Context, userid int32, name string) (db.Category, error) {
	return r.queries.GetCategoryByNameAndId(ctx, db.GetCategoryByNameAndIdParams{
		UserID: pgtype.Int4{Int32: userid, Valid: true},
		Name:   pgtype.Text{String: name, Valid: true},
	})
}

package items

import (
	"context"
	"fmt"
	"shared-todo/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository interface {
	GetItems(ctx context.Context, category_id int32) ([]db.ItemDTO, error)
	AddItem(ctx context.Context, category_id int32, name string, description string) (db.ItemDTO, error)
	GetCategoryId(ctx context.Context, item_id int32) (int32, error)
	DeleteItem(ctx context.Context, id int32) error
	UpdateItemDoneStatus(ctx context.Context, id int32, done bool) (db.ItemDTO, error)
}

type DBRepository struct {
	queries *db.Queries
}

func NewDBRepository(q *db.Queries) Repository {
	return &DBRepository{q}
}

func (r *DBRepository) GetItems(ctx context.Context, category_id int32) ([]db.ItemDTO, error) {
	itemsDB, err := r.queries.GetItems(ctx, pgtype.Int4{Int32: category_id, Valid: true})

	if err != nil {
		return nil, fmt.Errorf("categories error: %w", err)
	}

	itemsDTO := make([]db.ItemDTO, len(itemsDB))

	for i, item := range itemsDB {
		itemsDTO[i] = mapperItemDTO(item)
	}
	return itemsDTO, nil
}

func (r *DBRepository) GetCategoryId(ctx context.Context, item_id int32) (int32, error) {
	category_id, err := r.queries.GetCategoryByItemId(ctx, item_id)

	return category_id.Int32, err
}

func (r *DBRepository) AddItem(ctx context.Context, category_id int32, name string, description string) (db.ItemDTO, error) {
	item, err := r.queries.CreateItem(ctx, db.CreateItemParams{
		CategoryID:  pgtype.Int4{Int32: category_id, Valid: true},
		Name:        pgtype.Text{String: name, Valid: true},
		Description: pgtype.Text{String: description, Valid: true},
	})

	var itemDTO = mapperItemDTO(item)
	return itemDTO, err
}

func (r *DBRepository) DeleteItem(ctx context.Context, id int32) error {
	return r.queries.DeleteItem(ctx, id)
}

func (r *DBRepository) UpdateItemDoneStatus(ctx context.Context, id int32, done bool) (db.ItemDTO, error) {
	var item, err = r.queries.UpdateItemDoneStatus(ctx, db.UpdateItemDoneStatusParams{
		ID:   id,
		Done: pgtype.Bool{Bool: done, Valid: true},
	})

	var itemDTO = mapperItemDTO(item)
	fmt.Println(itemDTO)
	return itemDTO, err
}

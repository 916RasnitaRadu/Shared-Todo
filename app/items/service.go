package items

import (
	"context"
	"fmt"
	"shared-todo/db"
)

type Service struct {
	ctx        context.Context
	repository Repository
}

func NewService(ctx context.Context, repo Repository) *Service {
	return &Service{ctx: ctx, repository: repo}
}

func (srv *Service) GetItemsService(ctx context.Context, category_id int32) ([]db.ItemDTO, error) {
	items, err := srv.repository.GetItems(ctx, category_id)
	if err != nil {
		fmt.Println("Something went wrong..")
		return []db.ItemDTO{}, err
	}

	return items, err
}

func (srv *Service) GetCategoryIdService(ctx context.Context, item_id int32) (int32, error) {
	categ_id, err := srv.repository.GetCategoryId(ctx, item_id)

	if err != nil {
		fmt.Println("Something went wrong..")
		return -1, err
	}
	return categ_id, err
}

func (srv *Service) AddItemService(ctx context.Context, category_id int32, name string, description string) (db.ItemDTO, error) {
	item, err := srv.repository.AddItem(ctx, category_id, name, description)
	if err != nil {
		fmt.Println("Something went wrong..")
		return db.ItemDTO{}, err
	}
	return item, err
}

func (srv *Service) DeleteItemService(ctx context.Context, id int32) error {
	err := srv.repository.DeleteItem(ctx, id)
	if err != nil {
		fmt.Println("Something went wrong..")
	}
	return err
}

func (srv *Service) UpdateItemDoneStatusService(ctx context.Context, id int32, done bool) (db.ItemDTO, error) {
	item, err := srv.repository.UpdateItemDoneStatus(ctx, id, done)

	if err != nil {
		fmt.Println("Something went wrong..")
		return db.ItemDTO{}, err
	}
	return item, err
}

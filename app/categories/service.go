package categories

import (
	"context"
	"errors"
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

func (srv *Service) GetCategoriesService(ctx context.Context, username string) []db.CategoryDTO {
	var user_id, err = srv.repository.GetUserID(ctx, username)

	if err != nil {
		fmt.Println("Something went wrong..")
		return nil
	}

	categories, err := srv.repository.GetCategories(ctx, user_id)
	if err != nil {
		fmt.Println("Something went wrong..")
		return nil
	}

	return categories

}

var ErrCategoryExists = errors.New("category already exists for this user")

func (srv *Service) AddCategoryService(ctx context.Context, username string, name string, priority int32) (db.CategoryDTO, error) {
	var user_id, err = srv.repository.GetUserID(ctx, username)
	if err != nil {
		fmt.Println("Something went wrong..")
		return db.CategoryDTO{}, err
	}

	// check if category exists
	_, err = srv.repository.GetCategoryByNameAndId(ctx, user_id, name)
	if err == nil { // if the category exists for this user we throw an error
		fmt.Println("Category already exists for this user")
		return db.CategoryDTO{}, ErrCategoryExists
	}

	id, err := srv.repository.AddCategory(ctx, user_id, name, priority)
	if err != nil {
		fmt.Println("Something went wrong..")
		return db.CategoryDTO{}, err
	}

	var newCategory = db.CategoryDTO{
		ID:       id,
		UserID:   user_id,
		Name:     name,
		Priority: priority,
	}
	return newCategory, err
}

func (srv *Service) DeleteCategoryService(ctx context.Context, id int32) error {
	err := srv.repository.DeleteCategory(ctx, id)
	if err != nil {
		fmt.Println("Something went wrong..")
	}

	return err
}

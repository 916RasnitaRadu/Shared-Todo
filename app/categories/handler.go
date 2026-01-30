package categories

import (
	"context"
	"errors"
	"net/http"
	"shared-todo/view/dashboard"
	"shared-todo/view/start"
	"strconv"

	"os"
	"shared-todo/db"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func GetService(ctx context.Context) *Service {
	conn, _ := pgx.Connect(ctx, os.Getenv("DB_CONNECTION_STRING"))
	queries := db.New(conn)
	repo := NewDBRepository(queries)
	return NewService(ctx, repo)
}

func HandleCategoriesPage(w http.ResponseWriter, r *http.Request) {
	categories := HandleGetCategories(w, r)
	start.HtmxPage(dashboard.Dashboard(categories)).Render(r.Context(), w)
}

func HandleGetCategories(w http.ResponseWriter, r *http.Request) []db.CategoryDTO {
	service := GetService(r.Context())

	username := extractUsername(w, r)
	return service.GetCategoriesService(r.Context(), username)
}

func HandlePostCategory(w http.ResponseWriter, r *http.Request) {
	var er error = r.ParseForm()
	var categoryDTO db.CategoryDTO
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	service := GetService(r.Context())

	username := extractUsername(w, r)
	name := r.FormValue("name")
	priorityString := r.FormValue("priority")

	priority64, _ := strconv.ParseInt(priorityString, 10, 32)

	var priority int32 = int32(priority64)
	categoryDTO, er = service.AddCategoryService(r.Context(), username, name, priority)

	if er != nil {
		if errors.Is(er, ErrCategoryExists) {
			w.Header().Set("hx-trigger", "category-exists")
		} else {
			http.Error(w, "Failed to add category", http.StatusInternalServerError)
		}
		return
	}

	dashboard.CategoryItem(categoryDTO).Render(r.Context(), w)

}

func HandleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	service := GetService(r.Context())

	err = service.DeleteCategoryService(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	// re-render the list
	username := extractUsername(w, r)
	categories := service.GetCategoriesService(r.Context(), username)

	dashboard.CategoryList(categories).Render(r.Context(), w)
}

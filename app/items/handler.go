package items

import (
	"context"
	"net/http"
	"os"
	"shared-todo/db"
	"shared-todo/view/item_dashboard"
	"shared-todo/view/start"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func GetService(ctx context.Context) *Service {
	conn, _ := pgx.Connect(ctx, os.Getenv("DB_CONNECTION_STRING"))
	queries := db.New(conn)
	repo := NewDBRepository(queries)
	return NewService(ctx, repo)
}

func HandlerGetItems(w http.ResponseWriter, r *http.Request) {
	service := GetService(r.Context())
	idParam := chi.URLParam(r, "category_id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	items, err := service.GetItemsService(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Problem at fetching items", http.StatusInternalServerError)
		return
	}

	start.HtmxPage(item_dashboard.ItemDashboard(items, int32(id))).Render(r.Context(), w)
}

func HandlePostItem(w http.ResponseWriter, r *http.Request) {
	var category_id int32
	var name string
	var description string

	service := GetService(r.Context())
	idParam := chi.URLParam(r, "category_id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	category_id = int32(id)

	name = r.FormValue("name")
	description = r.FormValue("description")

	itemDTO, err := service.AddItemService(r.Context(), category_id, name, description)
	if err != nil {
		http.Error(w, "Failed to add item", http.StatusInternalServerError)
		return
	}
	item_dashboard.Item(itemDTO).Render(r.Context(), w)
}

func HandleDeleteItem(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "item_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}
	service := GetService(r.Context())

	// use category id so that we can bring the item list after deleting
	category_id, err := service.GetCategoryIdService(r.Context(), int32(id))
	if err != nil {
		return
	}

	err = service.DeleteItemService(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return
	}

	item_list, err := service.GetItemsService(r.Context(), category_id)
	if err != nil {
		http.Error(w, "Problem at fetching items", http.StatusInternalServerError)
		return
	}

	item_dashboard.ItemList(item_list).Render(r.Context(), w)
}

func HandleUpdateItemDoneStatus(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "item_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}
	doneParam := r.FormValue("done")

	var done bool
	if doneParam == "on" {
		done = true
	} else {
		done = false
	}

	service := GetService(r.Context())
	item, err := service.UpdateItemDoneStatusService(r.Context(), int32(id), done)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	item_dashboard.Item(item).Render(r.Context(), w)

}

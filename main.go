package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"shared-todo/app"
	"shared-todo/app/auth"
	"shared-todo/app/categories"
	"shared-todo/app/items"
	"shared-todo/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5"
)

const (
	SERVER_ADDR = ":8080"
)

func main() {
	// sample connection to db
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		fmt.Printf("Database connection error %v\n", err)
		return
	}

	queries := db.New(conn)
	r, err := queries.SampleQuery(context.Background())
	if err != nil {
		fmt.Printf("Database query error %v\n", err)
	}

	fmt.Printf("Got from db: %v\n", r)

	// sample routing
	auth.Init()

	router := chi.NewRouter()

	router.Use(jwtauth.Verifier(auth.TokenAuth))

	router.Get("/hello", app.HandleHelloWorld)
	router.Get("/hello/{name}", app.HandleHelloName)

	router.Route("/login", func(r chi.Router) {
		//r.Use(redirectToDashboard)

		r.Get("/", auth.HandleLoginPage)
		r.Post("/", auth.HandleLogin)
	})
	router.Route("/categories", func(r chi.Router) {
		r.Use(redirectToLogin)

		r.Get("/", categories.HandleCategoriesPage)
		r.Get("/{category_id}/items", items.HandlerGetItems)
		r.Post("/{category_id}/items", items.HandlePostItem)
		r.Post("/", categories.HandlePostCategory)
		r.Delete("/{id}", categories.HandleDeleteCategory)
	})
	router.Route("/items", func(r chi.Router) {
		r.Use(redirectToLogin)

		r.Delete("/{item_id}", items.HandleDeleteItem)
		r.Put("/{item_id}", items.HandleUpdateItemDoneStatus)
	})
	router.Post("/logout", auth.HandleLogout)

	fmt.Printf("Server started at http://localhost%s\n", SERVER_ADDR)
	log.Fatal(http.ListenAndServe(SERVER_ADDR, router))
}

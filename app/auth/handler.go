package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"shared-todo/db"
	"shared-todo/view/login"
	"shared-todo/view/start"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetService(ctx context.Context) *Service {
	conn, _ := pgx.Connect(ctx, os.Getenv("DB_CONNECTION_STRING"))
	queries := db.New(conn)
	repo := NewDBRepository(queries)
	return NewService(ctx, repo)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var er error = r.ParseForm()
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	service := GetService(r.Context())

	username := r.FormValue("username")
	password := r.FormValue("password")

	_, tokenString, er := service.Login(username, password)

	if er != nil {
		log.Println(er)
		w.Header().Set("hx-trigger", "wrong-credentials")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 2),
		Path:     "/",
		HttpOnly: true,
	})

	w.Header().Set("HX-Redirect", "/categories")
}

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	start.HtmxPage(login.LoginPage()).Render(r.Context(), w)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Set to a past time to delete the cookie
		Path:     "/",
		HttpOnly: true,
	})

	// Optionally, redirect the user to the login page or home page after logout
	w.Header().Set("HX-Redirect", "/login")
}

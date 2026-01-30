package auth

import (
	"context"
	"log"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	ctx        context.Context
	repository Repository
}

func NewService(ctx context.Context, repo Repository) *Service {
	return &Service{ctx, repo}
}

func (srv *Service) Login(username string, password string) (token jwt.Token, tokenString string, err error) {
	user, err := srv.repository.GetUser(srv.ctx, username)

	if err != nil {
		log.Println("The user is not in the db")
		return nil, tokenString, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(password))
	if err != nil {
		log.Println("The passwords do not match")
		return nil, tokenString, err
	}

	claims := map[string]interface{}{
		"username": username,
	}

	jwtauth.SetExpiryIn(claims, 2*time.Hour)
	return TokenAuth.Encode(claims)
}

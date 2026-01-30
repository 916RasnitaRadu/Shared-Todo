package auth

import (
	"log"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

func Init() {
	secretKey := os.Getenv("KEY")
	if secretKey == "" {
		log.Fatal("Key is missing!")
		return
	}
	TokenAuth = jwtauth.New("HS256", []byte(secretKey), nil)

}

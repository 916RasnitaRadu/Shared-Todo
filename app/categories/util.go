package categories

import (
	"net/http"
	"shared-todo/db"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func extractUsername(w http.ResponseWriter, r *http.Request) string {
	// Extract the token from the request context
	token, claims, err := jwtauth.FromContext(r.Context())
	if err != nil || token == nil || jwt.Validate(token) != nil {
		http.Error(w, "Invalid or missing token", http.StatusUnauthorized)
		return ""
	}

	// Extract the claims from the token
	username := claims["username"].(string)
	return username
}

func mapperCategoryDTO(category db.Category) db.CategoryDTO {
	var categoryDTO = db.CategoryDTO{
		ID:       category.ID,
		UserID:   category.UserID.Int32,
		Name:     category.Name.String,
		Priority: category.Priority.Int32,
	}

	return categoryDTO
}

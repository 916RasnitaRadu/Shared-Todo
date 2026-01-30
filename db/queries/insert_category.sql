-- name: CreateCategory :one
INSERT INTO public.categories (user_id, name, priority) VALUES ($1, $2, $3)
RETURNING id;
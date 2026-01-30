-- name: CreateItem :one
INSERT INTO public.items (category_id, name, description, done, created_at) VALUES ($1, $2, $3, false, NOW())
RETURNING id, category_id, name, description, done, created_at;
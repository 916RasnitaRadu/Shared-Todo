-- name: GetCategories :many
SELECT * FROM public.categories
WHERE user_id = $1 ORDER BY priority;

-- name: GetCategoryByNameAndId :one
SELECT * FROM public.categories
WHERE user_id = $1 AND name = $2;

-- name: GetCategoryByItemId :one
SELECT category_id FROM public.items
WHERE id = $1;
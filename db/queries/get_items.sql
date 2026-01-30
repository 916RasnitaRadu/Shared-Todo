-- name: GetItems :many
SELECT * FROM public.items WHERE category_id=$1;
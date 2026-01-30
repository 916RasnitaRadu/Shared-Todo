-- name: UpdateItemDoneStatus :one
UPDATE public.items SET done = $2
WHERE id = $1
RETURNING id, category_id, name, description, done, created_at;
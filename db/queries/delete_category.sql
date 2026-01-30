-- name: DeleteCategory :exec
DELETE FROM public.categories WHERE id=$1;
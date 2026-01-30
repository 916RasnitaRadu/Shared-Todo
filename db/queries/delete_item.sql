-- name: DeleteItem :exec
DELETE FROM public.items WHERE id=$1;
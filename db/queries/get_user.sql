-- name: GetUser :one
SELECT * FROM public.users
WHERE username = $1;

-- name: GetUserId :one
SELECT ID FROM public.users
WHERE username = $1;
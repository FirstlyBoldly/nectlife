-- name: CreateRole :one
INSERT INTO "role" (
  "key_name",
  "description"
) VALUES (
  $1, $2
) RETURNING *;

-- name: FlushRoles :exec
DELETE FROM "role" CASCADE;

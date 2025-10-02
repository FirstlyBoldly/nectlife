-- name: CreateSession :one
INSERT INTO "session" (
  "token", "data", "expires_at"
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ListSessions :many
SELECT * FROM "session"
ORDER BY "expires_at";

-- name: GetSession :one
SELECT * FROM "session"
WHERE "token" = $1
LIMIT 1;

-- name: UpdateSession :one
UPDATE "session"
SET "token" = $2,
"data" = $3,
"expires_at" = $4
WHERE "id" = $1
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM "session"
WHERE "token" = $1;

-- name: GarbageCollect :exec
DELETE FROM "session"
WHERE "expires_at" >= NOW();

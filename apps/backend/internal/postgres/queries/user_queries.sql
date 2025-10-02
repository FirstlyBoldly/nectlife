-- name: GetUser :one
SELECT * FROM "user"
WHERE "id" = $1 LIMIT 1;

-- name: GetUserByStudentId :one
SELECT * FROM "user"
WHERE "student_id" = $1 LIMIT 1;

-- name: GetUserIdByStudentId :one
SELECT "id" FROM "user"
WHERE "student_id" = $1 LIMIT 1;

-- name: GetActiveUser :one
SELECT * FROM "active_user"
WHERE "id" = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY "last_name", "first_name";

-- name: ListActiveUsers :many
SELECT * FROM "active_user"
ORDER BY "last_name", "first_name";

-- name: CreateUser :one
INSERT INTO "user" (
  "course_id",
  "role_id",
  "student_id",
  "first_name",
  "last_name",
  "email",
  "password_hash"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: UpdateUser :exec
UPDATE "user"
SET "course_id" = $2,
"role_id" = $3,
"student_id" = $4,
"first_name" = $5,
"middle_name" = $6,
"last_name" = $7,
"display_name" = $8,
"headline" = $9,
"description" = $10,
"email" = $11,
"password_hash" = $12,
"status" = $13
WHERE "id" = $1;

-- name: SoftDeleteUser :exec
UPDATE "user"
SET "deleted_at" = NOW()
WHERE "id" = $1;

-- name: PermaDeleteUser :exec
DELETE FROM "user"
WHERE "id" = $1;

-- name: GetPasswordHashByStudentId :one
SELECT "password_hash"
FROM "user"
WHERE "student_id" = $1
LIMIT 1;

-- name: UserExists :one
SELECT EXISTS (
  SELECT 1
  FROM "user"
  WHERE "id" = $1
);

-- name: FlushUsers :exec
DELETE FROM "user" CASCADE;

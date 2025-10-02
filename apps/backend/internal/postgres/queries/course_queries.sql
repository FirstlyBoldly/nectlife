-- name: CreateCourse :one
INSERT INTO "course" (
  "department_id",
  "key_name"
) VALUES (
  $1, $2
) RETURNING *;

-- name: FlushCourses :exec
DELETE FROM "course" CASCADE;

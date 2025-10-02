-- name: CreateInvitation :one
INSERT INTO "invitation" (
  "invited_by_user_id",
  "student_id",
  "email",
  "token",
  "role_id",
  "expires_at"
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: AcceptInvitationByToken :exec
UPDATE "invitation"
SET "status" = 'accepted'
WHERE "token" = $1;

-- name: CheckForExpiredInvitations :exec
UPDATE "invitation"
SET "status" = 'expired'
WHERE NOW() >= "expires_at"
AND "status" = 'pending';

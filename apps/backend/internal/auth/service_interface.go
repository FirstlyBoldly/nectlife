package auth

import (
	"context"

	"github.com/nectgrams-webapp-team/nectlife/internal/session"
)

type AuthServiceInterface interface {
	Register(
		ctx context.Context,
		first_name string,
		last_name string,
		role_id string,
		student_id string,
		course_id string,
		email string,
		password string,
	) error
	VerifyCredentials(
		ctx context.Context,
		email string,
		password string,
	) error
	Login(
		ctx context.Context,
		sh *session.SessionHandler,
		user_id int32,
	) error
	Logout(
		ctx context.Context,
		sh *session.SessionHandler,
	) error
}

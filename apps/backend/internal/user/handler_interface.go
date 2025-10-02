package user

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
	"github.com/nectgrams-webapp-team/nectlife/internal/auth"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
)

type CreateUserInput struct {
	Body postgres.CreateUserParams
}

type CreateUserOutput struct {
	Body postgres.User
}

type UserIDInput struct {
	ID int32 `path:"id" doc:"User ID"`
}

type ListUsersOutput struct {
	Body []postgres.User
}

type GetUserOutput struct {
	Body postgres.User
}

type UpdateUserInput struct {
	ID   int32 `path:"id" doc:"User ID to update"`
	Body postgres.UpdateUserParams
}

type UserHandlerInterface interface {
	CreateUser(context.Context, *CreateUserInput) (*CreateUserOutput, error)
	GetUserById(context.Context, *UserIDInput) (*GetUserOutput, error)
	ListUsers(context.Context, *struct{}) (*ListUsersOutput, error)
	UpdateUser(context.Context, *UpdateUserInput) (*struct{}, error)
	DeleteUser(context.Context, *UserIDInput) (*struct{}, error)
	Routes(*chi.Mux, huma.API, auth.AuthHandlerInterface)
}

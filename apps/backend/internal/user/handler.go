package user

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/nectgrams-webapp-team/nectlife/internal/auth"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
)

type UserHandler struct {
	db *postgres.Queries
}

func NewUserHandler(db *postgres.Queries) UserHandlerInterface {
	return &UserHandler{db: db}
}

func (h *UserHandler) CreateUser(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	user, err := h.db.CreateUser(ctx, input.Body)
	if err != nil {
		text := "failed to create new user"
		slog.Error(err.Error())
		return nil, huma.Error500InternalServerError(text, err)
	}

	return &CreateUserOutput{Body: user}, nil
}

func (h *UserHandler) ListUsers(ctx context.Context, _ *struct{}) (*ListUsersOutput, error) {
	users, err := h.db.ListUsers(ctx)
	if err != nil {
		text := "failed to get users"
		slog.Error(err.Error())
		return nil, huma.Error500InternalServerError(text, err)
	}

	return &ListUsersOutput{Body: users}, nil
}

func (h *UserHandler) GetUserById(ctx context.Context, input *UserIDInput) (*GetUserOutput, error) {
	user, err := h.db.GetUser(ctx, input.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			text := fmt.Sprintf("no user with ID %d exists", input.ID)
			slog.Error(err.Error())
			return nil, huma.Error404NotFound(text, err)
		}

		text := fmt.Sprintf("failed to get the user with ID %d", input.ID)
		slog.Error(err.Error())
		return nil, huma.Error500InternalServerError(text, err)
	}

	return &GetUserOutput{Body: user}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, input *UpdateUserInput) (*struct{}, error) {
	if err := h.db.UpdateUser(ctx, input.Body); err != nil {
		text := fmt.Sprintf("failed to update user with ID %d", input.ID)
		slog.Error(err.Error())
		return nil, huma.Error500InternalServerError(text, err)
	}

	return nil, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, input *UserIDInput) (*struct{}, error) {
	if err := h.db.PermaDeleteUser(ctx, input.ID); err != nil {
		text := fmt.Sprintf("failed to delete user with ID %d", input.ID)
		slog.Error(err.Error())
		return nil, huma.Error500InternalServerError(text, err)
	}

	return nil, nil
}

func (h *UserHandler) Routes(mux *chi.Mux, api huma.API, authHndlr auth.AuthHandlerInterface) {
	huma.Register(api, huma.Operation{
		OperationID: "get-user-by-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/users/{id}",
		Summary:     "Get a user instance by its id",
		Tags:        []string{"Users"},
	}, h.GetUserById)

	huma.Register(api, huma.Operation{
		OperationID: "list-users",
		Method:      http.MethodGet,
		Path:        "/api/v1/users",
		Summary:     "List all existing user instances",
		Tags:        []string{"Users"},
	}, h.ListUsers)

	huma.Register(api, huma.Operation{
		OperationID:   "create-user",
		Method:        http.MethodPost,
		Path:          "/api/v1/users",
		Summary:       "Create a new user instance",
		Tags:          []string{"Users"},
		DefaultStatus: http.StatusCreated,
	}, h.CreateUser)

	huma.Register(api, huma.Operation{
		OperationID:   "update-user",
		Method:        http.MethodPatch,
		Path:          "/api/v1/users/{id}",
		Summary:       "Update a user instance",
		Tags:          []string{"Users"},
		DefaultStatus: http.StatusNoContent,
	}, h.UpdateUser)

	huma.Register(api, huma.Operation{
		OperationID:   "delete-user",
		Method:        http.MethodDelete,
		Path:          "/api/v1/users/{id}",
		Summary:       "Delete a user instance",
		Tags:          []string{"Users"},
		DefaultStatus: http.StatusNoContent,
	}, h.DeleteUser)
}

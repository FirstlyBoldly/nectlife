package auth

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
	"github.com/nectgrams-webapp-team/nectlife/internal/session"
)

type AuthHandler struct {
	svc       *AuthService
	sessHndlr *session.SessionHandler
}

func (h *AuthHandler) Status(ctx context.Context, _ *struct{}) (*StatusOutput, error) {
	sess, ok := session.GetRequestSession(ctx)
	if !ok {
		slog.Error("failed to get session")
		return nil, nil
	}

	userId, ok := sess.Data["user_id"]
	if !ok {
		return nil, nil
	}

	userIdInt, err := strconv.ParseInt(*userId, 10, 32)
	if err != nil {
		slog.Error(err.Error())
		return nil, nil
	}

	return &StatusOutput{
		Body: struct {
			Authenticated bool
			UserID        int32
		}{
			ok,
			int32(userIdInt),
		},
	}, nil
}

func (h *AuthHandler) Register(ctx context.Context, input *RegisterInput) (*struct{}, error) {
	if err := h.svc.Register(
		ctx,
		input.Body.FirstName,
		input.Body.LastName,
		input.Body.RoleID,
		input.Body.StudentID,
		input.Body.CourseID,
		input.Body.Email,
		input.Body.Password,
	); err != nil {
		text := "failed to register new user"
		slog.Error(err.Error())
		huma.Error500InternalServerError(text, err)
	}

	return nil, nil
}

func (h *AuthHandler) Login(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	var id int32
	const wait = 1 * time.Second
	start := time.Now()

	err := h.svc.VerifyCredentials(ctx, input.Body.StudentID, input.Body.Password)
	if err == nil {
		id, err = h.svc.db.GetUserIdByStudentId(ctx, input.Body.StudentID)
		if err != nil {
			text := "failed to get user ID"
			slog.Error(err.Error())
			return nil, huma.Error500InternalServerError(text, err)
		}

		if err = h.svc.Login(ctx, h.sessHndlr, id); err != nil {
			text := "failed to login"
			slog.Error(err.Error())
			return nil, huma.Error500InternalServerError(text, err)
		}
	}

	if time.Since(start) < wait {
		time.Sleep(wait - time.Since(start))
	}

	if err != nil {
		text := "invalid credentials"
		slog.Error(err.Error())
		return nil, huma.Error401Unauthorized(text, err)
	}

	return &LoginOutput{Body: struct{ ID int32 }{ID: id}}, nil
}

func (h *AuthHandler) Logout(ctx context.Context, _ *struct{}) (*struct{}, error) {
	if err := h.svc.Logout(ctx, h.sessHndlr); err != nil {
		text := "failed to logout"
		slog.Error(err.Error())
		return nil, huma.Error500InternalServerError(text, err)
	}

	return nil, nil
}

func (h *AuthHandler) Routes(mux *chi.Mux, api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "status",
		Method:        http.MethodGet,
		Path:          "/api/v1/auth/status",
		Summary:       "Session status",
		Description:   "Returns a boolean indicating if a session is authenticated and the session data",
		Tags:          []string{"Auth"},
		DefaultStatus: http.StatusOK,
	}, h.Status)

	huma.Register(api, huma.Operation{
		OperationID:   "login",
		Method:        http.MethodPost,
		Path:          "/api/v1/auth/login",
		Summary:       "Login user",
		Tags:          []string{"Auth"},
		DefaultStatus: http.StatusOK,
	}, h.Login)

	huma.Register(api, huma.Operation{
		OperationID:   "register",
		Method:        http.MethodPost,
		Path:          "/api/v1/auth/register",
		Summary:       "Register a new user",
		Tags:          []string{"Auth"},
		DefaultStatus: http.StatusCreated,
	}, h.Register)

	huma.Register(api, huma.Operation{
		OperationID:   "logout",
		Method:        http.MethodGet,
		Path:          "/api/v1/auth/logout",
		Summary:       "Logout user",
		Tags:          []string{"Auth"},
		DefaultStatus: http.StatusNoContent,
	}, h.Logout)
}

func NewAuthHandler(db *postgres.Queries, sessSvc session.SessionServiceInterface, sessHndlr session.SessionHandlerInterface) AuthHandlerInterface {
	return &AuthHandler{
		svc:       NewAuthService(db, sessSvc).(*AuthService),
		sessHndlr: sessHndlr.(*session.SessionHandler),
	}
}

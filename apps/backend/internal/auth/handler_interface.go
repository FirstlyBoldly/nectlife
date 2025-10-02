package auth

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
)

type StatusOutput struct {
	Body struct {
		Authenticated bool
		UserID        int32
	}
}

type RegisterInput struct {
	Body struct {
		FirstName string `json:"first_name" doc:"User first name"`
		LastName  string `json:"last_name" doc:"User last name"`
		RoleID    string `json:"role_id" doc:"User role ID"`
		StudentID string `json:"student_id" doc:"User student ID"`
		CourseID  string `json:"course_id" doc:"User course ID"`
		Email     string `json:"email" doc:"User email address"`
		Password  string `json:"password" doc:"User password"`
	}
}

type LoginInput struct {
	Body struct {
		StudentID string `json:"student_id" doc:"User student ID"`
		Password  string `json:"password" doc:"User password"`
	}
}

type LoginOutput struct {
	Body struct {
		ID int32
	}
}

type AuthHandlerInterface interface {
	Status(context.Context, *struct{}) (*StatusOutput, error)
	Register(context.Context, *RegisterInput) (*struct{}, error)
	Login(context.Context, *LoginInput) (*LoginOutput, error)
	Logout(context.Context, *struct{}) (*struct{}, error)
	Routes(*chi.Mux, huma.API)
}

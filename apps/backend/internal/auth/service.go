package auth

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
	"github.com/nectgrams-webapp-team/nectlife/internal/session"
	"golang.org/x/crypto/bcrypt"
)

const userIdKey = "user_id"

type AuthService struct {
	db    *postgres.Queries
	store session.SessionServiceInterface
}

func (s *AuthService) Register(
	ctx context.Context,
	first_name string,
	last_name string,
	role_id string,
	student_id string,
	course_id string,
	email string,
	password string,
) error {
	course_id_int, err := strconv.Atoi(course_id)
	if err != nil {
		return err
	}

	role_id_int, err := strconv.Atoi(role_id)
	if err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	_, err = s.db.CreateUser(
		ctx,
		postgres.CreateUserParams{
			CourseID:     int32(course_id_int),
			RoleID:       int32(role_id_int),
			StudentID:    student_id,
			FirstName:    first_name,
			LastName:     last_name,
			Email:        email,
			PasswordHash: string(passwordHash),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) VerifyCredentials(
	ctx context.Context,
	student_id string,
	password string,
) error {
	dbPassword, err := s.db.GetPasswordHashByStudentId(ctx, student_id)
	if err != nil {
		return fmt.Errorf("invalid or non-existent email: %w", err)
	}

	if dbPassword == "" {
		return fmt.Errorf("stored password hass is not valid")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(dbPassword),
		[]byte(password),
	)
	if err != nil {
		return fmt.Errorf("invalid or wrong password: %w", err)
	}

	return nil
}

func (s *AuthService) Login(
	ctx context.Context,
	sh *session.SessionHandler,
	user_id int32,
) error {
	sess, ok := session.GetRequestSession(ctx)
	if !ok {
		return fmt.Errorf("failed to get session from client request")
	}

	if err := sh.Migrate(ctx, sess); err != nil {
		return err
	}

	val := strconv.Itoa(int(user_id))
	if err := s.store.PutSessionData(ctx, sess, userIdKey, val); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Logout(
	ctx context.Context,
	sh *session.SessionHandler,
) error {
	sess, ok := session.GetRequestSession(ctx)
	if !ok {
		return fmt.Errorf("failed to get session from client request")
	}

	_, ok = sess.Data["user_id"]
	if !ok {
		return fmt.Errorf("non authorized session or your session has been timed out")
	}

	if err := sh.Migrate(ctx, sess); err != nil {
		return err
	}

	s.store.DeleteSessionData(ctx, sess, userIdKey)
	return nil
}

func NewAuthService(db *postgres.Queries, sessSvc session.SessionServiceInterface) AuthServiceInterface {
	return &AuthService{
		db:    db,
		store: sessSvc,
	}
}

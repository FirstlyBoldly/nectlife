package invitation

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nectgrams-webapp-team/nectlife/internal/config"
	"github.com/nectgrams-webapp-team/nectlife/internal/mail"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
)

type InvitationService struct {
	mailer *mail.MailService
	db     *postgres.Queries
}

func NewInvitationService(mailer *mail.MailService, db *postgres.Queries) InvitationServiceInterface {
	return &InvitationService{mailer: mailer, db: db}
}

func (s *InvitationService) Create(ctx context.Context, invitedByUserId, roleId int32, studentId, email string) (postgres.Invitation, error) {
	var token string
	return s.db.CreateInvitation(ctx, postgres.CreateInvitationParams{
		InvitedByUserID: pgtype.Int4{
			Int32: roleId,
			Valid: true,
		},
		StudentID: studentId,
		Email:     email,
		RoleID:    roleId,
		Token:     token,
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().UTC().Add(time.Duration(config.Data.EXPIRES_AT_IN_HOURS) * time.Hour),
			Valid: true,
		},
	})
}

func (s *InvitationService) Send(studentId, email, subject, link string) error {
	data := struct {
		Name string
		Link string
	}{
		Name: studentId,
		Link: link,
	}
	return s.mailer.Send(email, subject, "invitation.html", data)
}

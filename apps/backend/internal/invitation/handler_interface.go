package invitation

import "github.com/nectgrams-webapp-team/nectlife/internal/postgres"

type InviteInput struct {
	Body postgres.CreateInvitationParams
}

type InvitationHandlerInterface interface {
}

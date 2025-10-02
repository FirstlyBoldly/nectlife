package invitation

import "context"

type InvitaionHandler struct {
	svc *InvitationService
}

func NewInvitationHandler(svc *InvitationService) InvitationHandlerInterface {
	return &InvitaionHandler{svc: svc}
}

func (h *InvitaionHandler) Invite(ctx context.Context, input *InviteInput) (*struct{}, error) {
	return nil, nil
}

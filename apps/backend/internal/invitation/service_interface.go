package invitation

type InvitationServiceInterface interface {
	Send(studentId, email, subject, link string) error
}

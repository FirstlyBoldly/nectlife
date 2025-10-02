package mail

type MailServiceInterface interface {
	Send(to, subject, tplFile string, data interface{}) error
}

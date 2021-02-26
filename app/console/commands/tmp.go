package commands

// import (
// 	"confetti-framework/app/support"
// 	"github.com/confetti-framework/contract/inter"
// 	"io"
// )
//
// type SendEmails struct {
// 	Names []string `short:"N" flag:"names"`
// 	Ids []int `short:"I" flag:"ids"`
// }
//
// func (s SendEmails) Name() string {
// 	return "mail:send"
// }
//
// func (s SendEmails) Description() string {
// 	return "Send a marketing email to a user."
// }
//
// func (s SendEmails) Handle(app inter.App, writer io.Writer) inter.ExitCode {
// 	mailer := app.Make(support.DripEmailer{}).(support.DripEmailer)
// 	mailer.Send(s.Email)
//
// 	return inter.Success
// }

package function

import (
	"pusherGo/domain"
	"pusherGo/global"

	"gopkg.in/gomail.v2"
)

func sendEmail(request *domain.EmailRequest) error {
	mailHeader := map[string][]string{
		"From":    {request.From},
		"To":      request.To,
		"Subject": {request.Subject},
	}
	m := gomail.NewMessage()
	m.SetHeaders(mailHeader)
	m.SetBody("text/html", request.Body)

	d := gomail.NewDialer("smtp.gmail.com", 587, request.From, request.Password)
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}

func SendEmail(text string) {
	request := &domain.EmailRequest{
		From:     global.Configs.Email.From,
		Password: global.Configs.Email.Password,
		To:       global.Configs.Email.To,
		Subject:  global.Configs.Email.Subject,
		Body:     text,
	}
	err := sendEmail(request)
	if err != nil {
		panic(err)
	}
}

func SendError(err error) {
	request := &domain.EmailRequest{
		From:     global.Configs.Email.From,
		Password: global.Configs.Email.Password,
		To:       global.Configs.Email.To,
		Subject:  "Error",
		Body:     err.Error(),
	}
	err = sendEmail(request)
	if err != nil {
		panic(err)
	}
}

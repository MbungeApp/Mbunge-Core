package utils

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/config"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var temp, key string
var err error

func init() {
	temp, err = config.GetKey("sendgrid", "template_id")
	if err != nil {
		panic(err)
	}
	key, err = config.GetKey("sendgrid", "api_key")
	if err != nil {
		panic(err)
	}
}

func dynamicTemplateEmail(email string, otpCode int) []byte {
	m := mail.NewV3Mail()

	address := "auth-noreply@mbunge.app"
	name := "Mbunge App"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	m.SetTemplateID(temp)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("", email),
	}

	p.AddTos(tos...)
	p.SetDynamicTemplateData("otp_code", otpCode)
	m.AddPersonalizations(p)

	return mail.GetRequestBody(m)
}

func SendMail(email string, otpCode int) {
	request := sendgrid.GetRequest(key, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = dynamicTemplateEmail(email, otpCode)
	request.Body = Body
	_, err := sendgrid.API(request)
	if err != nil {
		config.ErrorReporter(err.Error())
		fmt.Println(err)
	}
	//else {
	//	fmt.Println(response.StatusCode)
	//	fmt.Println(response.Body)
	//	fmt.Println(response.Headers)
	//}
}

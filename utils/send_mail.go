/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package utils

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/config"
	"log"

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

	address := "admin@mbungeapp.tech"
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

func SendRandomEmail(email string, message string) {
	from := mail.NewEmail("Mbunge App", "admin@mbungeapp.tech")
	subject := "Your mbungeApp admin login credentials"
	to := mail.NewEmail("", email)
	plainTextContent := "Hello there and welcome\n"
	htmlContent := fmt.Sprintf("Your password is <strong>%s\n</strong>", message)
	msg := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(key)
	response, err := client.Send(msg)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

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

}

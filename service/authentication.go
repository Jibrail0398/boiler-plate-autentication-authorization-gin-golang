package service

import (
	
	"fmt"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/helper"
	"gopkg.in/gomail.v2"
)


type AuthenticationService interface{
	SendVerificationCode() error
}

type authenticationService struct{}

func NewAuthenticationService() AuthenticationService {
	return &authenticationService{}
}


func (s * authenticationService) SendVerificationCode()error{

	emailConfig,err := helper.GetGomailConfig();
	if err!=nil{
		return err 
	}


	mailer := gomail.NewMessage()
    mailer.SetHeader("From", emailConfig.CONFIG_SENDER_NAME)
    mailer.SetHeader("To", "izamikatsuka@gmail.com",)
    
    mailer.SetHeader("Subject", "Test mail")
    mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
    mailer.Attach("./kucing muntah.jpg")

	dialer := gomail.NewDialer(
        emailConfig.CONFIG_SMTP_HOST,
        emailConfig.CONFIG_SMTP_PORT,
        emailConfig.CONFIG_AUTH_EMAIL,
        emailConfig.CONFIG_AUTH_PASSWORD,
		
    )

	
    err = dialer.DialAndSend(mailer)
    if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
    }

	return nil
}


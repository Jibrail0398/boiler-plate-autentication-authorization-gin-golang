package service

import (
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/helper"
	"fmt"
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

	randomCode,err := helper.GenerateCodeVerif(5)

	client := helper.ConnectToRedis()
	helper.StoreWithTime("randomCode",randomCode, 300, client)

	if err!=nil{
		return err
	}
	dataSend := map[string]string{
		"Email":emailConfig.CONFIG_AUTH_EMAIL,
		"Code":randomCode,
	}

	htmlBody,err := helper.ParseEmailTemplate("views/email.html",dataSend)

	if err!=nil{
		return err
	}

	mailer := gomail.NewMessage()
    mailer.SetHeader("From", emailConfig.CONFIG_SENDER_NAME)
    mailer.SetHeader("To", "izamikatsuka@gmail.com",)
    
    mailer.SetHeader("Subject", "Test mail")
    mailer.SetBody("text/html", htmlBody)
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


package service

import (
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/model"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)


type AuthenticationService interface{
	SendVerificationCode() error
}

type authenticationService struct{}

func NewAuthenticationService() AuthenticationService {
	return &authenticationService{}
}

func GetGomailConfig() (model.EmailVerifConfig,error) {
	err := godotenv.Load()

	if err!=nil{
		return model.EmailVerifConfig{}, fmt.Errorf("failed to load env file")
	}

	HOST := os.Getenv("CONFIG_SMTP_HOST");
	PORT := os.Getenv("CONFIG_SMTP_PORT");
	NAME := os.Getenv("CONFIG_SENDER_NAME");
	EMAIL := os.Getenv("CONFIG_AUTH_EMAIL");
	PASSWORD := os.Getenv("CONFIG_AUTH_PASSWORD");
	PORTINT,_ := strconv.Atoi(PORT)

	emailVerifConfig := model.EmailVerifConfig{
		CONFIG_SMTP_HOST:HOST,
		CONFIG_SMTP_PORT:PORTINT,
		CONFIG_SENDER_NAME:NAME,
		CONFIG_AUTH_EMAIL:EMAIL,
		CONFIG_AUTH_PASSWORD:PASSWORD,
	}
	
	return emailVerifConfig,nil
}

func (s * authenticationService) SendVerificationCode()error{

	emailConfig,err := GetGomailConfig();
	if err!=nil{
		return err 
	}


	mailer := gomail.NewMessage()
    mailer.SetHeader("From", emailConfig.CONFIG_SENDER_NAME)
    mailer.SetHeader("To", "izamikatsuka@gmail.com",)
    // mailer.SetAddressHeader("Cc", "if22.muhammadnatadilaga@mhs.ubpkarawang.ac.id", "Hello")
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


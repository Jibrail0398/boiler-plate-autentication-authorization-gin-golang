package service

import (
	"github.com/joho/godotenv"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/model"
	"os"
	"fmt"
	
)

type authenticationService struct{}

func NewAuthenticationService() *authenticationService {
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

	emailVerifConfig := model.EmailVerifConfig{
		CONFIG_SMTP_HOST:HOST,
		CONFIG_SMTP_PORT:PORT,
		CONFIG_SENDER_NAME:NAME,
		CONFIG_AUTH_EMAIL:EMAIL,
		CONFIG_AUTH_PASSWORD:PASSWORD,
	}
	

	return emailVerifConfig,nil


}

// func SendVerificationCode(){
	
// }


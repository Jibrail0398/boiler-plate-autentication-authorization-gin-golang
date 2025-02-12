package helper

import(
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/model"
	"github.com/joho/godotenv"
	"strconv"
	"os"
	"fmt"
)

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
package helper

import(
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/model"
	"github.com/joho/godotenv"
	"strconv"
	"os"
	"fmt"
	"crypto/rand"
	"math/big"
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

func GenerateCodeVerif(length int) (string,error){

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	code := make([]byte,length)

	
	for i:= range code{
		// Menghasilkan indeks acak untuk memilih karakter dari charset
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("generate random code failed: %v", err)
		}
		
		code[i] = charset[num.Int64()]
	}
	
	

	return string(code),nil

	
}
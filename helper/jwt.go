package helper

import (
	
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/model"
	"time"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"github.com/golang-jwt/jwt/v5"
)

func GetJWTKey() (string,error) {
	err := godotenv.Load()
	if err!=nil{
		return "", fmt.Errorf("failed to load env file")
	}

	jwtkey := os.Getenv("JWTKEY");
	
	return jwtkey,nil
}

func GenerateJWT(username string, email string) (string,error){
	expirationTime := time.Now().Add(5 * time.Minute) // 5 menit

	
	claims := &model.Claims{
		Username: username,
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtkey,err := GetJWTKey()
	jwtKey := []byte(jwtkey)
	if err!=nil{
		fmt.Println("error get jwt key")
	}

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString)

	return tokenString, nil
}

func ValidateTokenJWT(tknStr string) (* model.Claims, error) {
	claims := &model.Claims{}

	jwtkey,err := GetJWTKey()
	jwtKey := []byte(jwtkey)
	if err!=nil{
		fmt.Println("error get jwt key")
	}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
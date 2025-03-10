package helper

import(
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/model"
	"github.com/joho/godotenv"
	"strconv"
	"os"
	"fmt"
	"crypto/rand"
	"math/big"
	"html/template"
	"bytes"
	"encoding/base64"
	"regexp"
	"github.com/go-playground/validator/v10"
	
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

func ParseEmailTemplate(filename string, data map[string]string) (string,error){
	// Baca file HTML
	htmlBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	htmlString := string(htmlBytes)

	// Parse template
	tmpl, err := template.New("email").Parse(htmlString)
	if err != nil {
		return "", err
	}

	// Replace variabel dalam template
	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return "", err
	}

	return body.String(), nil
}

func GetOauthGoogleConfig() (model.GoogleOauthConfig,error) {
	err := godotenv.Load()
	if err!=nil{
		return model.GoogleOauthConfig{}, fmt.Errorf("failed to load env file")
	}

	client_id := os.Getenv("CLIENT_ID");
	client_secret := os.Getenv("CLIENT_SECRET");

	googleOauthConfig := model.GoogleOauthConfig{
		CLIENT_ID: client_id,
		CLIENT_SECRET: client_secret,
	}

	return googleOauthConfig,nil
}



func GenerateStateOauthCookie() string {
    b := make([]byte, 16)
    rand.Read(b)
    return base64.URLEncoding.EncodeToString(b)
}


func containsSpecialCharacter(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]+`)
	return re.MatchString(fl.Field().String())
}


func containsNumber(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[0-9]+`)
	return re.MatchString(fl.Field().String())
}

func RegisterNewValidator(validate validator.Validate){
	validate.RegisterValidation("contains_special", containsSpecialCharacter)
	validate.RegisterValidation("contains_number", containsNumber)
}


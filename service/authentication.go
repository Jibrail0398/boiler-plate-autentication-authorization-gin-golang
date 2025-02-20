package service

import (
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/db"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"
	
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)



type AuthenticationService interface{
	SendVerificationCode(email string) error
	RegisterGoogle(arg db.RegisterGoogleParams) error
	ManualRegister(arg db.RegisterManualParams) error
	ManualLogin(email string ,password string) (string ,error) 
	VerifyUser(key string, value string, arg db.VerifiedUserParams) (error)
}

type authenticationService struct{
	db *db.Queries
}

func NewAuthenticationService(db *db.Queries) AuthenticationService {
	return &authenticationService{db:db}
}


func (s * authenticationService) SendVerificationCode(email string)error{

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
		"Email":email,
		"Code":randomCode,
	}

	htmlBody,err := helper.ParseEmailTemplate("views/email.html",dataSend)

	if err!=nil{
		return err
	}

	mailer := gomail.NewMessage()
    mailer.SetHeader("From", emailConfig.CONFIG_SENDER_NAME)
    mailer.SetHeader("To", email)
    
    mailer.SetHeader("Subject", "Verification Code")
    mailer.SetBody("text/html", htmlBody)
    

	dialer := gomail.NewDialer(
        emailConfig.CONFIG_SMTP_HOST,
        emailConfig.CONFIG_SMTP_PORT,
        emailConfig.CONFIG_AUTH_EMAIL,
        emailConfig.CONFIG_AUTH_PASSWORD,
		
    )
	
    err = dialer.DialAndSend(mailer)
    if err != nil {
		return fmt.Errorf("failed to send email to email address %s: %w",email, err)
    }

	return nil
}

func (s * authenticationService) RegisterGoogle(arg db.RegisterGoogleParams ) error {
	ctx := context.Background()


	user,err:= s.db.GetUsersByEmail(ctx,arg.Email)
	if err!=nil{
		if errors.Is(err, sql.ErrNoRows){
			return fmt.Errorf("user with email %s not found",arg.Email)
		}
		return fmt.Errorf("error get user data by email: %v",err)
	}

	if !user.Password.Valid  && !user.OauthID.Valid{

		err = s.db.RegisterGoogle(ctx,arg)
	
		if err!=nil{
			return fmt.Errorf("error insert user to database: %v",err)
			
		}
	}
	


	return nil
}

func (s *authenticationService) ManualRegister(arg db.RegisterManualParams) error{

	//check if email has been registered
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	
	user,err := s.db.GetUsersByEmail(ctx,arg.Email)
	
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		
		return fmt.Errorf("error while get user by email: %w", err)
	}
	
	if user != (db.User{}) {
		
		return fmt.Errorf("email has been registered")
	}
	
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(arg.Password.String),bcrypt.DefaultCost)
	if err!=nil{
		return fmt.Errorf("error while hashing password: %v",err)
	}

	arg.Password.String = string(hashedPassword)


	err = s.db.RegisterManual(ctx, arg)

	if err!=nil{
		
		return fmt.Errorf("error while registering user: %v",err)
	}


	return nil
}

func(s *authenticationService) VerifyUser(key string, value string, arg db.VerifiedUserParams) (error){
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	client := helper.ConnectToRedis()
	code,err := helper.GetDataRedis(key,client)
	if err!=nil{
		return fmt.Errorf("verification code expired")
	}

	if value != code{
		return fmt.Errorf("otp doesn't match")
	}

	err = s.db.VerifiedUser(ctx,arg)

	if err!=nil{
		return fmt.Errorf("error while updating verify user: %v", err.Error())
	}

	return nil
}

func (s *authenticationService) ManualLogin(email string ,password string) (string,error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	//Get Data By Email
	user,err:= s.db.GetUsersByEmail(ctx,email);

	//throw error if user not found
	if err!= nil && errors.Is(err,sql.ErrNoRows){
		return "",fmt.Errorf("user with email %s not found",email)
	}

	//throw error if user not verified
	if !user.Verified{
		return "",fmt.Errorf("email has not been verified")
	}

	//throw error if password not correct
	
	err =  bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))
	if err!=nil{
		fmt.Println("password from db:",user.Password.String)
		fmt.Println(password)
		return "",fmt.Errorf("password inccorect")
	}

	//generate token JWT

	token,err := helper.GenerateJWT(user.Name,user.Email);
	if err!=nil{
		return "",fmt.Errorf("error generate token")
	}

	return token,nil
}


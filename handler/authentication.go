package handler

import (
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/db"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/helper"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/service"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/model"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

)



type authenticationHandler struct{
	AuthenticationService service.AuthenticationService
}

func NewAuthenticationHandler(authenticationService service.AuthenticationService) *authenticationHandler{
	return &authenticationHandler{
		AuthenticationService: authenticationService,
	}
}

func (h *authenticationHandler) TryHelper(c *gin.Context){
	tryHelper,err := helper.ValidateTokenJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkamkiLCJlbWFpbCI6ImFkamlAZ21haWwiLCJleHAiOjE3NDAwMTAyODF9.ra-OqjU80QxUz0cJJu_pmIDzarAZWVtaS3yhczqpNpM");

	// tryHelper,err := helper.GenerateJWT("adji","adji@gmail");

	if err!=nil{
		c.AbortWithStatusJSON(500,gin.H{
			"Error":err.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(200,gin.H{
		"Hasil Response":tryHelper,
	})
}
func(h *authenticationHandler) Login(c *gin.Context) {
	c.AbortWithStatusJSON(200,gin.H{
		"message":"Behasil login",
	})
}

func(h *authenticationHandler)SendVerificationCode(c *gin.Context){
	email:=c.Request.FormValue("email")
	
	err:= h.AuthenticationService.SendVerificationCode(email)
	
	if err!=nil{
		c.AbortWithStatusJSON(200,gin.H{"Error":err.Error()})
		return
	}

	c.AbortWithStatusJSON(
		200,gin.H{
			"Message":"Email already sent",
			
		})
}

func (h *authenticationHandler) LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

var secretConfig,_ = helper.GetOauthGoogleConfig()
var randomstate = helper.GenerateStateOauthCookie()

var(
	googleOauthConfig = &oauth2.Config{
		ClientID: secretConfig.CLIENT_ID,
		ClientSecret: secretConfig.CLIENT_SECRET,
		RedirectURL: "http://localhost:8080/google/callback",
		Scopes:       []string{"email", "profile"},
        Endpoint:     google.Endpoint,
	}
	
)

func (h *authenticationHandler) HandleGoogleLogin(c *gin.Context) {
	
    url := googleOauthConfig.AuthCodeURL(randomstate)
    c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *authenticationHandler) HandleGoogleCallback(c *gin.Context){

	randomstate := helper.GenerateStateOauthCookie()
	state := c.Query("state")
    if state != randomstate {
        c.AbortWithStatusJSON(400,gin.H{"Error":"States does not match"})
    }

	
	code := c.Query("code")
	
    token, err := googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	// log.Println("Received token:", token)

	client := googleOauthConfig.Client(context.Background(), token)
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info"})
        return
    }
    defer resp.Body.Close()	

	
    var userInfo struct {
        ID    string `json:"id"`
        Email string `json:"email"`
        Name  string `json:"name"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode user info"})
        return
    }

	registerGoogleParams:= db.RegisterGoogleParams{
		Name: userInfo.Name,
		Email: userInfo.Email,
		OauthProvider: sql.NullString{String: "Google", Valid: true},
		OauthID:       sql.NullString{String: userInfo.ID, Valid: true},
	}

	err = h.AuthenticationService.RegisterGoogle(registerGoogleParams)

	if err!=nil{
		c.AbortWithStatusJSON(500,gin.H{"error":err.Error()})
	}

}

func (h *authenticationHandler) ManualRegister(c *gin.Context){
	
	var newUser model.RegisterRequest

	//Get Request
	if err:= c.ShouldBind(&newUser); err!=nil{
		c.AbortWithStatusJSON(400,gin.H{
			"Error ":"Error while bind request ",
			"Message":err.Error(),
		})
		return
	}

	//Validation Process
	validate := validator.New()
	helper.RegisterNewValidator(*validate)
	if err:=validate.Struct(newUser); err!=nil{

		errs:=err.(validator.ValidationErrors)
		c.AbortWithStatusJSON(400,gin.H{"Error":errs.Error()})
		log.Println(errs)
		return
	}

	//Save User to Database

	registerManualParams := db.RegisterManualParams{
		Name: newUser.Name,
		Email: newUser.Email,
		Password: sql.NullString{String:newUser.Password,Valid: true},
		Verified: false,

	}
	err := h.AuthenticationService.ManualRegister(registerManualParams)
	if err!=nil{
		
		c.AbortWithStatusJSON(400,gin.H{"Error":err.Error()})
		return
	}

	//Send Verification Email
	err= h.AuthenticationService.SendVerificationCode(registerManualParams.Email)
	if err!=nil{
		c.AbortWithStatusJSON(500,gin.H{"Error":err.Error()})
		return
	}

	c.AbortWithStatusJSON(201,gin.H{
		"Status":"Success",
		"Message":"User registered successfully, See your email to verify your Account",
		
	})

	
}

func(s *authenticationHandler) VerifyUser(c *gin.Context) {

	email := c.Request.FormValue("email")
	value := c.Request.FormValue("code")

	verifiedUserParams := &db.VerifiedUserParams{
		Verified: true,
		Email: email,
	}

	err := s.AuthenticationService.VerifyUser("randomCode",value,*verifiedUserParams)
	if err!=nil{
		c.AbortWithStatusJSON(400,gin.H{"error":err.Error()})
		return 
	}

	c.AbortWithStatusJSON(200,gin.H{"Message":"User has been successfully verified"})

}


func(h *authenticationHandler) ManualLogin(c *gin.Context){
	
	var user model.LoginRequest

	if err := c.ShouldBind(&user); err!=nil{
		c.AbortWithStatusJSON(500,gin.H{
			"Error ":"Error while bind request ",
			"Message":err.Error(),
		})
		return
	}

	token,err:= h.AuthenticationService.ManualLogin(user.Email,user.Password)

	if err!=nil{
		c.AbortWithStatusJSON(500,gin.H{"Error":err.Error()})
		return
	}

	c.AbortWithStatusJSON(200,gin.H{
		"Message":"Login Succeed",
		"email":user.Email,
		"token":token,
	})
}
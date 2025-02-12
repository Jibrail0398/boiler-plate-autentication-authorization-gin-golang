package handler

import (
	"github.com/gin-gonic/gin"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/service"
)

type authenticationHandler struct{
	AuthenticationService service.AuthenticationService
}

func NewAuthenticationHandler(authenticationService service.AuthenticationService) *authenticationHandler{
	return &authenticationHandler{
		AuthenticationService: authenticationService,
	}
}

func(h *authenticationHandler) Login(c *gin.Context) {
	c.AbortWithStatusJSON(200,gin.H{
		"message":"Behasil login",
	})
}

func(h *authenticationHandler)SendVerificationCode(c *gin.Context){
	err:= h.AuthenticationService.SendVerificationCode()
	if err!=nil{
		c.AbortWithStatusJSON(500,gin.H{
			"Message":err,
		})
		return
	}

	c.AbortWithStatusJSON(
		500,gin.H{
			"Message":"Email already sent",
		})
}
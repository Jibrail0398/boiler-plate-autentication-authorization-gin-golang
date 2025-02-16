package handler

import (
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/helper"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/service"

	"github.com/gin-gonic/gin"
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
	tryHelper,_ := helper.GenerateCodeVerif(5);

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
	err:= h.AuthenticationService.SendVerificationCode()
	client:= helper.ConnectToRedis()

	randomCode := helper.GetDataRedis("randomCode", client)
	if err!=nil{
		c.AbortWithStatusJSON(500,gin.H{
			"Message":err,
		})
		return
	}

	c.AbortWithStatusJSON(
		200,gin.H{
			"Message":"Email already sent",
			"Random Code":randomCode,
		})
}
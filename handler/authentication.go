package handler

import "github.com/gin-gonic/gin"

type authenticationHandler struct{

}

func NewAuthenticationHandler() *authenticationHandler{
	return &authenticationHandler{

	}
}

func(h *authenticationHandler) Login(c *gin.Context) {
	c.AbortWithStatusJSON(200,gin.H{
		"message":"Behasil login",
	})
}
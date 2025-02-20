package middleware

import (
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenticationMiddleware struct{}

func (m *AuthenticationMiddleware) Authentication(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"Error": "Unauthorized request",
		})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.AbortWithStatusJSON(400, gin.H{
			"Error": "Invalid token format",
		})
		return
	}

	tokenString := parts[1]

	claims, err := helper.ValidateTokenJWT(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
		return
	}

	//Simpan ke context
	c.Set("userClaims", claims)

	c.Next()

}

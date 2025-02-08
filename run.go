package main

import "github.com/gin-gonic/gin"
import "Jibrail0398/boiler-plate-autentication-authorization-gin-golang/handler"

func RunServer() {
	r := gin.Default();

	handler := handler.NewAuthenticationHandler();

	r.GET("/",handler.Login)

	r.Run();
}
package run

import (
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/db"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/handler"
	"Jibrail0398/boiler-plate-autentication-authorization-gin-golang/service"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


func RunServer() {


	//Database
	database := db.NewDatabase();

	credential := db.Credential{
		Host		:"localhost",
		Username	:"postgres",
		Password	:"jibrailadji02",
		DatabaseName:"temenan",
		Port		: 5432,
	}

	_,err:=database.Connect(credential) 

	if err!=nil{
		log.Fatal("Database Connection Error")
	}

	defer database.DB.Close()

	err = database.Up()
	if err!=nil{
		log.Fatal("Migration Database Failed : ",err)
	}


	//Endpoint
	r := gin.Default();


	service := service.NewAuthenticationService()
	handler := handler.NewAuthenticationHandler(service);

	r.GET("/",handler.Login)
	r.POST("/send-email",handler.SendVerificationCode)

	

	if err!=nil{
		log.Fatal("Gagal load Gomail config")
	}

	

	r.Run();

	
}



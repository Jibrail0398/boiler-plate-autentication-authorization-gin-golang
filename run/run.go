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

	log.Println("Database Connected")

	defer database.DB.Close()

	err = database.Up()
	if err!=nil{
		log.Fatal("Migration Database Failed : ",err)
	}


	//Endpoint
	r := gin.Default();

	queries := db.New(database.DB)
	
	service := service.NewAuthenticationService(queries)
	handler := handler.NewAuthenticationHandler(service);
	
	//endpoint untuk coba helper function
	r.GET("/try-helper",handler.TryHelper)

	r.LoadHTMLFiles("index.html")

	r.GET("/",handler.Login)
	r.GET("/auth/oauth",handler.LoginHandler)
	r.GET("/google/login", handler.HandleGoogleLogin)
    r.GET("/google/callback", handler.HandleGoogleCallback)

	r.POST("/register",handler.ManualRegister)
	r.POST("/login", handler.ManualLogin)
	r.POST("/send-email",handler.SendVerificationCode)
	r.POST("/verify",handler.VerifyUser)
	

	if err!=nil{
		log.Fatal("Gagal load Gomail config")
	}


	r.Run();

	
}



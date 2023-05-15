package transport

import (
	"net/http"

	_ "github.com/shop/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go Restful API with Swagger
// @version 1.0
// @description Simple swagger implementation in Go HTTP

// @contact.name Linggar Primahastoko
// @contact.url https://linggar.asia
// @contact.email x@linggar.asia

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8082
// @BasePath /
func ServerApp() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	http.HandleFunc("/registration", Registration)
	// http.HandleFunc("/users/profile", userProfile)

	http.ListenAndServe(":8000", nil)
}
